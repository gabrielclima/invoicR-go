package rest

import (
	"fmt"
	"net/http"
)

const TextHTML = "text/html; charset=UTF-8"

func ApiResource(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", TextHTML)
	fmt.Fprint(w, "Hello")
}
