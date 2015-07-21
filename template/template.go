package template

import (
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jmervine/getdown/Godeps/_workspace/src/gopkg.in/jmervine/ll.v1"

	"github.com/jmervine/getdown/config"
	"github.com/jmervine/getdown/markdown"
)

// BlacklistPrefix contains a list of values that will blacklist files and
// directories when traversing the file tree.
//
// Update / add:
//
//     markdown.BlacklistPrefix = append(markdown.BlacklistPrefix, "X")
//
// Replace:
//
//     markdown.BlacklistPrefix = []string{"~"}
//var BlacklistPrefix = []string{"Godeps", ".", "_", "~"}
var BlacklistPrefix = []string{".", "_", "~"}

// wrapper for additional template handling later
func NewTemplate(file string) (*template.Template, error) {
	return template.ParseFiles(file)
}

// Payload is the type definition for the rendered template.
type Payload struct {
	template *template.Template
	Markdown *markdown.Markdown
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
		Markdown: md,
		Title:    cfg.Title,
		Style:    cfg.Style,
		Body:     md.Body,
		Files:    make(map[string][]string),
	}

	ll.Debug(nil, map[string]interface{}{
		"at":      "NewPayload",
		"payload": fmt.Sprintf("%+v", payload),
	})

	// Note: I hate this, there has to be a better way.

	walk := func(p string, f os.FileInfo, err error) error {
		if markdown.IsMarkdown(p) && !blacklist(p) {
			p = strings.TrimPrefix(p, cfg.Basedir)
			dir, name := path.Split(p)
			if !strings.HasPrefix(p, "/") {
				dir = "/" + dir
			}
			payload.Files[dir] = append(payload.Files[dir], name)
		}
		return nil
	}

	filepath.Walk(cfg.Basedir, walk)

	return
}

func (payload Payload) Render(w io.Writer) {
	payload.template.Execute(w, payload)
}

func blacklist(p string) bool {
	parts := strings.Split(p, "/")
	for _, b := range BlacklistPrefix {
		for _, p := range parts {
			if strings.HasPrefix(p, b) {
				return true
			}
		}
	}

	return false
}
