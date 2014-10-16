package hands

import (
	"fmt"
	"net/http"
)

type counter int

func (c *counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	(*c)++
	fmt.Fprint(w, c)
}
