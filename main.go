package main

import (
	"log"
	"net/http"
	"text/template"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html")
}

func playRound(w http.ResponseWriter, r *http.Request) {

}

// Main Loop
func main() {
	http.HandleFunc("/play", playRound)
	http.HandleFunc("/", homePage)
	log.Println("Starting web server at port 8080")
	http.ListenAndServe(":8080", nil)
}

// Generic Template Function
func renderTemplate(w http.ResponseWriter, page string) {
	t, err := template.ParseFiles(page)
	if err != nil {
		log.Println(err)
		return
	}
	err = t.Execute(w, nil)

	if err != nil {
		log.Println(err)
		return
	}

}
