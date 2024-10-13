package models

import (
	"github.com/dalissongabriel/go-html-template/database"
)

type Product struct {
	Id          int
	Name        string
	Description string
	Price       float64
	Amount      int
}

func FindAllProducts() []Product {
	db := database.ConnDatabase()

	var product Product
	var products []Product

	sqlFindAllProducts, err := db.Query("SELECT * from products")

	if err != nil {
		panic(err.Error())
	}

	for sqlFindAllProducts.Next() {
		var id, amount int
		var name, description string
		var price float64

		err = sqlFindAllProducts.Scan(&id, &name, &description, &price, &amount)

		if err != nil {
			panic(err.Error())
		}

		product.Id = id
		product.Name = name
		product.Description = description
		product.Price = price
		product.Amount = amount

		products = append(products, product)
	}

	defer db.Close()
	return products
}

func InsertProduct(name, description string, price float64, amount int) {
	db := database.ConnDatabase()

	sqlInsertProducts, err := db.Prepare("insert into products(name, description, price, amount) values ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	sqlInsertProducts.Exec(name, description, price, amount)

	defer db.Close()
}

func DeleteProduct(id int) {
	db := database.ConnDatabase()

	sqlDeleteProduct, err := db.Prepare("delete from products where id = $1")

	if err != nil {
		panic(err.Error())
	}

	sqlDeleteProduct.Exec(id)

	defer db.Close()
}
