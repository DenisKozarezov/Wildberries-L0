package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

var db *pgx.Conn

func Print(cfg *pgx.ConnConfig) {
	log.Println("============= INFO =============")
	log.Println("DB Name: \t", cfg.Database)
	log.Println("User: \t", cfg.User)
	log.Println("Host: \t", cfg.Host)
	log.Println("Port: \t", cfg.Port)
}

func Connect(ctx context.Context, login string, pass string, dbName string) error {
	log.Println("Opening a connection to database...")

	connectionStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", login, pass, "localhost:5432", dbName)
	cfg, err := pgx.ParseConnectionString(connectionStr)
	log.Println("Connection string:", connectionStr)

	defer Print(&cfg)

	// Get a database handle.
	db, err = pgx.Connect(cfg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	pingErr := Ping(ctx)
	if err != nil {
		return pingErr
	}

	log.Println("Connected to database!")
	return err
}

func Ping(ctx context.Context) error {
	return db.Ping(ctx)
}

func Configurate() {

}

func Disconnect() {
	log.Println("Disconnecting from the database...")
	db.Close()
}
