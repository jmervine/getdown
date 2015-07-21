package server

import (
	"testing"

	. "github.com/jmervine/getdown/Godeps/_workspace/src/gopkg.in/jmervine/GoT.v1"

	"github.com/jmervine/getdown/config"
)

func TestBind(T *testing.T) {
	svr := Bind(&config.Default)

	Go(T).RefuteNil(svr)
}
