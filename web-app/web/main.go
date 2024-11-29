package main

import (
	"fmt"
	"log"
	"net/http"
)
const port = ":3000"
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/About", About)

	log.Println(fmt.Sprintf("The server is running on Port %s", port))
	_ = http.ListenAndServe(port, nil)
}
