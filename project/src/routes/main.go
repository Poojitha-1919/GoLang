package routes

import (
	"Project/src/products"
	"net/http"
)

func SetUpRoutes() {
	http.HandleFunc("/products", products.GetAllProducts)
	//Add routes
}
