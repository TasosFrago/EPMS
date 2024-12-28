package authEndpoint

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/TasosFrago/epms/models"
	"github.com/TasosFrago/epms/utls"
	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/utls/security"
)

func (h AuthHandler) SignUpCons(w http.ResponseWriter, r *http.Request) {
	var consumer models.Consumer
	err := json.NewDecoder(r.Body).Decode(&consumer)
	if err != nil {
		httpError.InternalServerError(w, "Sign up of consumer, invalid JSON")
		return
	}

	err = validateConsumer(consumer)
	if err != nil {
		httpError.BadRequestError(w, fmt.Sprintf("Sign up of consumer, invalid data given:\n\t%v", err))
		return
	}

	consumer.Password, err = security.HashPassword(consumer.Password)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Sign up of consumer, failed hashing password:\n\t%v", err))
		return
	}

	err = createConsumer(h.dbSession, consumer)
	if err != nil {
		httpError.ConflictError(w, fmt.Sprintf("Sign up of consumer, failed creating consumer:\n\t%v", err))
		return
	}

	httpError.StatusCreated(w, "create consumer successfully", nil)
}

func createConsumer(dbSession *sql.DB, consumer models.Consumer) error {
	var err error
	var query string

	if consumer.Landline == nil {
		query = `
		INSERT INTO CONSUMER
		(first_name, last_name, email, cell, password)
		VALUES
		(?, ?, ?, ?, ?)
		`
		_, err = dbSession.Exec(
			query,
			consumer.FirstName,
			consumer.LastName,
			consumer.Email,
			consumer.Cell,
			consumer.Password,
		)
	} else {
		query = `
		INSERT INTO CONSUMER
		(first_name, last_name, email, cell, password, landline)
		VALUES
		(?, ?, ?, ?, ?, ?)
		`
		_, err = dbSession.Exec(
			query,
			consumer.FirstName,
			consumer.LastName,
			consumer.Email,
			consumer.Cell,
			consumer.Password,
			*consumer.Landline,
		)
	}
	if err != nil {
		return err
	}
	return nil
}

func validateConsumer(consumer models.Consumer) error {
	if strings.TrimSpace(consumer.FirstName) == "" {
		return fmt.Errorf("First name is required")
	}
	if strings.TrimSpace(consumer.LastName) == "" {
		return fmt.Errorf("Last name is required")
	}
	if strings.TrimSpace(consumer.Email) == "" {
		return fmt.Errorf("Email is required")
	}
	if !utls.IsValidEmail(strings.TrimSpace(consumer.Email)) {
		return fmt.Errorf("Email is not valid")
	}
	if strings.TrimSpace(consumer.Password) == "" {
		return fmt.Errorf("Password is required")
	}
	if !utls.IsValidPassword(strings.TrimSpace(consumer.Password)) {
		return fmt.Errorf("Password is not valid")
	}
	// TODO: Need to add cell and landline phone validator
	if strings.TrimSpace(consumer.Cell) == "" {
		return fmt.Errorf("Cell is required")
	}
	return nil
}
