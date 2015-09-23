package hands_test

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sbl/hands"
)

func TestGzipDecompress(t *testing.T) {
	buf := new(bytes.Buffer)
	gz := gzip.NewWriter(buf)
	gz.Write([]byte("howdy"))
	gz.Flush()
	gz.Close()

	req, _ := http.NewRequest("POST", "/", buf)
	req.Header.Set(hands.HeaderContentEncoding, "gzip")
	rec := httptest.NewRecorder()

	var have []byte
	verifier := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		have, _ = ioutil.ReadAll(r.Body)
	})

	h := hands.Degzip(verifier)
	h.ServeHTTP(rec, req)

	if rec.Header().Get(hands.HeaderContentEncoding) != "" {
		t.Error("expected content-encoding to be unset")
	}

	if string(have) != "howdy" {
		t.Errorf("expected decoded output\n got: %+v", string(have))
	}
}

func TestGzipNoHeader(t *testing.T) {
	buf := new(bytes.Buffer)
	gz := gzip.NewWriter(buf)
	gz.Write([]byte("howdy"))
	gz.Flush()
	gz.Close()

	req, _ := http.NewRequest("POST", "/", buf)
	rec := httptest.NewRecorder()

	var have []byte
	verifier := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		have, _ = ioutil.ReadAll(r.Body)
	})

	h := hands.Degzip(verifier)
	h.ServeHTTP(rec, req)

	if string(have) == "howdy" {
		t.Errorf("expected undecoded output\n got: %+v", string(have))
	}
}
