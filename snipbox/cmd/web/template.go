package main

import (
	"github/saaicasm/snipbox/internal/models"
	"path/filepath"
	"text/template"
	"time"
)

type templateData struct {
	CurrentYear int
	Snippet     models.Snippet
	Snippets    []models.Snippet
	Form        any
}

func readableDate(t time.Time) string {
	return t.Format("02 Jan 2006 at 15:04")
}

var functions = template.FuncMap{
	"readableDate": readableDate,
}

func createTemplateCache() (map[string]*template.Template, error) {

	cache := map[string]*template.Template{}

	pages, err := filepath.Glob("internal/ui/html/pages/*.tmpl")

	if err != nil {
		return nil, err
	}

	for _, page := range pages {

		name := filepath.Base(page)

		files := []string{
			"internal/ui/html/pages/base.tmpl",
			"internal/ui/html/pages/nav.tmpl",
			page,
		}

		ts, err := template.New(name).Funcs(functions).ParseFiles(files...)

		if err != nil {
			return nil, err
		}

		cache[name] = ts

	}

	return cache, nil

}
