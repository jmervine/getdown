# cli
--
    import "github.com/jmervine/getdown/cli"


## Usage

```go
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
```
slightly modified version of
https://github.com/codegangsta/cli/blob/v1.2.0/help.go#L13

#### func  Parse

```go
func Parse(args []string) *config.Config
```
