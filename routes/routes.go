package routes

import (
	"net/http"

	"github.com/dalissongabriel/go-html-template/controllers"
)

func RequestHandler() {
	http.HandleFunc("/", controllers.ProductsList)
	http.HandleFunc("/new", controllers.ProductsNew)
	http.HandleFunc("/insert", controllers.ProductsInsert)
	http.HandleFunc("/delete", controllers.ProductsDelete)
	http.HandleFunc("/edit", controllers.ProductEdit)
	http.HandleFunc("/update", controllers.ProductUpdate)
}
