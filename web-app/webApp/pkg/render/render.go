package render

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// var tc = make(map[string]*template.Template)

// func RenderTemplate(w http.ResponseWriter, t string) {
// 	var tmpl *template.Template
// 	var err error

// 	_, inMap := tc[t]
// 	if !inMap {
// 		// need to create the template
// 		log.Println("creating template and adding to cache")
// 		err = createTemplateCache(t)
// 		if err != nil {
// 			log.Println(err)
// 		}
// 	} else {

// 		log.Println("using cached template")
// 	}

// 	tmpl = tc[t]

// 	err = tmpl.Execute(w, nil)
// 	if err != nil {
// 		log.Println(err)
// 	}
// }

// func createTemplateCache(t string) error {
// 	templates := []string{
// 		fmt.Sprintf("./templates/%s", t),
// 		"./templates/base.layout.tmpl",
// 	}

// 	// parse the template
// 	tmpl, err := template.ParseFiles(templates...)
// 	if err != nil {
// 		return err
// 	}

// 	tc[t] = tmpl

// 	return nil
// }

func RenderTemplate(w http.ResponseWriter, tmpl string) {

	log.Println("This is the tmpl : ", tmpl)
	//create a template
	tc, err := createTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	//get the template from cache
	t, ok := tc[tmpl]

	if !ok {
		log.Fatal(err)
	}

	buf := new(bytes.Buffer)

	err = t.Execute(buf, nil)
	if err != nil {
		log.Println(err)
	}

	//render the template
	_, err = buf.WriteTo(w)
	if err != nil {
		log.Println(err)
	}

}

func createTemplateCache() (map[string]*template.Template, error) {
	//will take no args
	//will return map of string -> template , error

	//1 - create an empty map of string -> template
	myCache := map[string]*template.Template{}

	log.Println("This should be empty cache ", myCache)

	//2 - get all the pages with file named *.page.tmpl from ./templates
	pages, err := filepath.Glob("../.././templates/*.page.tmpl")

	log.Println("This should be all names of page.tmpls ", pages, err)

	if err != nil {
		return myCache, err
	}

	//3 - range through all the files ending with *.page.tmpl
	for _, page := range pages {
		name := filepath.Base(page)

		log.Println("The name of the page ", page, "is ", name)

		ts, err := template.New(name).ParseFiles(page)
		if err != nil {
			return myCache, err
		}

		matches, err := filepath.Glob("../.././templates/*layout.tmpl")
		if err != nil {
			return myCache, err
		}

		log.Println("The strings that matches layouts ", matches)

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("../.././templates/*layout.tmpl")
			if err != nil {
				return myCache, err
			}
		}

		myCache[name] = ts

	}

	log.Println("My cache at the end", myCache)

	return myCache, nil

}
