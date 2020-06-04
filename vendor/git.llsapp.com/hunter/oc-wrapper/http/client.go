package http

import (
	"go.opencensus.io/plugin/ochttp"
	"go.opencensus.io/trace"
	"net/http"
)

func NewTracedClient(remoteHost string, opts ...trace.StartOption) *http.Client {
	octransport := &ochttp.Transport{
		Base: newSpanAnnotator(http.DefaultTransport, remoteHost),
		GetStartOptions: func(*http.Request) trace.StartOptions {
			o := trace.StartOptions{SpanKind: trace.SpanKindClient}
			for _, opt := range opts {
				opt(&o)
			}
			return o
		},
	}
	return &http.Client{Transport: octransport}
}

type spanAnnotator struct {
	rt         http.RoundTripper
	remoteHost string
}

func newSpanAnnotator(rt http.RoundTripper, remoteHost string) http.RoundTripper {
	return &spanAnnotator{rt: rt, remoteHost: remoteHost}
}

func (s *spanAnnotator) RoundTrip(req *http.Request) (*http.Response, error) {
	if span := trace.FromContext(req.Context()); span != nil {
		span.AddAttributes(
			clientDefaultAttributes(s.remoteHost)...,
		)
	}
	return s.rt.RoundTrip(req)
}
