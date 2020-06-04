package http

import (
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"net/http"
)

func TracedHandler(h http.Handler) http.Handler {
	return wrapHandler(h)
}

func wrapHandler(next http.Handler) http.Handler {
	return &ochttp.Handler{
		Handler: http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				if span := trace.FromContext(r.Context()); span != nil {
					span.AddAttributes(
						serverDefaultAttributes()...,
					)
				}
				next.ServeHTTP(w, r)
			},
		)}
}
