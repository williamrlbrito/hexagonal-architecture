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

func (productDb *ProductDb) Save(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := productDb.Get(product.GetID())

	if err != nil {
		return productDb.create(product)
	}

	return productDb.update(product)
}

func (productDb *ProductDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := productDb.db.Prepare("insert into products(id, name, price, status) values(?,?,?,?)")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (productDb *ProductDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := productDb.db.Prepare("update products set name = ?, price = ?, status = ? where id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())

	if err != nil {
		return nil, err
	}

	return product, nil
}
