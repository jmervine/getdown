package markdown

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path"

	"github.com/jmervine/getdown/config"
	"github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v1"
)

type Markdown struct {
	Path     string
	Markdown []byte
	Body     string
}

var ValidExt = []string{".md", ".mdown"}

func IsMarkdown(p string) bool {
	for _, ext := range ValidExt {
		if path.Ext(p) == ext {
			return true
		}
	}
	return false
}

func (md Markdown) IsValid() bool {
	return IsMarkdown(md.Path)
}

func New(file string, cfg *config.Config) (md Markdown, err error) {
	if path.Ext(file) == "" {
		file = path.Join(file, cfg.Index)
	}

	md = Markdown{Path: path.Join(cfg.Basedir, file)}

	if !md.IsValid() {
		return md, errors.New(fmt.Sprintf("invalid extension: %s", path.Ext(md.Path)))
	}

	md.Markdown, err = ioutil.ReadFile(md.Path)
	md.Body = string(bluemonday.UGCPolicy().SanitizeBytes(blackfriday.MarkdownCommon(md.Markdown)))

	return
}
