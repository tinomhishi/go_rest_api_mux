package main

import (
	"database/sql"
	"errors"
)

type product struct{
	ID int `json:"id"`
	NAME  string `json:"name"`
	PRICE float64 `json:"price"`
}

\\ CRUD Functions

func (p *product) updateProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *product) getProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *product) deleteProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func (p *product) createProduct(db *sql.DB) error {
	return errors.New("Not Implemented")
}

func getProducts(db *sql.DB, start, count, int) ([] product, error) {
	return nil, errors.New("Not Implemented")
}