package cli

import (
	//"os"
	"path"

	"github.com/jmervine/getdown/config"

	"github.com/jmervine/getdown/Godeps/_workspace/src/gopkg.in/codegangsta/cli.v1"
)

// slightly modified version of
// https://github.com/codegangsta/cli/blob/v1.2.0/help.go#L13
var AppHelpTemplate = `Name:
    {{.Name}} - {{.Usage}}

Usage:
    {{.Usage}} [args...]

Version:
    {{.Version}}

Options:
    {{range .Flags}}{{.}}
    {{end}}
`

func Parse(args []string) *config.Config {
	// use custom help template
	cli.AppHelpTemplate = AppHelpTemplate

	app := cli.NewApp()

	app.Version = "0.0.1"
	app.Name = "getdown"
	app.Usage = "Markdown file parser and server written in Go and designed to run on Docker."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "addr, a",
			Value:  config.Default.Addr,
			Usage:  "listener address",
			EnvVar: "GETDOWN_ADDR",
		},
		cli.StringFlag{
			Name:   "port, p",
			Value:  config.Default.Port,
			Usage:  "listener port",
			EnvVar: "GETDOWN_PORT",
		},
		cli.StringFlag{
			Name:   "basedir, b",
			Value:  config.Default.Basedir,
			Usage:  "base markdown directory",
			EnvVar: "GETDOWN_BASEDIR",
		},
		cli.StringFlag{
			Name:   "static, s",
			Value:  path.Join(config.Default.Basedir, "static"),
			Usage:  "static file directory",
			EnvVar: "GETDOWN_STATIC",
		},
		cli.StringFlag{
			Name:   "index, i",
			Value:  config.Default.Index,
			Usage:  "root file default",
			EnvVar: "GETDOWN_INDEX",
		},
		cli.StringFlag{
			Name:   "template, T",
			Value:  config.Default.Template,
			Usage:  "path to html template",
			EnvVar: "GETDOWN_TEMPLATE",
		},
		cli.StringFlag{
			Name:   "title, t",
			Value:  config.Default.Title,
			Usage:  "rendered page title/header for default template",
			EnvVar: "GETDOWN_TITLE",
		},
		cli.StringFlag{
			Name:   "style",
			Value:  config.Default.Style,
			Usage:  "bootstrap compatable stylesheet url",
			EnvVar: "GETDOWN_STYLE",
		},
	}

	var cfg config.Config
	app.Action = func(c *cli.Context) {
		cfg = config.New(
			c.String("addr"),
			c.String("port"),
			c.String("basedir"),
			c.String("public"),
			c.String("index"),
			c.String("style"),
			c.String("title"),
			c.String("template"),
		)
	}

	app.Run(args)

	return &cfg
}
