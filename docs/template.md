# template
--
    import "github.com/jmervine/getdown/template"


## Usage

```go
var BlacklistPrefix = []string{".", "_", "~"}
```
BlacklistPrefix contains a list of values that will blacklist files and
directories when traversing the file tree.

Update / add:

    markdown.BlacklistPrefix = append(markdown.BlacklistPrefix, "X")

Replace:

    markdown.BlacklistPrefix = []string{"~"}

var BlacklistPrefix = []string{"Godeps", ".", "_", "~"}

#### func  NewTemplate

```go
func NewTemplate(file string) (*template.Template, error)
```
wrapper for additional template handling later

#### type Payload

```go
type Payload struct {
	Markdown *markdown.Markdown
	Title    string
	Style    string
	Body     string
	Files    map[string][]string
}
```

Payload is the type definition for the rendered template.

#### func  NewPayload

```go
func NewPayload(cfg *config.Config, md *markdown.Markdown) (payload Payload, err error)
```

#### func (Payload) Render

```go
func (payload Payload) Render(w io.Writer)
```
