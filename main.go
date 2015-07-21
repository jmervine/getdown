package main

import (
	"os"

	"github.com/jmervine/getdown/Godeps/_workspace/src/gopkg.in/jmervine/ll.v1"

	"github.com/jmervine/getdown/cli"
	"github.com/jmervine/getdown/server"
)

var cfg = cli.Parse(os.Args)

func main() {
	if cfg.Addr == "" { // will be empty on "help" or "version"
		os.Exit(0)
	}

	svr := server.Bind(cfg)
	ll.Log(nil, map[string]interface{}{
		"at":   "startup",
		"addr": cfg.Addr,
		"port": cfg.Port,
	})

	ll.Debug(nil, map[string]interface{}{
		"at":       "startup",
		"basedir":  cfg.Basedir,
		"index":    cfg.Index,
		"template": cfg.Template,
		"title":    cfg.Title,
		"style":    cfg.Style,
	})

	ll.Logger.Fatal(svr.ListenAndServe())
}
