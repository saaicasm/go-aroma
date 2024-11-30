package render

import (
	"bytes"
	"github/iamlego/go-web/pkg/config"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)


var app *config.AppConfig

func NewTemplates(a *config.AppConfig) {
	app = a
}

func RenderTemplate(w http.ResponseWriter, tmpl string) {
	var tc map[string]*template.Template
	if app.UseCache {
		// get template cache from app config
		tc = app.TemplateCache

	}else{
		tc, _ = CreateTemplateCache()
	}


	

	//get the template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal("Error getting template")
	}

	buf := new(bytes.Buffer)

	err := t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func CreateTemplateCache() (map[string]*template.Template, error) {
	//will take no args
	//will return map of string -> template , error


	myCache := map[string]*template.Template{}

	pages, err := filepath.Glob("../.././templates/*.page.tmpl")


	if err != nil {
		return myCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../.././templates/*layout.tmpl")
		if err != nil {
			return myCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../.././templates/*layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}


	return myCache, nil

}
