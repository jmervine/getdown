package template

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jmervine/getdown/config"
	"github.com/jmervine/getdown/markdown"
)

// wrapper for additional template handling later
func NewTemplate(file string) (*template.Template, error) {
	return template.ParseFiles(file)
}

// Payload is the type definition for the rendered template.
type Payload struct {
	template *template.Template
	markdown *markdown.Markdown
	Title    string
	Style    string
	Body     string
	Files    map[string][]string
}

func NewPayload(cfg *config.Config, md *markdown.Markdown) (payload Payload, err error) {
	tmpl, err := NewTemplate(cfg.Template)
	if err != nil {
		return
	}

	payload = Payload{
		template: tmpl,
		markdown: md,
		Title:    cfg.Title,
		Style:    cfg.Style,
		Body:     md.Body,
		Files:    make(map[string][]string),
	}

	walk := func(p string, f os.FileInfo, err error) error {
		if markdown.IsMarkdown(p) {
			p = strings.TrimPrefix(p, cfg.Basedir)
			dir, name := path.Split(p)
			payload.Files[dir] = append(payload.Files[dir], name)
		}
		return nil
	}

	filepath.Walk(cfg.Basedir, walk)

	return
}

func (payload Payload) Render(w io.Writer) {
	payload.template.ExecuteTemplate(w, payload.markdown.Path, payload)
}
