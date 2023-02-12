package handlers

import (
	"net/http"

	"github.com/sksahu2097/go-project/pkg/render"
)

func Home(w http.ResponseWriter, r *http.Request) {
	render.RenderTenplate(w, "home.page.tmpl")
}

func About(w http.ResponseWriter, r *http.Request) {
	render.RenderTenplate(w, "aboutUs.page.tmpl")
}
