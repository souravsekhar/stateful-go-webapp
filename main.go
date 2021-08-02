package main

import (
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/souravsekhar/webapp/database"
	"github.com/souravsekhar/webapp/product"
)

const apiBasePath = "/api"

func main() {
	database.SetupDatabase()
	product.SetUpRoutes(apiBasePath)
	http.ListenAndServe(":5001", nil)
}
