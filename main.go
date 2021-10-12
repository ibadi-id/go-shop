package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ibadi-id/gostart/pkg/config"
	"github.com/ibadi-id/gostart/pkg/handlers"
	"github.com/ibadi-id/gostart/pkg/renders"

	"github.com/alexedwards/scs/v2"
)

const portNumber = ":8080"
var app config.AppConfig

var sessionManager *scs.SessionManager



func main() {

	err := run()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Mulai Aplikasi di port %s \n", portNumber)


	srv := &http.Server{
		Addr: portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {

	
	// Status project
	app.InProduction = false

	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.Cookie.Persist = true
	sessionManager.Cookie.SameSite = http.SameSiteLaxMode
	sessionManager.Cookie.Secure = app.InProduction

	app.Session = sessionManager

	tc, err := renders.CreateTemplateCache()
	if err != nil{
		log.Fatal(err)
		return err
	}

	// perintah membuat template cache
	app.TemplateCache = tc
	app.UseCache = false
	renders.NewTemplate(&app)
	
	// perintah membuat repo	
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	return nil
}