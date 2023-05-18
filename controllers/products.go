package controllers

import (
	"html/template"
	"net/http"

	"projects/goStudyng/Web_Store_App/models"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob(("templates/*.html")))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SearchProducts()
	temp.ExecuteTemplate(w, "Index", products)
}
