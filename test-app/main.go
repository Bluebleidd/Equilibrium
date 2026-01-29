package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Hostname string
	Color    string
}

func main() {
	hostname, _ := os.Hostname()
	color := os.Getenv("APP_COLOR")
	if color == "" {
		color = "#3498db"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, err := template.ParseFiles("index.html")

		if err != nil {
			http.Error(w, "Błąd: Nie znaleziono pliku index.html", http.StatusInternalServerError)
			log.Println("Błąd ładowania szablonu:", err)
			return
		}

		t.Execute(w, PageData{Hostname: hostname, Color: color})
	})

	log.Println("Aplikacja testowa startuje na porcie 8081...")
	http.ListenAndServe(":8081", nil)
}
