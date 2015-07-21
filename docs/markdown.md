# markdown
--
    import "github.com/jmervine/getdown/markdown"


## Usage

```go
var ValidExt = []string{".md", ".mdown"}
```

#### func  IsMarkdown

```go
func IsMarkdown(p string) bool
```

#### type Markdown

```go
type Markdown struct {
	Path     string
	Markdown []byte
	Body     string
}
```


#### func  New

```go
func New(file string, cfg *config.Config) (md Markdown, err error)
```

#### func (Markdown) IsValid

```go
func (md Markdown) IsValid() bool
```
