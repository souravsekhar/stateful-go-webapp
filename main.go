package main

import (
	"net/http"

	"github.com/souravsekhar/webapp/product"
)

const apiBasePath = "/api"

func main() {
	product.SetUpRoutes(apiBasePath)
	http.ListenAndServe(":5001", nil)
}
