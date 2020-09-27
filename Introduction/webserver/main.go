package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Listening to Port 8080"))
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
