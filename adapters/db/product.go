package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/williamrlbrito/hexagonal-architecture/application"
)

type ProductDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProductDb {
	return &ProductDb{db: db}
}

func (productDb *ProductDb) Get(id string) (application.ProductInterface, error) {
	product := application.Product{}
	stmt, err := productDb.db.Prepare("select id, name, price, status from products where id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	if err := stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status); err != nil {
		return nil, err
	}

	return &product, nil
}
