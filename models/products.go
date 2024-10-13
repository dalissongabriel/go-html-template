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

	sqlFindAllProducts, err := db.Query("SELECT * FROM products ORDER BY id")

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

func FindOneProduct(productId int) Product {
	db := database.ConnDatabase()

	var product Product

	sqlFindProductById, err := db.Query("SELECT * FROM products WHERE id = $1", productId)

	if err != nil {
		panic(err.Error())
	}

	for sqlFindProductById.Next() {
		sqlFindProductById.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Amount)
	}

	defer db.Close()
	return product
}

func InsertProduct(name, description string, price float64, amount int) {
	db := database.ConnDatabase()

	sqlInsertProducts, err := db.Prepare("INSERT INTO products(name, description, price, amount) VALUES ($1, $2, $3, $4)")

	if err != nil {
		panic(err.Error())
	}

	sqlInsertProducts.Exec(name, description, price, amount)

	defer db.Close()
}

func UpdateProduct(name, description string, price float64, amount int, id int) {
	db := database.ConnDatabase()

	sqlUpdateProduct, err := db.Prepare("UPDATE products SET name = $1, description = $2, price = $3, amount = $4 WHERE id = $5")

	if err != nil {
		panic(err.Error())
	}

	_, errA := sqlUpdateProduct.Exec(name, description, price, amount, id)

	if errA != nil {
		panic(errA.Error())
	}

	defer db.Close()
}

func DeleteProduct(id int) {
	db := database.ConnDatabase()

	sqlDeleteProduct, err := db.Prepare("DELETE FROM products WHERE id = $1")

	if err != nil {
		panic(err.Error())
	}

	sqlDeleteProduct.Exec(id)

	defer db.Close()
}
