package hands

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRequestIDAssignRandomUUID(t *testing.T) {
	req, _ := http.NewRequest("GET", "", nil)
	rec := httptest.NewRecorder()

	c := new(counter)
	h := RequestID(c)
	h.ServeHTTP(rec, req)

	id := req.Header.Get(HeaderRequestID)
	if exp, got := 36, len(id); exp != got {
		t.Errorf("expected id with %d chars, got %d", exp, got)
	}

	if got := rec.Header().Get(HeaderRequestID); id != got {
		t.Errorf("expected same request-id on req and resp, got %#v", got)
	}

	req, _ = http.NewRequest("GET", "", nil)
	h.ServeHTTP(rec, req)
	if id == req.Header.Get(HeaderRequestID) {
		t.Errorf("expected alternating ids, again got: %#v", id)
	}
}

func TestRequestIDPassThruWhenPresent(t *testing.T) {
	id := "my-personal-very-random-request-id"

	req, _ := http.NewRequest("GET", "", nil)
	req.Header.Set(HeaderRequestID, id)
	rec := httptest.NewRecorder()

	c := new(counter)
	h := RequestID(c)
	h.ServeHTTP(rec, req)

	if got := rec.Header().Get(HeaderRequestID); id != got {
		t.Errorf("expected response request-id %#v, got %#v", id, got)
	}

	if got := req.Header.Get(HeaderRequestID); id != got {
		t.Errorf("expected request request-id %#v, got %#v", id, got)
	}
}
