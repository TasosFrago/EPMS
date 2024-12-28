package authEndpoint

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/utls/security"
	"github.com/TasosFrago/epms/utls/types"
)

type ConsumerLoginReqData struct {
	ID       int    `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (h AuthHandler) LogInConsumer(w http.ResponseWriter, r *http.Request) {
	var loginData ConsumerLoginReqData
	err := json.NewDecoder(r.Body).Decode(&loginData)
	if err != nil {
		httpError.InternalServerError(w, "Login of consumer, invalid JSON")
		return
	}

	consumerData, err := getConsumerByEmail(h.dbSession, loginData.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			httpError.UnauthorizedError(w, "Login of consumer, invalid email")
		} else {
			httpError.InternalServerError(w, fmt.Sprintf("Login of consumer, error getting consumer:\n\t%v", err))
		}
		return
	}

	if !security.CheckPassword(loginData.Password, consumerData.Password) {
		httpError.UnauthorizedError(w, "Login of consumer, invalid password")
		return
	}

	token, err := security.CreateToken(consumerData.ID, loginData.Email, types.CONSUMER, nil)
	if err != nil {
		httpError.InternalServerError(w, "Login of consumer, couldn't create token")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	httpError.StatusCreated(w, "Login successfull", map[string]interface{}{
		"user_id": consumerData.ID,
		"token":   token,
	})
}

func getConsumerByEmail(dbSession *sql.DB, email string) (ConsumerLoginReqData, error) {
	var consumer ConsumerLoginReqData

	row := dbSession.QueryRow("SELECT user_id, email, password FROM CONSUMER WHERE email = ?", email)

	err := row.Scan(&consumer.ID, &consumer.Email, &consumer.Password)
	if err != nil {
		return ConsumerLoginReqData{}, err
	}

	return consumer, nil
}
