package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
    // Pobieramy port z argumentu, np. "go run server.go 8081"
	port := "8081" 
	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Otrzymano zapytanie na porcie %s\n", port)
		fmt.Fprintf(w, "Cześć! Odpowiedź z Backend Server na porcie: %s", port)
	})

	log.Printf("Startowanie serwera backendowego na porcie: %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}