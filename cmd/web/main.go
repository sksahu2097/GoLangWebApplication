package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sksahu2097/go-project/pkg/config"
	"github.com/sksahu2097/go-project/pkg/handlers"
	"github.com/sksahu2097/go-project/pkg/render"
)

func main() {

	var app config.AppConfig
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatal("Error while creating application config")
	}
	app.TemplateCache = tc
	render.SetTemplateAppConfig(&app)
	repo := handlers.NewRepo(&app)
	handlers.SetRepo(repo)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	fmt.Println("Starting the application on 8080")
	_ = http.ListenAndServe(":8080", nil)

}
