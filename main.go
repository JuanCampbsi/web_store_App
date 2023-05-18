package main

import (
	"database/sql"
	"html/template"
	"net/http"
	"os"

	godotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func connectDatabase() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	connection := os.Getenv("CONNECTION_DATABASE")
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Product struct {
	Name, Description string
	Price             float64
	Amount            int
}

var temp = template.Must(template.ParseGlob(("templates/*.html")))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := connectDatabase()
	selectProductsListAll, err := db.Query("SELECT * FROM products")

	if err != nil {
		panic(err.Error())
	}

	p := Product{}
	products := []Product{}

	for selectProductsListAll.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = selectProductsListAll.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	temp.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}
