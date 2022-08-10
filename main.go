package main

import (
	"fmt"
	"log"
	"net/http"
)

func forms(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful")
	fmt.Println()
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "SurName : %s\n", surname)
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "method is not supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "hello")
}

func main() {
	fileserver := http.FileServer(http.Dir("./style"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", forms)
	http.HandleFunc("/hello", hello)

	fmt.Printf("starting server at port 8080 : ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
