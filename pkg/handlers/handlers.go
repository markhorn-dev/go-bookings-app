package handlers

import (
	"net/http"

	"github.com/markhorn-dev/go-bookings-app/pkg/config"
	"github.com/markhorn-dev/go-bookings-app/pkg/models"
	"github.com/markhorn-dev/go-bookings-app/pkg/render"
)

// Repo the repository used by the handlers
var Repo *Repository

// Repository is the repository type
type Repository struct {
	App *config.AppConfig
}

// NewRepo creates a new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}

// Home is the handler for the home page
func (m *Repository) Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr
	m.App.Session.Put(r.Context(), "remote_ip", remoteIP)

	render.RenderTemplate(w, "home.page.html", &models.TemplateData{})
}

// About is the handler for the about page
func (m *Repository) About(w http.ResponseWriter, r *http.Request) {
	// perform some logic
	stringMap := make(map[string]string)
	stringMap["test"] = "Hello, again"

	remoteIP := m.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	// send data to the template
	render.RenderTemplate(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}

// Contact is the handler for the contact page
func (m *Repository) Contact(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "contact.page.html", &models.TemplateData{})
}

// Availability is the handler for the availability page
func (m *Repository) Availability(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "availability.page.html", &models.TemplateData{})
}

// Reservation is the handler for the reservation page
func (m *Repository) Reservation(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "reservation.page.html", &models.TemplateData{})
}

// GeneralsSuite is the handler for the generals-suite page
func (m *Repository) GeneralsSuite(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "generals-suite.page.html", &models.TemplateData{})
}

// MajorsQuarters is the handler for the majors-quarters page
func (m *Repository) MajorsQuarters(w http.ResponseWriter, r *http.Request) {
	render.RenderTemplate(w, "majors-quarters.page.html", &models.TemplateData{})
}
