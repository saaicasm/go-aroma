package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

type SnippetCreateForm struct {
	Title       string
	Content     string
	Expires     int
	FieldErrors map[string]string
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	snippets, err := app.snippets.Latest()

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	td := app.newTemplateData(r)
	td.Snippets = snippets

	app.render(w, r, http.StatusOK, "home.tmpl", td)

}

func (app *application) snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}

	snippet, err := app.snippets.Get(id)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	td := app.newTemplateData(r)
	td.Snippet = snippet

	app.render(w, r, http.StatusOK, "view.tmpl", td)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	td := app.newTemplateData(r)
	td.Form = SnippetCreateForm{
		Expires: 365,
	}

	app.render(w, r, http.StatusOK, "create.tmpl", td)
}

func (app *application) snippetCreatePost(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}
	Form := SnippetCreateForm{
		Title:       r.PostForm.Get("title"),
		Content:     r.PostForm.Get("content"),
		Expires:     expires,
		FieldErrors: map[string]string{},
	}

	if strings.TrimSpace(Form.Title) == "" {
		Form.FieldErrors["title"] = "The title is Empty"
	} else if utf8.RuneCountInString(Form.Content) > 100 {
		Form.FieldErrors["content"] = "The content is too long must be less than 100 characters"
	}

	if strings.TrimSpace("content") == "" {
		Form.FieldErrors["content"] = "The content is empty"
	}

	if expires != 1 && expires != 7 && expires != 365 {
		Form.FieldErrors["expires"] = "Enter a valid expiry date"
	}

	if len(Form.FieldErrors) > 0 {
		td := templateData{
			Form: Form,
		}
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl", td)
		return
	}

	id, err := app.snippets.Insert(Form.Title, Form.Content, Form.Expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	log.Println(id)

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}
