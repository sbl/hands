package hands

import "net/http"

const forwaredProto = "X-Forwarded-Proto"

// EnforceSSL redirects all "http" requests to "https".
func EnforceSSL(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.TLS == nil && r.Header.Get(forwaredProto) != "https" && r.URL.Scheme != "https" {
			r.URL.Scheme = "https"
			r.URL.Host = r.Host

			var status int
			if r.Method == "GET" || r.Method == "HEAD" {
				status = http.StatusMovedPermanently
			} else {
				status = http.StatusTemporaryRedirect
			}
			http.Redirect(w, r, r.URL.String(), status)
			return
		}
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
