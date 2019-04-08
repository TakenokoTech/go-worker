package handlers

import (
	"fmt"
	"net/http"
)

func SyncHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}
