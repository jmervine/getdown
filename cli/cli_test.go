package cli

import (
	"os"
	"testing"

	. "github.com/jmervine/getdown/Godeps/_workspace/src/gopkg.in/jmervine/GoT.v1"
	"github.com/jmervine/getdown/config"
)

func TestParseWithDefaults(T *testing.T) {
	cfg := Parse([]string{"app"})

	Go(T).RefuteNil(cfg)
	Go(T).AssertEqual(config.Default.Addr, cfg.Addr)
}

func TestParseWithArgs(T *testing.T) {
	addr := "argaddr"
	cfg := Parse([]string{"app", "-a", addr})

	Go(T).RefuteNil(cfg)
	Go(T).AssertEqual(addr, cfg.Addr)
}

func TestParseWithEnvVars(T *testing.T) {
	addr := "envaddr"

	os.Setenv("GETDOWN_ADDR", addr)
	cfg := Parse([]string{"app"})

	Go(T).RefuteNil(cfg)
	Go(T).AssertEqual(addr, cfg.Addr)
}
