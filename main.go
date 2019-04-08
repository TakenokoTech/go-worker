package main

import (
	"net/http"

	"github.com/TakenokoTech/go-worker/handlers"
)

func main() {
	http.HandleFunc("/", handlers.SyncHandler)
	http.ListenAndServe(":8080", nil)
}
