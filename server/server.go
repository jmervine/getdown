package server

import (
	"log"
	"net/http"
	//"os"
	"time"

	"github.com/jmervine/getdown/config"
	"github.com/jmervine/getdown/markdown"
	"github.com/jmervine/getdown/template"
)

// Allow for overite of default logging fucntion when
// testing. Exporting to allow for other logging functions
// as well.
var LogF = log.Printf

// Logger provides a good loggging mechanism.
func Logger(at, meth string, stat int, url string, begin *time.Time, meta interface{}) {
	var metaLabel string
	var metaValue string

	switch m := meta.(type) {
	case string:
		metaLabel = "source"
		metaValue = m
	case error:
		metaLabel = "error"
		metaValue = "\"" + m.Error() + "\""
	}

	str := "at=%s method=%s status=%d %s=%v url=%s"
	if begin != nil { // mostly for testing
		str = str + " duration=%s\n"
		LogF(str, at, meth, stat, metaLabel, metaValue, url, time.Since(*begin))
	} else {
		str = str + "\n"
		LogF(str, at, meth, stat, metaLabel, metaValue, url)
	}
}

// FileServerWithLogger does what you would expect.
func FileServerWithLogger(begin *time.Time, cfg *config.Config) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		http.FileServer(http.Dir(cfg.Basedir)).ServeHTTP(w, r)
		Logger("request", r.Method, 200, r.URL.Path, begin, "static")
	}

	return http.HandlerFunc(handler)
}

// Bind http server.
func Bind(cfg *config.Config) *http.Server {
	handler := func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		begin := &now

		md, err := markdown.New(r.URL.Path, cfg)
		if err != nil {
			// serve static
			FileServerWithLogger(begin, cfg).ServeHTTP(w, r)
			return
		}

		payload, err := template.NewPayload(cfg, &md)
		if err != nil {
			// handle error
			Logger("request", r.Method, 500, r.URL.Path, begin, err)
			return
		}

		payload.Render(w)
		Logger("request", r.Method, 200, r.URL.Path, begin, md.Path)
	}

	return &http.Server{
		Addr:    cfg.Listener(),
		Handler: http.HandlerFunc(handler),
	}
}
