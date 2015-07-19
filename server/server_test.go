package server

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	. "gopkg.in/jmervine/GoT.v1"

	"github.com/jmervine/getdown/config"
)

func TestBind(T *testing.T) {
	svr := Bind(&config.Default)

	Go(T).RefuteNil(svr)
}

func TestLogger(T *testing.T) {
	var recorded string
	LogF = func(s string, i ...interface{}) {
		recorded = strings.TrimSpace(fmt.Sprintf(s, i...))
	}

	Logger("at", "meth", 200, "url", nil, "meta")
	Go(T).AssertEqual("at=at method=meth status=200 source=meta url=url",
		recorded)

	Logger("at", "meth", 500, "url", nil, errors.New("error message"))
	Go(T).AssertEqual("at=at method=meth status=500 error=\"error message\" url=url",
		recorded)

}
func ExampleLogger() {
	// at=at method=meth status=200 source=meta url=url
	// at=at method=meth status=500 error=error url=url
}
