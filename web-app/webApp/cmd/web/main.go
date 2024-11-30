package main

import (
	"fmt"
	"github/iamlego/go-web/pkg/config"
	handler "github/iamlego/go-web/pkg/handlers"
	"github/iamlego/go-web/pkg/render"
	"log"
	"net/http"
)

const port = ":3000"

func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cant create templ cache")
	}

	app.TemplateCache = tc
	app.UseCache = false

	repo := handler.NewRepo(&app) // get repo from handler and add app to it
	handler.NewHandler(repo)      // send back the updated repo with tc to handler

	render.NewTemplates(&app)

	http.HandleFunc("/", handler.Repo.Home)
	http.HandleFunc("/About", handler.Repo.About)

	log.Println(fmt.Sprintf("The server is running on Port %s", port))
	_ = http.ListenAndServe(port, nil)
}

// package main

// import (
// 	"log"
// 	"path/filepath"
// )

// func main() {

// 	pages, err := filepath.Glob("../.././templates/*.page.tmpl")

// 	log.Println(pages, err)

// }
