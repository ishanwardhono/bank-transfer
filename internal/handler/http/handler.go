package http

import (
	"fmt"
	"net/http"
)

func testFunc(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	w.Write([]byte(fmt.Sprint("okee")))
}
