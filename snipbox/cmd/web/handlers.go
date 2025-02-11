package main

import (
	"errors"
	"fmt"
	"github/saaicasm/snipbox/internal/models"
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

type UserSignUpForm struct {
	Name                string `form:"name"`
	Email               string `form:"email"`
	Password            string `form:"password"`
	validator.Validator `form:"-"`
}

type UserLoginForm struct {
	Email               string `form:"email"`
	Password            string `form:"password"`
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

	fmt.Println(td.IsAuthenticated)

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

	fmt.Println(td.IsAuthenticated)

	app.render(w, r, http.StatusOK, "view.tmpl", td)

}

func (app *application) snippetCreate(w http.ResponseWriter, r *http.Request) {

	td := app.newTemplateData(r)
	td.Form = SnippetCreateForm{
		Expires: 365,
	}

	fmt.Println(td.IsAuthenticated)

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

	form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")
	form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long")
	form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
	form.CheckField(validator.PermittedValue(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form

		fmt.Println(data.IsAuthenticated)

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

func (app *application) userSignup(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = UserSignUpForm{}

	fmt.Println(data.IsAuthenticated)

	app.render(w, r, http.StatusOK, "signup.tmpl", data)

}
func (app *application) UserSignUpPost(w http.ResponseWriter, r *http.Request) {
	var form UserSignUpForm

	err := app.decodePostForm(r, &form)

	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	fmt.Println(form.Name)

	form.CheckField(validator.NotBlank(form.Name), "name", "Field Name cannot be empty")
	form.CheckField(validator.NotBlank(form.Email), "email", "Field cannot be empty")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "Field needs to be valid email be empty")
	form.CheckField(validator.NotBlank(form.Password), "password", "Field cannot be empty")
	form.CheckField(validator.MinChars(form.Password, 8), "password", "Must be atleast 8 chars")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form

		fmt.Println(data.IsAuthenticated)

		app.render(w, r, http.StatusUnprocessableEntity, "signup.tmpl", data)
		return
	}

	err = app.users.Insert(form.Name, form.Email, form.Password)

	if err != nil {
		if errors.Is(err, models.ErrDuplicateEmail) {
			form.AddFieldError("email", "Email address is already is use")

			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, r, http.StatusUnprocessableEntity, "signup.tmpl", data)
		} else {
			app.serverError(w, r, err)
		}

		return
	}

	app.sessionManager.Put(r.Context(), "flash", "Your Signup was successful. Please log in")

	http.Redirect(w, r, "/user/login", http.StatusSeeOther)

}
func (app *application) UserLogin(w http.ResponseWriter, r *http.Request) {
	data := app.newTemplateData(r)
	data.Form = UserLoginForm{}
	app.render(w, r, http.StatusOK, "login.tmpl", data)
}
func (app *application) UserLoginPost(w http.ResponseWriter, r *http.Request) {
	var form UserLoginForm

	err := app.decodePostForm(r, &form)
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	form.CheckField(validator.NotBlank(form.Email), "email", "Email cannot be empty")
	form.CheckField(validator.Matches(form.Email, validator.EmailRX), "email", "Enter a valid Email")
	form.CheckField(validator.NotBlank(form.Password), "password", "Password cannot be blank")

	if !form.Valid() {
		data := app.newTemplateData(r)
		data.Form = form
		app.render(w, r, http.StatusUnprocessableEntity, "login.tmpl", data)
		return
	}

	id, err := app.users.Authenticate(form.Email, form.Password)
	if err != nil {
		if errors.Is(err, models.ErrInvalidCredentials) {
			form.AddNonFieldError("Email or Password Invalid")

			data := app.newTemplateData(r)
			data.Form = form
			app.render(w, r, http.StatusUnprocessableEntity, "login.tmpl", data)
		} else {
			app.serverError(w, r, err)
		}
		return
	}

	err = app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	app.sessionManager.Put(r.Context(), "authenticatedUserID", id)

	http.Redirect(w, r, "/snippet/create", http.StatusSeeOther)

}
func (app *application) UserLogoutPost(w http.ResponseWriter, r *http.Request) {

	app.sessionManager.Remove(r.Context(), "authenticatedID")

	err := app.sessionManager.RenewToken(r.Context())
	if err != nil {
		app.serverError(w, r, err)
		return
	}
	app.sessionManager.Put(r.Context(), "flash", "Log out Succesful!")

	app.sessionManager.Clear(r.Context())

	http.Redirect(w, r, "/", http.StatusSeeOther)

}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("OK"))
}
