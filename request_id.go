package hands

import (
	"net/http"

	"code.google.com/p/go-uuid/uuid"
)

const headerRequestID = "X-Request-ID"

// RequestID sets up "X-Request-Id" header if it has not been previously been
// set up.
func RequestID(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		id := r.Header.Get(headerRequestID)
		if id == "" {
			id = uuid.New()
		}
		w.Header().Set(headerRequestID, id)
		r.Header.Set(headerRequestID, id)

		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
