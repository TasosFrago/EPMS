package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"log"
	"os"

	_ "github.com/gorilla/handlers"
	"github.com/joho/godotenv"

	"github.com/TasosFrago/epms/db_connection"
	"github.com/TasosFrago/epms/models"
	"github.com/TasosFrago/epms/router"
)

func main() {
	if _, err := os.Stat("../.env"); err == nil {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatalf("Error loading environment vars: %s", err)
		}
	}

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

	api := router.NewServer("0.0.0.0:8080")
	if err := api.Run(); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func consumerData(db *sql.DB) ([]models.Consumer, error) {
	var consumers []models.Consumer

	rows, err := db.Query("SELECT * FROM CONSUMER")
	if err != nil {
		return nil, fmt.Errorf("consumerData: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var cons models.Consumer
		if err := rows.Scan(
			&cons.ID,
			&cons.FirstName,
			&cons.LastName,
			&cons.Email,
			&cons.Password,
			&cons.Cell,
			&cons.Landline,
			&cons.CreditInfo,
		); err != nil {
			return nil, fmt.Errorf("consumerData: %v", err)
		}
		consumers = append(consumers, cons)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("consumerData: %v", err)
	}
	return consumers, nil
}
