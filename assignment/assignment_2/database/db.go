package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var (
	username = "postgres"
	password = "naufalaldy23"
	host     = "localhost"
	port     = "5432"
	dbName   = "assignment2"
	db       *sql.DB
	err      error
)

func createRequiredTable() {
	// ! Membuat tabel order
	orderTable := `CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		customer_name VARCHAR(255) NOT NULL,
		ordered_at timestamptz NOT NULL DEFAULT (now())
		);
	`

	_, err := db.Exec(orderTable)
	if err != nil {
		log.Fatal("Error creating order table:", err.Error())
	}

	// ! Membuat tabel items
	itemsTable := `CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		item_code VARCHAR(255) NOT NULL,
		description TEXT NOT NULL,
		quantity INT NOT NULL,
		order_id INT NOT NULL,
		FOREIGN KEY (order_id) REFERENCES orders(id)
	);
	`

	_, err = db.Exec(itemsTable)
	if err != nil {
		log.Fatal("Error creating items table:", err.Error())
	}
}

func InitializeDB() {
	// postgresql://username:password@host:5432/assignment2?sslmode=disable
	config := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, port, dbName)
	db, err = sql.Open("postgres", config)

	if err != nil {
		log.Fatal("Error connecting to database:", err.Error())
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error while trying to ping database connection:", err.Error())
	}

	fmt.Println("Successfully connected to database!")
	createRequiredTable()
}

func GetDB() *sql.DB {
	return db
}
