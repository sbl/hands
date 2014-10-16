# hands [WIP]

[![Build
Status](https://travis-ci.org/sbl/hands.svg)](https://travis-ci.org/sbl/hands)
[![GoDoc](https://godoc.org/github.com/sbl/hands?status.svg)](https://godoc.org/github.com/sbl/hands)

helpful `net/http` conformant handlers.

- `hands.RequestID(next http.Handler)` add a request ID
- `hands.Runtime(next http.Handler)` measure runtime
- `hands.EnforceSSL(next http.Handler)` redirect to SSL
- `hands.BasicAuth(next http.Handler, user, password, realm string)`
  http basic auth
