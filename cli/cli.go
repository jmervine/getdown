package cli

import (
	"gopkg.in/codegangsta/cli.v1"
	"os"

	"github.com/jmervine/getdown/config"
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
		cfg = config.Config{
			Addr:     c.String("addr"),
			Port:     c.String("port"),
			Basedir:  c.String("basedir"),
			Index:    c.String("index"),
			Template: c.String("template"),
			Title:    c.String("title"),
			Style:    c.String("style"),
		}

		if cfg.Addr == "" || cfg.Port == "" || cfg.Basedir == "" || cfg.Index == "" {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
	}

	app.Run(args)

	return &cfg
}
