package bd

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func ConnectionDB() *sql.DB {
	connectionStr := "user=postgres dbname=alura_loja password=D@nS@142720 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connectionStr)
	isErro(err)
	return db
}

func isErro(err error) {
	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
		log.Fatal(err)
		panic(err.Error())
	}
}
