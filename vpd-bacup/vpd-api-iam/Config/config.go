package configuration

import (
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func DbConn() (db *sqlx.DB) {
	db, err := sqlx.Connect("postgres", "user=postgres password=123 dbname=vpd_payroll_db sslmode=disable")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
