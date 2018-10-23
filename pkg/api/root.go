package api

import (
	"fmt"
	"net/http"
)

// RootHandler handles requests to /
func RootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hytta API is on its way! :)")
}