package main

import (
	"net/http"

	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()
	fileserver := http.FileServer(http.Dir(("./internal/ui/static")))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileserver))

	// unprotected
	dynamic := alice.New(app.sessionManager.LoadAndSave, noSurf, app.authenticate)

	mux.Handle("GET /{$}", dynamic.ThenFunc(app.home))
	mux.Handle("GET /snippet/view/{id}", dynamic.ThenFunc(app.snippetView))
	mux.Handle("GET /user/signup", dynamic.ThenFunc(app.userSignup))
	mux.Handle("POST /user/signup", dynamic.ThenFunc(app.UserSignUpPost))
	mux.Handle("GET /user/login", dynamic.ThenFunc(app.UserLogin))
	mux.Handle("POST /user/login", dynamic.ThenFunc(app.UserLoginPost))

	// protected
	protected := dynamic.Append(app.requiresAuthentication)

	mux.Handle("GET /snippet/create", protected.ThenFunc(app.snippetCreate))
	mux.Handle("POST /snippet/create", protected.ThenFunc(app.snippetCreatePost))
	mux.Handle("POST /user/logout", protected.ThenFunc(app.UserLogoutPost))

	chain := alice.New(app.recoverPanic, app.logRequest, commonHeaders)

	return chain.Then(mux)

}
