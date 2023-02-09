package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainPage(w http.ResponseWriter, r *http.Request) { //func for main page
	fmt.Fprintf(w, "Привет!")
}

func routes() {
	http.HandleFunc("/", mainPage) //Added route for main page
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
