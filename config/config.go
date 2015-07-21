package config

import (
	"fmt"
	"path"
)

var Default = Config{
	Addr:     "localhost",
	Port:     "3000",
	Basedir:  ".",
	Index:    "README.md",
	Style:    "https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css",
	Title:    "getdown",
	Template: "./template.html",
}

type Config struct {
	Addr     string
	Port     string
	Basedir  string
	Static   string
	Index    string
	Style    string
	Title    string
	Template string
}

func New(addr, port, base, static, index, style, title, tmpl string) Config {
	cfg := *(&Default) // clone Default

	if addr != "" {
		cfg.Addr = addr
	}

	if port != "" {
		cfg.Port = port
	}

	if base != "" {
		cfg.Basedir = base
	}

	// todo:
	// - starts with "/" || "./" == static
	// - starts with "[a-z]" == join(base, static)
	if static != "" {
		cfg.Static = static
	} else {
		cfg.Static = path.Join(cfg.Basedir, "static")
	}

	if index != "" {
		cfg.Index = index
	}

	if title != "" {
		cfg.Title = title
	}

	if tmpl != "" {
		cfg.Template = tmpl
	}

	return cfg
}

func (c Config) Listener() string {
	return fmt.Sprintf("%s:%s", c.Addr, c.Port)
}
