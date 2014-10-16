package hands

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuth(t *testing.T) {
	u := "foobie"
	p := "passie"
	realm := "super secret"

	table := []struct {
		u    string
		p    string
		code int
	}{
		{"", "", 401},
		{"simply", "wrong", 401},
		{u, "wrong-pass", 401},
		{"wrong-user", p, 401},
		{u, p, 200},
	}
	for _, tt := range table {
		req, _ := http.NewRequest("GET", "", nil)
		if tt.u != "" {
			req.SetBasicAuth(tt.u, tt.p)
		}
		rec := httptest.NewRecorder()

		var c counter
		h := BasicAuth(&c, u, p, realm)

		h.ServeHTTP(rec, req)

		if exp, got := tt.code, rec.Code; exp != got {
			t.Errorf("expected %+v, got %+v, for: %+v", exp, got, tt)
		}
		if rec.Code == 401 && rec.Header().Get(authenticateHeader) == "" {
			t.Errorf("requires %s header when no unauthorized", authenticateHeader)
		}
	}
}
