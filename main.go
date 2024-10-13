package main

import (
	"fmt"
	"net/http"

	"github.com/dalissongabriel/go-html-template/routes"
)

func main() {
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.RequestHandler()
	http.ListenAndServe(":3001", nil)
}
