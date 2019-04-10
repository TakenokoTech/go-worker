package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func SyncHandler(w http.ResponseWriter, r *http.Request) {
	go sync()
	w.WriteHeader(202)
}

func sync() {
	fmt.Println("sync")
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println(i)
	}
}
