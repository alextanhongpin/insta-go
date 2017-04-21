package helper

import (
	"html/template"
	"net/http"
)

var templates map[string]*template.Template

func init() {
	if templates == nil {
		templates = make(map[string]*template.Template)
	}

	templates["login"] = template.Must(template.ParseFiles("templates/login.html", "templates/base.html"))
	templates["register"] = template.Must(template.ParseFiles("templates/register.html", "templates/base.html"))
	templates["profile"] = template.Must(template.ParseFiles("templates/profile.html", "templates/base.html"))

}

// RenderTemplate is a helper that renders a view based on the template name
// e.g. RenderTemplate(w, "index", "base", noteStore)
func RenderTemplate(w http.ResponseWriter, name string, template string, viewmodel interface{}) {
	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist", http.StatusInternalServerError)
	}

	err := tmpl.ExecuteTemplate(w, template, viewmodel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
