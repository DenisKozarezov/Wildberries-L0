package main

import (
	db "myapp/database"
	_ "myapp/http"
)

func main() {
	err := db.Connect("postgres", "root", "localhost:5432", "postgres")

	if err != nil {
		panic(err)
	} else {
		db.Configurate()
	}

	db.Disonnect()
}
