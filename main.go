package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/williamrlbrito/hexagonal-architecture/adapters/db"
	"github.com/williamrlbrito/hexagonal-architecture/application"
)

func main() {
	sqlite, _ := sql.Open("sqlite3", "sqlite.db")
	defer sqlite.Close()
	productDbAdapter := db.NewProductDb(sqlite)
	productService := application.NewProductService(productDbAdapter)
	product, _ := productService.Create("Product 1", 10)
	productService.Enable(product)
}
