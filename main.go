package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Name struct {
	Name string `json:"name"`
}

func mainPage(w http.ResponseWriter, r *http.Request) { //func for main page
	fmt.Fprintf(w, "NEW Changes for checking CI/CD pipeline")
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	var n Name
	err := json.NewDecoder(r.Body).Decode(&n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Fprintf(w, "Hello, %s!", n.Name)
}

func routes() {
	http.HandleFunc("/", mainPage) //Added route for main page
	http.HandleFunc("/api/hello", hello)
}

func handleFunc() {
	routes()                                 //as you know, we can have many routes, i deceided to make func for this
	err := http.ListenAndServe(":8080", nil) // configuring the port
	log.Println("Server is running!")
	if err != nil { //excepting errors
		log.Fatal("ListenAndServe: ", err)
	}
}

func main() { //The entry point
	handleFunc()
}
