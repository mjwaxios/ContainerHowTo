package main

import (
	"fmt"
	"log"
	"net/http"
)

var val int

func homePage(w http.ResponseWriter, r *http.Request) {
	val++
	fmt.Fprintf(w, "serv Test program\n %v", val)
}

func main() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":10100", nil))
}
