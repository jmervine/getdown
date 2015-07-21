# ll

short for "linelogger"

## [Read the Docs](https://godoc.org/github.com/jmervine/ll)

## Example

```go
package main
import (
    "gopkg.in/jmervine/ll.v1"
    "time"
)

func main() {
    begin := time.Now()

    // do stuff
    ll.Log(&begin, map[string]interface{}{
        "at": "main",
        "foo": "bar",
    })
}
```

*outputs*

```
2015/07/21 04:57:22 level=info at=main foo=bar durration=102.972us
```

