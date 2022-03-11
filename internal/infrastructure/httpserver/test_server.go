package httpserver

import (
	"net/http"
	"net/http/httptest"
)

func NewTestServer(handler http.Handler) *httptest.Server {
	return httptest.NewServer(handler)
}
