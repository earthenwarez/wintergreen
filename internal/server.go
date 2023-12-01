package server

import (
	"net/http"

	"log"

	"os"
)

// Start starts the server.
func Start() {
	HealthHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("healthy"))
	}
	ReadHandler := func(w http.ResponseWriter, r *http.Request) {
		data, err := os.ReadFile("data/data.txt")
		if err != nil {
			log.Fatal(err)
		}
		w.Write(data)
	}
	http.HandleFunc("/", HealthHandler)
	http.HandleFunc("/read", ReadHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
