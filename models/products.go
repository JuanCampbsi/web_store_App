package models

import (
	"log"
	db "projects/goStudyng/Web_Store_App/db"
)

type Product struct {
	Id, Amount        int
	Name, Description string
	Price             float64
}

func SearchProducts() []Product {
	db := db.ConnectDatabase()
	selectProductsListAll, err := db.Query("select * from products")

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
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Amount = amount

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateProduct(model *Product) {
	db := db.ConnectDatabase()

	insertData, err := db.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertData.Exec(model.Name, model.Description, model.Price, model.Amount)
	defer db.Close()
	log.Println("Product object:", model)
}

func DeleteProduct(id string) {
	log.Println("Product id delete:", id)
	db := db.ConnectDatabase()
	deleteProduct, err := db.Prepare("delete from products where id=$1")
	if err != nil {
		panic(err.Error())
	}
	deleteProduct.Exec(id)
	defer db.Close()
}
