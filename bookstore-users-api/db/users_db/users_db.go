package users_db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var (
	Client *sql.DB
)

func init() {
	var err error

	if err = godotenv.Load(); err != nil {
		log.Fatal(err)
	}

	datasourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_SCHEMA"),
	)

	Client, err = sql.Open("mysql", datasourceName)
	if err != nil {
		panic(err)
	}
	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")

}
