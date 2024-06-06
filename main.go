package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprint(w, "ParseForm() err: ", err)
		return
	}
	fmt.Fprint(w, "POST request sucessfull\n")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprint(w, "Name =  ", name)
	fmt.Fprint(w, " \n")
	fmt.Fprint(w, "Address =  ", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "hello")
}

func main() {
	fileServe := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServe)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
