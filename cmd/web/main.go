package main

import (
	"fmt"
	"github.com/alexedwards/scs/v2"
	"hotel_reservation/pkg/config"
	"hotel_reservation/pkg/handlers"
	"hotel_reservation/pkg/render"
	"log"
	"net/http"
	"time"
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
		log.Fatal(err)
	}

	app.UseCache = false

	app.TemplateCache = tc
	repo := handlers.InitRepo(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(),
	}

	fmt.Printf("🚀 Server running on http://localhost%s\n", portNumber)
	log.Panic(server.ListenAndServe())
}
