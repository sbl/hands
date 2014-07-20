package hands

import (
	"fmt"
	"net/http"
	"time"
)

const headerRuntime = "X-Runtime"

// Runtime measures the runtime of a request and sets up a corresponding
// X-Runtimte header.
func Runtime(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		if w.Header().Get(headerRuntime) == "" {
			dur := time.Since(start)
			w.Header().Set(headerRuntime, fmt.Sprintf("%0.6f", dur.Seconds()))
		}
	}
	return http.HandlerFunc(fn)
}
