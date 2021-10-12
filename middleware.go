package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

// NoSurf adds CSRF protection to all post request
func NoSurf(next http.Handler) http.Handler{
	csrfHandler := nosurf.New(next)

	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path: "/",
		Secure: app.InProduction,
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad load and save the session on every request
func SessionLoad(next http.Handler) http.Handler{
	return sessionManager.LoadAndSave(next)
}