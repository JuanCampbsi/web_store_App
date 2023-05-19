package models

import (
	"log"
	db "store/db"
	"store/store"
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

func CreateProduct(serializedData *store.Product) {
	db := db.ConnectDatabase()

	insertData, err := db.Prepare("insert into products(name, description, price, amount) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	insertData.Exec(serializedData.Name(), serializedData.Description(), serializedData.Price(), serializedData.Amount())
	defer db.Close()
	log.Println("Product object:", serializedData.Name())
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

func EditProduct(id string) Product {
	db := db.ConnectDatabase()
	productDataBase, err := db.Query("select * from products where id = $1", id)
	if err != nil {
		panic(err.Error())
	}
	productToUpdate := Product{}
	for productDataBase.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = productDataBase.Scan(&id, &name, &description, &price, &amount)
		if err != nil {
			panic(err.Error())
		}
		productToUpdate.Id = id
		productToUpdate.Name = name
		productToUpdate.Description = description
		productToUpdate.Price = price
		productToUpdate.Amount = amount
	}
	defer db.Close()
	return productToUpdate
}

func UpdateProduct(serializedData *store.Product) {
	db := db.ConnectDatabase()

	updateData, err := db.Prepare("update products set name = $2, description = $3, price = $4, amount = $5 where id = $1")
	if err != nil {
		panic(err.Error())
	}
	updateData.Exec(serializedData.Id(), serializedData.Name(), serializedData.Description(), serializedData.Price(), serializedData.Amount())
	defer db.Close()
	log.Println("Product object update:", serializedData.Id())
}
