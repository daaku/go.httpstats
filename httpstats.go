// Package httpstats provides a HTTP handler that will keep track of
// some useful request statistics.
package httpstats

import (
	"net/http"
	"time"
)

type Stats interface {
	Inc(name string)
	Record(name string, value float64)
}

type Handler struct {
	Handler http.Handler
	Name    string
	Stats   Stats
}

func NewHandler(name string, handler http.Handler) *Handler {
	return &Handler{
		Handler: handler,
		Name:    name,
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Stats.Inc("web request")
	h.Stats.Inc("web request - method=" + r.Method)
	start := time.Now()
	h.Handler.ServeHTTP(w, r)
	h.Stats.Record("web request gen time", float64(time.Since(start).Nanoseconds()))
}
