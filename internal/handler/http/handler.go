package http

import (
	"fmt"
	"net/http"
	"time"
)

func testFunc(w http.ResponseWriter, r *http.Request) {
	// ctx := r.Context()

	time.Sleep(10 * time.Second)

	w.Write([]byte(fmt.Sprint("okee")))
}
