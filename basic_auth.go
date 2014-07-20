package hands

import (
	"encoding/base64"
	"net/http"
)

const authenticateHeader = "WWW-Authenticate"

// BasicAuth provides http basic authentication for a given username, password
// combination.
func BasicAuth(username, password, realm string, next http.Handler) http.Handler {
	secret := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	fn := func(w http.ResponseWriter, r *http.Request) {
		if validAuth(secret, r.Header.Get("Authorization")) {
			next.ServeHTTP(w, r)
		} else {
			w.Header().Set(authenticateHeader, `Basic realm="`+realm+`"`)
			w.WriteHeader(http.StatusUnauthorized)
		}
	}
	return http.HandlerFunc(fn)
}

func validAuth(secret string, authHead string) bool {
	if len(authHead) < 6 {
		return false
	}

	auth := authHead[6:]

	if auth != secret {
		return false
	}
	return true
}
