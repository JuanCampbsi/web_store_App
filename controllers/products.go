package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"projects/goStudyng/Web_Store_App/models"

	_ "github.com/lib/pq"
)

var temp = template.Must(template.ParseGlob(("templates/*.html")))

func Index(w http.ResponseWriter, r *http.Request) {
	products := models.SearchProducts()
	temp.ExecuteTemplate(w, "Index", products)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceConv, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Err conversion price:", err)
		}

		amountConv, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Err conversion amount:", err)
		}
		products := models.Product{
			Name:        name,
			Description: description,
			Price:       priceConv,
			Amount:      amountConv,
		}
		models.CreateProduct(&products)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")

	models.DeleteProduct(idProduct)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduct := r.URL.Query().Get("id")
	product := models.EditProduct(idProduct)
	temp.ExecuteTemplate(w, "Edit", product)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		idConv, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Err conversion ID:", err)
		}
		priceConv, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Err conversion Price:", err)
		}
		amountConv, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Err conversion Amount:", err)
		}
		productsData := models.Product{
			Id:          idConv,
			Name:        name,
			Description: description,
			Price:       priceConv,
			Amount:      amountConv,
		}
		models.UpdateProduct(&productsData)
	}
	http.Redirect(w, r, "/", 301)
}
