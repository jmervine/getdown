package cli

import (
	"fmt"
	"gopkg.in/codegangsta/cli.v1"
	"os"
)

type Config struct {
	Port    string
	Addr    string
	Basedir string
	Index   string
	Style   string
	Title   string
}

func (c Config) Listener() string {
	return fmt.Sprintf("%s:%s", c.Addr, c.Port)
}

func Parse(args []string) (config Config) {
	app := cli.NewApp()

	app.Version = "0.0.1"
	app.Name = "getdown"
	app.Usage = "Markdown file parser and server written in Go and designed to run on Docker."

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "addr, a",
			Value:  "localhost",
			Usage:  "listener address",
			EnvVar: "GETDOWN_ADDR",
		},
		cli.StringFlag{
			Name:   "port, p",
			Value:  "3000",
			Usage:  "listener port",
			EnvVar: "GETDOWN_PORT",
		},
		cli.StringFlag{
			Name:   "title, t",
			Value:  "getdown",
			Usage:  "rendered page title/header, pass empty string to use basedir name",
			EnvVar: "GETDOWN_TITLE",
		},
		cli.StringFlag{
			Name:   "basedir, b",
			Value:  ".",
			Usage:  "base markdown directory",
			EnvVar: "GETDOWN_BASEDIR",
		},
		cli.StringFlag{
			Name:   "index, i",
			Value:  "README.md",
			Usage:  "root file default",
			EnvVar: "GETDOWN_INDEX",
		},
		cli.StringFlag{
			Name:   "style",
			Value:  "https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css",
			Usage:  "bootstrap compatable stylesheet url",
			EnvVar: "GETDOWN_STYLE",
		},
	}

	app.Action = func(c *cli.Context) {
		config.Addr = c.String("addr")
		config.Port = c.String("port")
		config.Basedir = c.String("basedir")
		config.Index = c.String("index")
		config.Style = c.String("style")
		config.Title = c.String("title")

		if config.Addr == "" || config.Port == "" || config.Basedir == "" || config.Index == "" {
			cli.ShowAppHelp(c)
			os.Exit(1)
		}
	}

	app.Run(args)

	return
}
