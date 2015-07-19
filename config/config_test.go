package config

import (
	"testing"

	. "gopkg.in/jmervine/GoT.v1"
)

func TestDefault(T *testing.T) {
	Go(T).AssertEqual(Default.Addr, Config{Addr: "localhost"}.Addr)
}

func TestListener(T *testing.T) {
	c := Config{
		Port: "3000",
		Addr: "localhost",
	}
	Go(T).AssertEqual("localhost:3000", c.Listener())
}
