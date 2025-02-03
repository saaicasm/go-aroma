package main

import (
	"fmt"
	"github/saaicasm/snipbox/internal/validator"
	"net/http"
	"strconv"
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

	var form SnippetCreateForm

	err := app.decodePostForm(r, &form)

	fmt.Println(form.Title)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	fmt.Println(validator.NotBlank(form.Title))

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

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

	app.sessionManager.Put(r.Context(), "flash", fmt.Sprintf("Snippet %d was created", id))

	http.Redirect(w, r, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)

}
