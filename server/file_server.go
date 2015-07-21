package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/jmervine/getdown/Godeps/_workspace/src/gopkg.in/jmervine/ll.v1"

	"github.com/jmervine/getdown/config"
)

type hijackFileServer struct {
	http.ResponseWriter
	R    *http.Request
	Code int
}

func (h *hijackFileServer) WriteHeader(code int) {
	if code != 200 {
		h.Code = code
		panic(h)
	} else {
		h.ResponseWriter.WriteHeader(code)
	}
}

// FileServerWithLogger does what you would expect.
func FileServerWithLogger(begin *time.Time, cfg *config.Config) http.HandlerFunc {
	handler := func(w http.ResponseWriter, r *http.Request) {
		hijack := &hijackFileServer{ResponseWriter: w, R: r}
		var status int
		defer func() {
			if p := recover(); p != nil {
				if p == hijack {
					status = p.(*hijackFileServer).Code
				} else {
					panic(p)
				}
			}

			ll.Log(begin, map[string]interface{}{
				"at":     "request",
				"method": r.Method,
				"status": status,
				"url":    r.URL.Path,
				"source": "static",
			})

			if status != 200 {
				hijack.ResponseWriter.WriteHeader(status)
				if status == 404 {
					hijack.ResponseWriter.Write([]byte("Not Found"))
				} else {
					hijack.ResponseWriter.Write([]byte("Unknown Error"))
				}
			}
		}()

		ll.Debug(begin, map[string]interface{}{
			"at":      "server.FileServerWithLogger #handler",
			"request": fmt.Sprintf("%+v", *r),
		})

		http.FileServer(http.Dir(cfg.Static)).ServeHTTP(w, r)
	}

	return http.HandlerFunc(handler)
}
