package main

import (
	"fmt"
	"github/saaicasm/snipbox/internal/validator"
	"net/http"
	"strconv"
	"strings"
	"unicode/utf8"
)

type SnippetCreateForm struct {
	Title               string `form:"title"`
	Content             string `form:"content"`
	Expires             int    `form:"expires"`
	validator.Validator `form:"-"`
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

	var form snippetCreateForm

	err = app.formDecoder.Decode(&form, r.PostForm)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	expires, err := strconv.Atoi(r.PostForm.Get("expires"))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	if strings.TrimSpace(form.Title) == "" {
		form.FieldErrors["title"] = "The title is Empty"
	} else if utf8.RuneCountInString(form.Content) > 100 {
		form.FieldErrors["content"] = "The content is too long must be less than 100 characters"
	}

	if strings.TrimSpace("content") == "" {
		form.FieldErrors["content"] = "The content is empty"
	}

	if expires != 1 && expires != 7 && expires != 365 {
		form.FieldErrors["expires"] = "Enter a valid expiry date"
	}

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "create.tmpl", data)
		return
	}

	id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)

	if err != nil {
		app.serverError(w, r, err)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}
