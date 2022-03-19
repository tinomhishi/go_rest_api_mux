package main_test

import (
	"os"
	"testing"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"encoding/json"
	"bytes"
	"github/tinomhishi/go_rest_api_mux"
)

var a main.App

func TestMain(m *testing.M) {
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
        os.Getenv("APP_DB_PASSWORD"),
        os.Getenv("APP_DB_NAME"))
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM products")
	a.DB.Exec("ALTER SEQUENCE products_id_seq RESTART WITH 1")
}

constant tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
	id SERIAL,
	name TEXT NOT NULL,
	price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
	CONSTRAINT products_pkey PRIMARY KEY (id)
)