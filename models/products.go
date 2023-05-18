package models

import db "projects/goStudyng/Web_Store_App/db"

type Product struct {
	Name, Description string
	Price             float64
	Amount            int
}

func SearchProducts() []Product {
	db := db.ConnectDatabase()
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
	defer db.Close()
	return products
}
