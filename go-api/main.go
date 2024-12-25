package main

import (
	"log"
	"os"

	"github.com/TasosFrago/epms/db_connection"
	"github.com/TasosFrago/epms/router"
	"github.com/TasosFrago/epms/utls"

	_ "github.com/gorilla/handlers"
)

func main() {
	// Load Environment variables
	utls.LoadEnv()

	config := db_connection.CredentialConfig{
		Usrname:    os.Getenv("USERNAME"),
		Passwd:     os.Getenv("PASSWORD"),
		ServerHost: os.Getenv("HOST"),
		ServerPort: os.Getenv("PORT"),
		DBHost:     "localhost:3306",
		DBName:     "lab2425omada1_EPMS",
	}

	db, err := db_connection.ConnectDBoSSH(config)
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}

	defer db.Cleanup()

	api := router.NewServer("0.0.0.0:8080", db.Conn)
	if err := api.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
