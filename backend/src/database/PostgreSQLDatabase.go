package database

import (
	"database/sql"
	"fmt"
	"log"
	"myapp/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Config struct {
	Database string
	User     string
	Password string
	Host     string
	Protocol string
}

func Connect(login string, pass string, addr string, dbName string) error {
	log.Println("Opening a connection to database...")

	cfg := Config{
		Database: dbName,
		User:     login,
		Password: pass,
		Host:     addr,
		Protocol: http.PROT_TCP,
	}

	connectionStr := fmt.Sprintf("%s://%s:%s@%s/postgres?sslmode=disable", cfg.Database, cfg.User, cfg.Password, cfg.Host)
	log.Println("Connection string:", connectionStr)

	defer Print(cfg)

	// Get a database handle.
	var err error
	db, err = sql.Open("postgres", connectionStr)
	if err != nil {
		log.Fatal(err)
		return err
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
		return pingErr
	}
	log.Println("Connected to database!")
	return err
}

func Configurate() {
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(time.Minute * 5)
}

func Disonnect() {
	log.Println("Disconnecting from the database...")
	db.Close()
}

func Print(cfg Config) {
	log.Println("============= INFO =============")
	log.Println("DB name: \t", cfg.Database)
	log.Println("User: \t", cfg.User)
	log.Println("Host: \t", cfg.Host)
	log.Println("Protocol:\t", cfg.Protocol)
	log.Println("Driver:  \t PostgreSQL")
}
