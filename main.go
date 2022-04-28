package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {

	fmt.Fprint(w, "Hello, world Vai Curintia!!")

}

func main() {
	http.HandleFunc("/", homePage)

	log.Println("Starting web server at port 8080")
	http.ListenAndServe(":8080", nil)

}
