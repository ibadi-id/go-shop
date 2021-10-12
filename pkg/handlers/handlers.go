package handlers

import (
	"net/http"

	"github.com/ibadi-id/gostart/pkg/config"
	"github.com/ibadi-id/gostart/pkg/models"
	"github.com/ibadi-id/gostart/pkg/renders"
)

var Repo *Repository

type Repository struct{
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository{
	return &Repository{
		App : a,
	}
}

func NewHandlers(r *Repository){
	Repo = r
}

func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {

	remoteIp := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIp)

	renders.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// business logic untuk mendapatkan data

	stringMap := make(map[string]string)
	stringMap["test"] = "Hello Robbi"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	
	renders.RenderTemplate(w, "about.page.html", &models.TemplateData{StringMap: stringMap})

}
