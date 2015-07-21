package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmervine/getdown/Godeps/_workspace/src/gopkg.in/jmervine/ll.v1"

	"github.com/jmervine/getdown/config"
	"github.com/jmervine/getdown/markdown"
	"github.com/jmervine/getdown/template"
)

// Bind http server.
func Bind(cfg *config.Config) *http.Server {
	ll.Debug(nil, map[string]interface{}{
		"at":  "server.Bind",
		"cfg": fmt.Sprintf("%+v", *cfg),
	})

	handler := func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		begin := &now

		ll.Debug(nil, map[string]interface{}{
			"at":      "server.Bind #handler",
			"request": fmt.Sprintf("%+v", *r),
		})

		md, err := markdown.New(r.URL.Path, cfg)
		if err != nil {
			// serve static
			// see: server/file_server.go
			FileServerWithLogger(begin, cfg).ServeHTTP(w, r)
			return
		}

		payload, err := template.NewPayload(cfg, &md)
		if err != nil {
			// handle error
			ll.Log(begin, map[string]interface{}{
				"at":     "request",
				"method": r.Method,
				"status": 500,
				"url":    r.URL.Path,
				"error":  err,
			})
			return
		}

		payload.Render(w)
		ll.Log(begin, map[string]interface{}{
			"at":     "request",
			"method": r.Method,
			"status": 200,
			"url":    r.URL.Path,
			"source": md.Path,
		})
	}

	svr := &http.Server{
		Addr:    cfg.Listener(),
		Handler: http.HandlerFunc(handler),
	}

	ll.Debug(nil, map[string]interface{}{
		"at":  "server.Bind",
		"svr": fmt.Sprintf("%+v", *svr),
	})

	return svr
}
