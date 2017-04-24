package helper

import (
	"html/template"
	"net/http"
	"reflect"
)

var templates map[string]*template.Template

func init() {
	templates = makeTemplates()
}

func makeTemplates() map[string]*template.Template {
	templates = make(map[string]*template.Template)

	templates["login"] = template.Must(template.New("").Funcs(template.FuncMap{"hasField": hasField}).ParseFiles("templates/login.html", "templates/base.html"))
	templates["register"] = template.Must(template.New("").Funcs(template.FuncMap{"hasField": hasField}).ParseFiles("templates/register.html", "templates/base.html"))
	templates["profile"] = template.Must(template.New("").Funcs(template.FuncMap{"hasField": hasField}).ParseFiles("templates/profile.html", "templates/base.html"))
	templates["users"] = template.Must(template.New("").Funcs(template.FuncMap{"hasField": hasField}).ParseFiles("templates/users.html", "templates/base.html"))
	templates["user"] = template.Must(template.New("").Funcs(template.FuncMap{"hasField": hasField}).ParseFiles("templates/user.html", "templates/base.html"))
	return templates
}

func hasField(v interface{}, name string) bool {
	rv := reflect.ValueOf(v)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
	}
	if rv.Kind() != reflect.Struct {
		return false
	}
	return rv.FieldByName(name).IsValid()
}

// RenderTemplate is a helper that renders a view based on the template name
// e.g. RenderTemplate(w, "index", "base", noteStore)
func RenderTemplate(w http.ResponseWriter, name string, template string, viewmodel interface{}) {
	if templates == nil {
		templates = makeTemplates()
	}

	tmpl, ok := templates[name]
	if !ok {
		http.Error(w, "The template does not exist", http.StatusInternalServerError)
	}

	err := tmpl.ExecuteTemplate(w, template, viewmodel)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
