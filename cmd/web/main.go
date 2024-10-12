package main

import (
	"log"
	"net/http"

	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mihasiok/bookings/pkg/config"
	"github.com/mihasiok/bookings/pkg/handlers"
	"github.com/mihasiok/bookings/pkg/render"
)

const portNumber = ":8080"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	// change this to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal()
	}

	app.TemplateCache = tc
	app.UseCacbe = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTamplates(&app)

	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
