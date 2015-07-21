# server
--
    import "github.com/jmervine/getdown/server"


## Usage

#### func  Bind

```go
func Bind(cfg *config.Config) *http.Server
```
Bind http server.

#### func  FileServerWithLogger

```go
func FileServerWithLogger(begin *time.Time, cfg *config.Config) http.HandlerFunc
```
FileServerWithLogger does what you would expect.
