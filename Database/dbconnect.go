package Database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/project104/Models"
)

var DB *sqlx.DB
var err error

func Connect() {

	dsn := "host=localhost user=postgres password=8055 dbname=project104db port=5432 sslmode=disable"
	dbconnect, err := sqlx.Open("postgres", dsn)

	if err != nil {
		//		panic(err)fmt
		fmt.Println("err")
	}
	if err = dbconnect.Ping(); err != nil {
		//		panic(err)
		fmt.Println("err2")
	}
	DB = dbconnect
	dbconnect.AutoMigrate(&Models.User{})
	dbconnect.AutoMigrate(&Models.List{})
}
