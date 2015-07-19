package config

import (
	"fmt"
)

var Default = Config{
	Addr:     "localhost",
	Port:     "3000",
	Basedir:  ".",
	Index:    "README.md",
	Template: "./template.html",
	Title:    "getdown",
	Style:    "https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css",
}

type Config struct {
	Port     string
	Addr     string
	Basedir  string
	Index    string
	Style    string
	Title    string
	Template string
}

func (c Config) Listener() string {
	return fmt.Sprintf("%s:%s", c.Addr, c.Port)
}
