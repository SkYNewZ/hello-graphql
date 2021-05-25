package main

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

var _ http.RoundTripper = (*transport)(nil)

type transport struct {
	Original http.RoundTripper
}

func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Debugf("URL being requested: %s\n", req.URL.String())
	return t.Original.RoundTrip(req)
}

// newHTTPClient returns an http.Client with our custom transport
func newHTTPClient() *http.Client {
	return &http.Client{Transport: &transport{http.DefaultTransport}}
}

// loggingMiddleware enables http access log
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.WithFields(log.Fields{
			"method":      r.Method,
			"request_uri": r.RequestURI,
			"user_agent":  r.UserAgent(),
		}).Printf("%s %s", r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
