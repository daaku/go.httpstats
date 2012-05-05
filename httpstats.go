// Package provides a HTTP handler that will keep track of some useful
// request statistics and expose them via expvar.
package httpstats

import (
	"expvar"
	"net/http"
)

type Handler struct {
	Handler       http.Handler
	Name          string
	totalRequests *expvar.Int
}

func NewHandler(name string, handler http.Handler) *Handler {
	return &Handler{
		Handler:       handler,
		Name:          name,
		totalRequests: expvar.NewInt(name + "_total_requests"),
	}
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.totalRequests.Add(1)
	h.Handler.ServeHTTP(w, r)
}
