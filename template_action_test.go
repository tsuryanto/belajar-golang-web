package belajar_golang_web

import (
	"belajar-golang-web/helper"
	"html/template"
	"net/http"
	"testing"
)

/* TEMPLATE ACTION IF */

func TemplateActionIf(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(w, "if.gohtml", Page{
		Title: "Template Data Struct",
		Name:  "Taufiq Suryanto",
	})
}

func TestTemplateActionIf(t *testing.T) {
	helper.RunHttpTest(TemplateActionIf, localhost, http.MethodGet)
}

/* TEMPLATE ACTION OPERATOR */

func TemplateActionOperator(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(w, "comparator.gohtml", map[string]interface{}{
		"Title":      "Template Action Operator",
		"FinalValue": 90,
	})
}

func TestTemplateActionOperator(t *testing.T) {
	helper.RunHttpTest(TemplateActionOperator, localhost, http.MethodGet)
}

/* TEMPLATE ACTION RANGE */

func TemplateActionRange(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(w, "range.gohtml", map[string]interface{}{
		"Title": "Template Action Range",
		"Hobbies": []string{
			"game", "read", "code", "design",
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	helper.RunHttpTest(TemplateActionRange, localhost, http.MethodGet)
}

/* TEMPLATE ACTION WITH */

func TemplateActionWith(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(w, "with.gohtml", map[string]interface{}{
		"Title": "Template Action With",
		"Name":  "Taufiq Suryanto",
		"Address": map[string]interface{}{
			"Street": "Jl. Teluk Gong Raya",
			"City":   "Jakarta Utara",
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	helper.RunHttpTest(TemplateActionWith, localhost, http.MethodGet)
}

/* TEMPLATE LAYOUT */

func TemplateLayout(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles(
		"./templates/template_header.gohtml",
		"./templates/template_footer.gohtml",
		"./templates/template_body.gohtml",
	))
	t.ExecuteTemplate(w, "template_body.gohtml", map[string]interface{}{
		"Title": "Template Layout",
		"Name":  "Taufiq Suryanto",
	})
}

func TestTemplateLayout(t *testing.T) {
	helper.RunHttpTest(TemplateLayout, localhost, http.MethodGet)
}
