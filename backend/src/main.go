package main

import (
	"context"
	"log"
	db "myapp/database"
	http "myapp/http"
	"myapp/services"
)

func ConnectToDatabase() {
	ctx := context.Background()
	err := db.Connect(ctx, "postgres", "root", "wildberriesDb")

	if err != nil {
		log.Fatalln(err)
	} else {
		db.Configurate()
	}
}

func main() {
	ConnectToDatabase()

	services.RestoreCache()
	services.ConnectToNATS()

	defer db.Disconnect()

	http.StartListening(":8080")
}
