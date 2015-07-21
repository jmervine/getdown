/*
ll (short for linelogger) is a vary simple Scrolls'y (https://github.com/asenchi/scrolls)
style logger.

Example:

    package main
    import "time"
    import "gopkg.in/jmervine/ll.v1"

    func main() {
        begin := time.Now()

        // ... do stuff ...

        ll.Log(&begin, map[string]interface{} {
            "at": "main",
            "data": "foo",
        }
    }

    // Output:
    //
    // 2015/07/21 04:57:22 level=info at=main foo=bar durration=102.972us
*/
package ll

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// Logger is exported as a pass through to log.Logger under the
// hood, so default functions can still be called.
//
// Example:
//
// Logger.Fatal("ack")
var Logger *log.Logger

// Debugger is exported as a pass through to log.Logger under the
// hood, so default functions can still be called.
//
// Example:
//
// Debugger.Println("ack")
//
// Warning: this will exit with no output when not in debug mode.
// Debugger.Fatal("ack")
//
// TODO: create noop logger for when not in debug mode.
var Debugger *log.Logger

func init() {
	Logger = log.New(os.Stdout, "", log.Flags())

	// if anything but '/true/i', '/t/i' or '1' hide debug output
	if !hasDebug() {
		Debugger = log.New(ioutil.Discard, "", log.Flags())
	} else {
		Debugger = Logger
	}
}

func hasDebug() bool {
	if ok, _ := strconv.ParseBool(os.Getenv("DEBUG")); !ok {
		return false
	}
	return true
}

// Logger provides a good loggging mechanism.
func logger(target *log.Logger, level string, begin *time.Time, meta map[string]interface{}) {
	toS := func(k string, i interface{}) string {
		return fmt.Sprintf("%s=%v", k, i)
	}
	toQ := func(k string, i interface{}) string {
		return fmt.Sprintf("%s=\"%s\"", k, i)
	}

	var line = []string{toS("level", level)}
	for key, val := range meta {
		var pair string
		switch v := val.(type) {
		case error:
			pair = toQ(key, v.Error())
		case []string:
			pair = toQ(key, strings.Join(v, ","))
		}

		if pair == "" {
			pair = toS(key, val)
		}

		line = append(line, pair)
	}

	if begin != nil {
		line = append(line, toS("durration", time.Since(*begin)))
	}

	target.Println(strings.Join(line, " "))
}

// SetOutput allows you to change the output destination of both Logger
// and Debugger in one shot.
//
// Example:
//
//     SetOutput(os.Stderr)
func SetOutput(out io.Writer) {
	Logger = log.New(out, "", log.Flags())
	if hasDebug() {
		Debugger = log.New(out, "", log.Flags())
	}
}

// Log is the standard logger, always logging to os.Stdout by default.
//
// Example usage:
//
//     begin := time.Now()
//     // do stuff
//     Log(&begin, map[string]interface{} {
//         "at": "request",
//         "method": "GET",
//         "url: "/path/to/file.html",
//         "error": errors.New("something bad happened"),
//     })
//
//     // Outputs:
//     YYYY-MM-DD HH:MM:SS at=request method=GET url=/path/to/file.html error="something bad happened" durration=#ns
//
//     // without time
//     Log(nil, map[string]interface{}{
//         "at": "request",
//         "request": fmt.Sprintf("%+v", *req),
//     })
//
//     // Outputs:
//     YYYY-MM-DD HH:MM:SS at=request request={ ... request args ... }
//
func Log(begin *time.Time, meta map[string]interface{}) {
	logger(Logger, "info", begin, meta)
}

// Debug is the debug (DEBUG=true) logger, logging to os.Stdout by default
// when os.Getenv("DEBUG") is true.
func Debug(begin *time.Time, meta map[string]interface{}) {
	logger(Debugger, "debug", begin, meta)
}
