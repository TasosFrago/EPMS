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
	log.Println("Starting app...")
	utls.LoadEnv()

	address := fmt.Sprintf(":%s", os.Getenv("PORT"))
	log.Printf("Server address %s\n", address)

	server := &http.Server{
		Addr: address,
	}

	config := db_connection.CredentialConfig{
		Usrname:    os.Getenv("USERNAME"),
		Passwd:     os.Getenv("PASSWORD"),
		ServerHost: os.Getenv("HOST"),
		ServerPort: os.Getenv("SSH_PORT"),
		DBHost:     "localhost:3306",
		DBName:     "lab2425omada1_EPMS",
	}

	// Channel to receive db connection
	dbChan := make(chan *db_connection.DBConn, 1)
	errChan := make(chan error, 1)

	go func() {
		var db *db_connection.DBConn = nil
		for {
			if db == nil || db.Conn.Ping() != nil {
				log.Println("Attemting to connecct database via SSH tunnel...")
				if db == nil {
					log.Println("Connecting to db...")
				} else {
					log.Println("Reconnecting to the db...")
				}
				newDB, err := db_connection.ConnectDBoSSH(config)
				if err != nil {
					errChan <- err
					log.Println("WARNING: Error while connecting to DB:", err)
					time.Sleep(30 * time.Second)
					continue
				}
				db = newDB
				dbChan <- db
				log.Println("Successfully connected to db!")
			}
			time.Sleep(2 * time.Minute)
		}
	}()

	// Start API server immediately
	api := router.NewServer(address)
	go func() {
		if err := api.RunWithTemporaryHandlers(server); err != nil {
			log.Println("Info: Stopping temporary http server.")
		}
	}()

	// Wait for database connection result asynchronously
	select {
	case db := <-dbChan:
		log.Println("Database connection established!")
		api.SetDB(db.Conn)

		defer db.Cleanup()
		server.Close()
		if err := api.Run(); err != nil {
			log.Fatalf("Error starting server: %v", err)
		}
	case err := <-errChan:
		log.Fatalf("Database connection error: %v", err)
	case <-time.After(50 * time.Second):
		log.Fatalf("Timeout waiting for database connection")
	}

	// defer db.Cleanup()
	//
	// api := router.NewServer(address, db.Conn)
}
