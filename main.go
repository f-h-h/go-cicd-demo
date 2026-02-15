package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	version := os.Getenv("VERSION")
	if version == "" {
		version = "1.0.0"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from Go CI/CD Demo! Version: %s\n", version)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "ok")
	})

	log.Printf("Server starting on port %s, version %s", port, version)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
