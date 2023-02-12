package handlers

import (
	"net/http"

	"github.com/sksahu2097/go-project/pkg/config"
	"github.com/sksahu2097/go-project/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

// set the AppConfig in the Repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func SetRepo(a *Repository) {
	Repo = a
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTenplate(w, "home.page.tmpl")
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	render.RenderTenplate(w, "aboutUs.page.tmpl")
}
