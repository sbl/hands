package hands

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestSetRuntime(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	rec := httptest.NewRecorder()
	c := new(counter)
	h := Runtime(c)
	h.ServeHTTP(rec, req)

	run := rec.Header().Get(headerRuntime)

	flreg := regexp.MustCompile(`^\d\.+`)
	if !flreg.MatchString(run) {
		t.Errorf("expected a float got %#v", run)
	}
}

func TestDontOverwrite(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	rec := httptest.NewRecorder()
	rec.Header().Set(headerRuntime, "foo")
	c := new(counter)
	h := Runtime(c)
	h.ServeHTTP(rec, req)

	run := rec.Header().Get(headerRuntime)

	if run != "foo" {
		t.Errorf("should not overwrite got %#v", run)
	}
}
