package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = ":8080"

// Function handler for Home
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is Home Page")
}

// Function handler for About
func About(w http.ResponseWriter, r *http.Request) {
	var sum int
	sum = addValues(6, 9)

	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the About Page and 6 + 9 is %d ", sum))

}

// Adds two intergers
func addValues(x int, y int) int {
	return x + y
}

func main() {

	//Routes for Home and About

	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)

	//Starts the Server on the port provided
	log.Println(fmt.Sprintf("The server is running on Port %s", port))
	_ = http.ListenAndServe(port, nil)
}
