package handlers

import (
	"html/template"
	"net/http"
	"path"
)

// TODO: probably need to move to components, or something
type CompaniesPage struct {
	Name string
}

func CompaniesHandler(w http.ResponseWriter, r *http.Request) {
	name := "MEGACORP" // TODO: pull this from db
	p := &CompaniesPage{Name: name}
	renderTemplate(w, "companies", p)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *CompaniesPage) {
	templatePath := path.Join("web/templates/pages", tmpl+".html")

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = t.Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
