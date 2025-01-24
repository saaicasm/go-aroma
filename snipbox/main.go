package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Let's Go!"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))

	if err != nil {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte(fmt.Sprintf("Display Id no %d", id)))

}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	inp := r.PathValue("query")
	fmt.Println("Query", inp)
	w.Write([]byte("Create a new Snippet!"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is Initial"))
	w.WriteHeader(http.StatusCreated)

	fmt.Println("The snippet was created!!")

}

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create/{query}", snippetCreate)
	mux.HandleFunc("POST /snippet/create/{action}", snippetCreatePost)

	log.Println("Server running on port 4000")
	err := http.ListenAndServe(":4000", mux)

	if err != nil {
		log.Fatal(err)
	}

}
