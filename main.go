package main

import (
	"html/template"
	"net/http"

	models "projects/goStudyng/Web_Store_App/models"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob(("templates/*.html")))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	products := models.SearchProducts()
	temp.ExecuteTemplate(w, "Index", products)
}
