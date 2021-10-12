package main

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/ibadi-id/gostart/pkg/config"
	"github.com/ibadi-id/gostart/pkg/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gotailwindcss/tailwind/twembed"
	"github.com/gotailwindcss/tailwind/twhandler"
)

func routes(app *config.AppConfig) http.Handler{
	mux := chi.NewRouter()

	//middleware
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)

	// aktifkan session , untuk develop tailwind non aktifkan dulu aja
	mux.Use(SessionLoad)

	//router
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	// css file using tailwind
	mux.Handle("/css/*", twhandler.New(http.Dir("css"), "/css", twembed.New()))

	// Create a route along /files that will serve contents from
	// the ./data/ folder.
	workDir, _ := os.Getwd()
	filesDir := http.Dir(filepath.Join(workDir, "image"))
	FileServer(mux, "/image", filesDir)

	return mux
}

// FileServer conveniently sets up a http.FileServer handler to serve
// static files from a http.FileSystem.
func FileServer(r chi.Router, path string, root http.FileSystem) {
	if strings.ContainsAny(path, "{}*") {
		panic("FileServer does not permit any URL parameters.")
	}

	if path != "/" && path[len(path)-1] != '/' {
		r.Get(path, http.RedirectHandler(path+"/", http.StatusMovedPermanently).ServeHTTP)
		path += "/"
	}
	path += "*"

	r.Get(path, func(w http.ResponseWriter, r *http.Request) {
		rctx := chi.RouteContext(r.Context())
		pathPrefix := strings.TrimSuffix(rctx.RoutePattern(), "/*")
		fs := http.StripPrefix(pathPrefix, http.FileServer(root))
		fs.ServeHTTP(w, r)
	})
}