package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"store/models"
	"store/store"

	flatbuffers "github.com/google/flatbuffers/go"
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

		builder := flatbuffers.NewBuilder(1024)

		name := builder.CreateString(r.FormValue("name"))
		description := builder.CreateString(r.FormValue("description"))

		priceConv, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Err conversion price:", err)
		}

		amountConv, err := strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			log.Println("Err conversion amount:", err)
		}

		store.ProductStart(builder)
		store.ProductAddName(builder, name)
		store.ProductAddDescription(builder, description)
		store.ProductAddPrice(builder, priceConv)
		store.ProductAddAmount(builder, int32(amountConv))

		serializedData := store.ProductEnd(builder)
		builder.Finish(serializedData)
		buf := builder.FinishedBytes()
		data := store.GetRootAsProduct(buf, 0)

		log.Println("Product object:", data.Name())
		models.CreateProduct(data)
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
		builder := flatbuffers.NewBuilder(1024)

		name := builder.CreateString(r.FormValue("name"))
		description := builder.CreateString(r.FormValue("description"))

		idConv, err := strconv.Atoi(r.FormValue("id"))
		if err != nil {
			log.Println("Err conversion ID:", err)
		}

		priceConv, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			log.Println("Err conversion price:", err)
		}

		amountConv, err := strconv.Atoi(r.FormValue("amount"))
		if err != nil {
			log.Println("Err conversion amount:", err)
		}

		store.ProductStart(builder)
		store.ProductAddId(builder, int32(idConv))
		store.ProductAddName(builder, name)
		store.ProductAddDescription(builder, description)
		store.ProductAddPrice(builder, priceConv)
		store.ProductAddAmount(builder, int32(amountConv))

		serializedData := store.ProductEnd(builder)
		builder.Finish(serializedData)
		buf := builder.FinishedBytes()
		data := store.GetRootAsProduct(buf, 0)

		models.UpdateProduct(data)
	}
	http.Redirect(w, r, "/", 301)
}
