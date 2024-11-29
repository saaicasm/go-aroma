package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		_, err := fmt.Fprintf(w, "Hello World")

		if err != nil {
			log.Println("Error", err)
		}

	})

	http.ListenAndServe(":8080", nil)

}
