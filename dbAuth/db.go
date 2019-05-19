package dbAuth

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

func InitDB() {

	var err error

	db, err = sql.Open("postgres", "dbname=mydb user=postgres password="+os.Getenv("PGSQL")+" host=localhost port=5432 sslmode=disable")

	if err != nil {
		panic(err)
	}

}
