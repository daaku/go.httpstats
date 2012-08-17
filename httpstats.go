// Package httpstats provides a HTTP handler that will keep track of
// some useful request statistics and log them via go.stats.
package httpstats

import (
	"github.com/daaku/go.stats"
	"net/http"
	"time"
)

type Handler struct {
	Handler http.Handler
	Name    string
}

func NewHandler(name string, handler http.Handler) *Handler {
	return &Handler{
		Handler: handler,
		Name:    name,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	stats.Inc("web request")
	stats.Inc("web request - method=" + r.Method)
	start := time.Now()
	h.Handler.ServeHTTP(w, r)
	stats.Record("web request gen time", time.Since(start).Nanoseconds())
}
