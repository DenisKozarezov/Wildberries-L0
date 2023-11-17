package main

import (
	db "myapp/database"
	http "myapp/http"
	"myapp/services"
)

func main() {
	db.ConnectToDatabase()

	services.RestoreCache()
	services.ConnectToNATS()

	defer db.Disconnect()

	http.StartListening(":8080")
}
