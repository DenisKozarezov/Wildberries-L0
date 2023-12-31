package database

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx"
)

var db *pgx.Conn

func print(cfg *pgx.ConnConfig) {
	log.Println("============= INFO =============")
	log.Println("DB Name: \t", cfg.Database)
	log.Println("User: \t", cfg.User)
	log.Println("Host: \t", cfg.Host)
	log.Println("Port: \t", cfg.Port)
}

func connect(ctx context.Context, login string, pass string, dbName string) error {
	log.Println("Opening a connection to database...")

	connectionStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", login, pass, "localhost:5432", dbName)
	cfg, err := pgx.ParseConnectionString(connectionStr)
	log.Println("Connection string:", connectionStr)

	defer print(&cfg)

	// Get a database handle.
	db, err = pgx.Connect(cfg)
	if err != nil {
		log.Fatal(err)
		return err
	}

	// Check the connection
	if pingErr := db.Ping(ctx); pingErr != nil {
		return pingErr
	}
	log.Println("Connected to database!")
	return err
}

func ConnectToDatabase() {
	ctx := context.Background()
	err := connect(ctx, "postgres", "root", "wildberriesDb")

	if err != nil {
		log.Fatalln(err)
	}
}

func Disconnect() {
	log.Println("Disconnecting from the database...")
	db.Close()
}
