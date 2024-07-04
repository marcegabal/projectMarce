package db

import (
	"database/sql"
	"fmt"
)

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "marce5891"
	dbname   = "bd_pruebas"
)

func MainDb() {
	fmt.Println("conecto")
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	CheckError(err)
	defer db.Close()
}

func CheckError(err error){
	if err != nil{
		panic(err)
	}
}
