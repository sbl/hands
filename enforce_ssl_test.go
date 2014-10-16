package hands

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestEnforceSSL(t *testing.T) {
	emptyHeader := map[string][]string{}
	forwardHeader := map[string][]string{forwaredProto: {"https"}}

	table := []struct {
		method string
		header map[string][]string
		scheme string
		exp    int
	}{
		{"GET", emptyHeader, "https", 200},
		{"GET", forwardHeader, "http", 200},
		{"GET", forwardHeader, "https", 200},
		{"POST", emptyHeader, "https", 200},
		{"POST", forwardHeader, "https", 200},
		{"GET", emptyHeader, "http", 301},
		{"HEAD", emptyHeader, "http", 301},
		{"PUT", emptyHeader, "http", 307},
		{"PUT", emptyHeader, "http", 307},
	}

	for _, tt := range table {
		rec := httptest.NewRecorder()
		req := &http.Request{
			Method: tt.method,
			URL:    &url.URL{Scheme: tt.scheme, Path: "/foo"},
			Host:   "example.com",
			Header: tt.header,
		}

		var c counter
		h := EnforceSSL(&c)
		h.ServeHTTP(rec, req)

		if got := rec.Code; tt.exp != got {
			t.Errorf("expected %+v, got %+v", tt.exp, got)
		}

		// check location on redirect
		if exp, got := "https://example.com/foo", rec.Header().Get("Location"); rec.Code > 299 && exp != got {
			t.Errorf("expected %+v, got %+v", exp, got)
		}
	}
}
