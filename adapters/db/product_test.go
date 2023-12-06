package db_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/williamrlbrito/hexagonal-architecture/adapters/db"
)

var Db *sql.DB

func setUP() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := "create table products (id text, name text, price float, status text);"
	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := "insert into products values ('abc', 'Product Test', 10, 'enabled');"
	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUP()
	defer Db.Close()

	productDb := db.NewProductDb(Db)
	product, err := productDb.Get("abc")
	assert.Nil(t, err)
	assert.Equal(t, "abc", product.GetID())
	assert.Equal(t, "Product Test", product.GetName())
	assert.Equal(t, 10.0, product.GetPrice())
	assert.Equal(t, "enabled", product.GetStatus())
}
