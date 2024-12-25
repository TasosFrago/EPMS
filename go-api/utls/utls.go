package utls

import (
	"log"
	"os"
	"regexp"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if _, err := os.Stat("../.env"); err == nil {
		err := godotenv.Load("../.env")
		if err != nil {
			log.Fatalf("Error loading environment vars: %s", err)
		}
	}
}

func IsValidEmail(email string) bool {
	re := regexp.MustCompile(`^[a-z0-9._%+-]+@[a-z0-9.-]+\.[a-z]{2,}$`) // Using magic to find if the email is valid
	return re.MatchString(email)
}

func IsValidPassword(password string) bool {
	return len(password) >= 8
}
