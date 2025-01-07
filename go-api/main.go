package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/TasosFrago/epms/db_connection"
	"github.com/TasosFrago/epms/router"
	"github.com/TasosFrago/epms/utls"
)

func main() {
	// Load Environment variables
	fmt.Println("Starting app...")
	utls.LoadEnv()

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("Server address %s\n", address)

	server := &http.Server{
		Addr: address,
	}

	// Channel to receive db connection
	dbChan := make(chan *db_connection.DBConn, 1)
	errChan := make(chan error, 1)

	go func() {
		config := db_connection.CredentialConfig{
			Usrname:    os.Getenv("USERNAME"),
			Passwd:     os.Getenv("PASSWORD"),
			ServerHost: os.Getenv("HOST"),
			ServerPort: os.Getenv("SSH_PORT"),
			DBHost:     "localhost:3306",
			DBName:     "lab2425omada1_EPMS",
		}

		fmt.Println("Connecting to database via SSH tunnel...")
		db, err := db_connection.ConnectDBoSSH(config)
		if err != nil {
			errChan <- err
			return
		}
		dbChan <- db
	}()

	// Start API server immediately
	api := router.NewServer(address)
	go func() {
		if err := api.RunWithTemporaryHandlers(server); err != nil {
			log.Printf("Error starting server: %v", err)
		}
	}()

	// Wait for database connection result asynchronously
	select {
	case db := <-dbChan:
		fmt.Println("Database connection established!")
		api.SetDB(db.Conn) // Assuming router.Server has a SetDB method
		defer db.Cleanup()
		server.Close()
		if err := api.Run(); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	case err := <-errChan:
		log.Fatalf("Database connection error: %v", err)
	case <-time.After(15 * time.Second):
		log.Fatalf("Timeout waiting for database connection")
	}

	// defer db.Cleanup()
	//
	// api := router.NewServer(address, db.Conn)
}
