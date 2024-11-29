package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

const port = ":3000"

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

// Function handler for Divide
func Divide(w http.ResponseWriter, r *http.Request) {
	f, err := divideValues(100.0, 0.0)

	if err != nil {
		fmt.Fprintf(w, fmt.Sprintf("The error is : %s", err))
	} else {
		fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is equal to %f", 100.0, 0.0, f))
	}

}

// Adds two intergers
func addValues(x, y int) int {
	return x + y
}

// Divides two numbers
func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		err := errors.New("Cannot Divide by 0")
		return 0.0, err
	}

	result := x / y
	return result, nil
}

func main() {
	//Routes for Home and About
	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)
	http.HandleFunc("/Divide", Divide)

	//Starts the server
	log.Println(fmt.Sprintf("The server is running on Port %s", port))
	_ = http.ListenAndServe(port, nil)
}
