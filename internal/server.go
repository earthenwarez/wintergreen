package server

import (
	"net/http"

	"log"
)

// Start starts the server.
func Start() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	}
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
