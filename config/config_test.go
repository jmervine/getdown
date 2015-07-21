package config

import (
	"path"
	"testing"

	. "github.com/jmervine/getdown/Godeps/_workspace/src/gopkg.in/jmervine/GoT.v1"
)

func TestDefault(T *testing.T) {
	Go(T).AssertEqual(Default.Addr, Config{Addr: "localhost"}.Addr)
}

func TestNew(T *testing.T) {
	c := New("", "", "basedir", "", "", "", "", "")

	Go(T).AssertEqual(c.Addr, Default.Addr)
	Go(T).AssertEqual(c.Static, path.Join(c.Basedir, "static"))

	c = New("", "", "basedir", "dir/static", "", "", "", "")

	Go(T).AssertEqual(c.Static, "dir/static")
}

func TestListener(T *testing.T) {
	c := Config{
		Port: "3000",
		Addr: "localhost",
	}
	Go(T).AssertEqual("localhost:3000", c.Listener())
}
