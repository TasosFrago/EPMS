package main

import (
    "database/sql"
    "fmt"
    "os"
    "log"
    _ "encoding/json"
    "net/http"

    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"

    . "github.com/TasosFrago/epms/db_connection"
    . "github.com/TasosFrago/epms/models"
)


func main() {
    if _, err := os.Stat("../.env"); err == nil {
        err := godotenv.Load("../.env")
        if err != nil {
            fmt.Printf("Error loading environment vars: %s", err)
        }
    }

    config := CredentialConfig{
        Usrname: os.Getenv("USERNAME"),
        Passwd: os.Getenv("PASSWORD"),
        ServerHost: os.Getenv("HOST"),
        ServerPort: os.Getenv("PORT"),
        DBHost: "localhost:3306",
        DBName: "lab2425omada1_EPMS",
    }

    db, err := ConnectDBoSSH(config) 
    if err != nil {
        log.Fatalf("Error connecting to db: %v", err)
    }
    defer db.Cleanup()
    
    // for _, consumer := range consumers {
    //     jsonData, err := json.MarshalIndent(consumer, "", "  ")
    //     if err != nil {
    //         fmt.Println("Error marshalling to JSON:", err)
    //         continue
    //     }
    //     fmt.Println(string(jsonData))
    // }

    r := gin.Default()

    r.Use(func(c *gin.Context) {
        c.Set("db", db.Conn)
        c.Next()
    })

    r.GET("/", func(c *gin.Context) {
        c.JSON(http.StatusOK, gin.H{
            "msg": "hello world",
        })
    })
    r.GET("/consumers", getConsumers)

    r.Run("0.0.0.0:8080")

}

func getConsumers(c *gin.Context) {
    db, ok := c.MustGet("db").(*sql.DB)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not db connection"})
    }

    consumers, err := consumerData(db)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err})
    }
    c.JSON(http.StatusOK, consumers)
}

func consumerData(db *sql.DB) ([]Consumer, error) {
    var consumers []Consumer

    rows, err := db.Query("SELECT * FROM CONSUMER")
    if err != nil {
        return nil, fmt.Errorf("consumerData: %w", err)
    }
    defer rows.Close()

    for rows.Next() {
        var cons Consumer
        if err := rows.Scan(
            &cons.ID,
            &cons.FirstName,
            &cons.LastName,
            &cons.Email,
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
