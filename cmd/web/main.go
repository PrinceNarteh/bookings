package main

import (
	"fmt"
	"hotel_reservation/pkg/config"
	"hotel_reservation/pkg/handlers"
	"hotel_reservation/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

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
		Handler: routes(&app),
	}

	fmt.Printf("🚀 Server running on http://localhost%s\n", portNumber)
	log.Panic(server.ListenAndServe())
}
