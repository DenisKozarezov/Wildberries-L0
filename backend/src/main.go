package main

import (
	"context"
	"log"
	db "myapp/database"
	http "myapp/http"
)

func main() {
	ctx := context.Background()
	err := db.Connect(ctx, "postgres", "root", "wildberriesDb")

	if err != nil {
		log.Fatalln(err)
	} else {
		db.Configurate()
	}

	defer db.Disconnect()

	http.StartListening(":8080")
}
