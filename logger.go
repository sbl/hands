package hands

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"net/http"
)

// Logger provides a very simple logger in l2met format.
func Logger(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		wr := newWrapper(w)
		next.ServeHTTP(wr, r)

		str := fmt.Sprintf("method=%s path=%+v host=%+v code=%d bytes=%d",
			r.Method,
			r.URL.Path,
			r.Host,
			wr.status,
			wr.bytes,
		)
		if reqID := wr.Header().Get(HeaderRequestID); reqID != "" {
			str = fmt.Sprintf("%s request_id=%s",
				str,
				reqID,
			)
		}
		log.Print(str)
	}
	return http.HandlerFunc(fn)
}

// wrapper allows further introspection of generated responses.
type wrapper struct {
	http.ResponseWriter
	wroteHeader bool
	status      int
	bytes       int
}

func newWrapper(w http.ResponseWriter) *wrapper {
	return &wrapper{ResponseWriter: w}
}

func (wr *wrapper) WriteHeader(code int) {
	if !wr.wroteHeader {
		wr.ResponseWriter.WriteHeader(code)
		wr.wroteHeader = true
		wr.status = code
	}
}

func (wr *wrapper) Write(data []byte) (int, error) {
	wr.WriteHeader(http.StatusOK)
	n, err := wr.ResponseWriter.Write(data)
	wr.bytes += n
	return n, err
}

// Flush implements http.Flusher.
func (wr *wrapper) Flush() {
	if f, ok := wr.ResponseWriter.(http.Flusher); ok {
		f.Flush()
	}
}

// CloseNotify implements http.CloseNotifier.
func (wr *wrapper) CloseNotify() <-chan bool {
	if cn, ok := wr.ResponseWriter.(http.CloseNotifier); ok {
		return cn.CloseNotify()
	}
	return nil
}

// Hijack implements http.Hijacker.
func (wr *wrapper) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	if hj, ok := wr.ResponseWriter.(http.Hijacker); ok {
		return hj.Hijack()
	}
	return nil, nil, nil
}
