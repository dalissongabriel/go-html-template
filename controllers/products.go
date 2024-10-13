package controllers

import (
	"log"
	"net/http"
	"strconv"
	"text/template"

	"github.com/dalissongabriel/go-html-template/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func ProductsList(w http.ResponseWriter, r *http.Request) {
	products := models.FindAllProducts()
	templates.ExecuteTemplate(w, "Index", products)
}

func ProductsNew(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "New", nil)
}

func ProductsInsert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preco:", err)
		}

		amountInteger, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		models.InsertProduct(name, description, priceFloat, amountInteger)
	}

	defer http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func ProductsDelete(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")

	parsedProductId, err := strconv.Atoi(productId)

	if err != nil {
		panic(err.Error())
	}

	models.DeleteProduct(parsedProductId)

	defer http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

func ProductEdit(w http.ResponseWriter, r *http.Request) {
	productId := r.URL.Query().Get("id")
	parsedProductId, err := strconv.Atoi(productId)
	if err != nil {
		panic(err.Error())
	}

	product := models.FindOneProduct(parsedProductId)
	templates.ExecuteTemplate(w, "Edit", product)
}

func ProductUpdate(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		description := r.FormValue("description")
		price := r.FormValue("price")
		amount := r.FormValue("amount")

		priceFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("Erro na conversão do preço:", err)
		}

		amountInteger, err := strconv.Atoi(amount)
		if err != nil {
			log.Println("Erro na conversão da quantidade:", err)
		}

		idInteger, err := strconv.Atoi(id)
		if err != nil {
			log.Println("Erro na conversão do ID:", err)
		}

		models.UpdateProduct(name, description, priceFloat, amountInteger, idInteger)
	}

	defer http.Redirect(w, r, "/", http.StatusMovedPermanently)
}
