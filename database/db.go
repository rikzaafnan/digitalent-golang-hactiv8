package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	username = "root"
	password = "root"
	dbDriver = "postgres"
	dbPort   = "5432"
	dbName   = "golang_test"
	host     = "localhost"
	db       *sqlx.DB
	err      error
)

func createRequiredTables() {
	orderTable := `
	CREATE TABLE IF NOT EXISTS orders (
		id SERIAL PRIMARY KEY,
		customer_name varchar(255) NOT NULL,
		ordered_at timestamptz NOT NULL DEFAULT (now()),
		updated_at timestamptz NOT NULL DEFAULT (now())
	);`

	itemTable := `
	CREATE TABLE IF NOT EXISTS items (
		id SERIAL PRIMARY KEY,
		item_code varchar(255) NOT NULL,
		description varchar(255) NOT NULL,
		quantity int NOT NULL,
		order_id int NOT NULL,
		created_at timestamptz NOT NULL DEFAULT (now()),
		updated_at timestamptz NULL DEFAULT NULL,
		CONSTRAINT item_order_id_fk
			FOREIGN KEY(order_id) 
				REFERENCES orders(id)
					ON DELETE SET NULL	
	);`

	createTableQueries := fmt.Sprintf("%s  %s", orderTable, itemTable)
	_, err = db.Exec(createTableQueries)

	if err != nil {
		log.Fatal("error while creating movies table =>", err.Error())
	}
}

func InitializeDB() {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", username, password, host, dbPort, dbName)

	db, err = sqlx.Connect(dbDriver, dsn)

	if err != nil {
		log.Fatal("error connecting to database", err.Error())
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("error while tyring to ping the database connection", err.Error())
	}

	fmt.Println("successfully connected to my database")
	createRequiredTables()
}

func GetDB() *sqlx.DB {
	return db
}
