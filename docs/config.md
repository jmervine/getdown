# config
--
    import "github.com/jmervine/getdown/config"


## Usage

```go
var Default = Config{
	Addr:     "localhost",
	Port:     "3000",
	Basedir:  ".",
	Public:   "./public",
	Index:    "README.md",
	Template: "./template.html",
	Title:    "getdown",
	Style:    "https://maxcdn.bootstrapcdn.com/bootstrap/3.3.5/css/bootstrap.min.css",
}
```

#### type Config

```go
type Config struct {
	Port     string
	Addr     string
	Basedir  string
	Public   string
	Index    string
	Style    string
	Title    string
	Template string
}
```


#### func (Config) Listener

```go
func (c Config) Listener() string
```
