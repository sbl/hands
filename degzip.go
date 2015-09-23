package hands

import (
	"compress/gzip"
	"net/http"
)

const (
	HeaderContentType     = "Content-Type"
	HeaderContentEncoding = "Content-Encoding"
)

// Degzip incoming requests.
func Degzip(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if isGzipped(r) {
			gz, err := gzip.NewReader(r.Body)
			defer gz.Close()
			if err != nil {
				http.Error(w, "error reading content body", http.StatusBadRequest)
				return
			}
			r.Header.Del(HeaderContentEncoding)
			r.Body = gz
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func isGzipped(r *http.Request) bool {
	return r.Body != nil && r.Header.Get(HeaderContentEncoding) == "gzip"
}
