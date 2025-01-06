package main

import (
	"fmt"
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
		ServerPort: os.Getenv("SSH_PORT"),
		DBHost:     "localhost:3306",
		DBName:     "lab2425omada1_EPMS",
	}

	db, err := db_connection.ConnectDBoSSH(config)
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}

	defer db.Cleanup()

	address := fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT"))
	api := router.NewServer(address, db.Conn)
	if err := api.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
