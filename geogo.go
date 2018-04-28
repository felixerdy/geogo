package main

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t, _ := template.ParseFiles("main.html")
		t.Execute(w, nil)
	})

	http.HandleFunc("/countryName", func(w http.ResponseWriter, r *http.Request) {
		result := executeQuery(r.URL.Query().Get("query"), schema)
		json.NewEncoder(w).Encode(result)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
