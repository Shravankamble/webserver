package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func table(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	fmt.Fprintf(w, "Table of given Integer : ")
	fmt.Fprintln(w, "   ")
	fmt.Println()
	num := r.FormValue("Integer")
	Int, err := strconv.Atoi(num)
	if err != nil {
		panic(err)
	}
	for i := 1; i < 11; i++ {
		ans := Int * i
		fmt.Fprintf(w, "%v %v %v %v %v\n", num, "x", i, "=", strconv.Itoa(ans))
	}
}

func content(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "PasrsForm() err : %v", err)
		return
	}
	// fmt.Fprintf(w, "ANSWER : ")
	number1 := r.FormValue("number1")
	number2 := r.FormValue("number2")
	num1, err := strconv.Atoi(number1)
	if err != nil {
		panic(err)
	}
	num2, err := strconv.Atoi(number2)
	if err != nil {
		panic(err)
	}
	sum := num1 + num2
	add := strconv.Itoa(sum)
	// publish := sum
	// fmt.Fprintf(w, "num1 : %v\n", number1)
	// fmt.Fprintf(w, "num2 : %v\n", number2)
	fmt.Fprintf(w, "the sum of given two numbers is : %v\n", add)
}

func forms(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "ParseForm() err : %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful! ")
	fmt.Println()
	name := r.FormValue("name")
	surname := r.FormValue("surname")
	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "SurName : %s\n", surname)
}

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "404 Not Found", http.StatusNotFound)
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
	http.HandleFunc("/content", content)
	http.HandleFunc("/table", table)

	fmt.Printf("starting server at port 8080 : ")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
