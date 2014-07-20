package hands

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"time"
)

const headerRuntime = "X-Runtime"

// Runtime measures the runtime of a request and sets up a corresponding
// X-Runtimte header.
func Runtime(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := httptest.NewRecorder()
		next.ServeHTTP(rec, r)

		for k, v := range rec.Header() {
			w.Header()[k] = v
		}
		if w.Header().Get(headerRuntime) == "" {
			dur := time.Since(start)
			w.Header().Set(headerRuntime, fmt.Sprintf("%0.6f", dur.Seconds()))
		}
		w.WriteHeader(rec.Code)
		w.Write(rec.Body.Bytes())
	}
	return http.HandlerFunc(fn)
}
