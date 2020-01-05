package configuration

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

//InitializeDbCon :
func InitializeDbCon() *sqlx.DB {
	var err error
	db, err := sqlx.Open("postgres", "user=postgres password=1234 dbname=vpd_payroll_db sslmode=disable")
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	fmt.Println("Connected to database successfully!")
	return db

}
