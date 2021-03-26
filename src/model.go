package main

import (
	"database/sql"
	"strings"
)

type product struct {
	ID    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func (p *product) getProduct(db *sql.DB) error {
	return db.QueryRow("SELECT name, price FROM products WHERE id=$1", p.ID).Scan(&p.Name, &p.Price)
}

func (p *product) updateProduct(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE products SET name=$1, price=$2 WHERE id=$3",
			p.Name, p.Price, p.ID)

	return err
}

func (p *product) deleteProduct(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM products WHERE id=$1", p.ID)

	return err
}

func (p *product) createProduct(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO products(name, price) VALUES($1, $2) RETURNING id",
		p.Name, p.Price).Scan(&p.ID)

	if err != nil {
		return err
	}

	return nil
}

func getProducts(db *sql.DB, start, count int) ([]product, error) {
	rows, err := db.Query(
		"SELECT id, name,  price FROM products LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}

func getProductsCount(db *sql.DB) (int, error) {
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM products")
	err := row.Scan(&count)

	return count, err
}

type pricebin struct {
	Name  string  `json:"name"`
	Count float64 `json:"count"`
}

func getProductBins(db *sql.DB) ([]pricebin, error) {
	rows, err := db.Query(`
		SELECT bin as name, Count(bin) as count
		FROM (SELECT
				CASE
					WHEN price >= 0 and price <= 250 then '<=250'
					WHEN price > 250 and price <= 500 then '>250 <=500'
					WHEN price > 500 and price <= 750 then '>500 <=750'
					WHEN price > 750 and price <= 1000 then '>750 <=1000'
					ELSE '>1000'
				End as bin
			FROM products) bins
		GROUP BY bin`)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	pricebins := []pricebin{}
	for rows.Next() {
		var p pricebin
		if err := rows.Scan(&p.Name, &p.Count); err != nil {
			return nil, err
		}
		pricebins = append(pricebins, p)
	}

	return pricebins, nil
}

func searchProducts(db *sql.DB, search string, start, count int) ([]product, error) {
	rows, err := db.Query(`
		SELECT id, name, price
		FROM products
		WHERE LOWER(name) like $1
		LIMIT $2 OFFSET $3
	`,
		"%"+strings.ToLower(search)+"%",
		count,
		start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	products := []product{}

	for rows.Next() {
		var p product
		if err := rows.Scan(&p.ID, &p.Name, &p.Price); err != nil {
			return nil, err
		}
		products = append(products, p)
	}

	return products, nil
}
