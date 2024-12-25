package consumerEndpoint

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TasosFrago/epms/models"
	"github.com/TasosFrago/epms/router/middleware"
	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/utls/types"

	"github.com/gorilla/mux"
)

type ConsumerHandler struct {
	dbSession *sql.DB
}

func NewConsumerHandler(db *sql.DB) *ConsumerHandler {
	return &ConsumerHandler{
		dbSession: db,
	}
}

func AddConsumerSubRouter(router *mux.Router, db *sql.DB) error {
	// Define endpoints subrouter
	subRouter := router.PathPrefix("/consumer").Subrouter()

	consHandl := NewConsumerHandler(db)

	// Define are endpoints for /consumer
	subRouter.HandleFunc("/", consHandl.GetConsumer).Methods("GET")

	privateRouter := subRouter.PathPrefix("/").Subrouter()
	privateRouter.Use(middleware.AuthMiddleware)

	privateRouter.HandleFunc("/consumerd", consHandl.ProtectedConsumer).Methods("GET")

	return nil
}

func (h ConsumerHandler) ProtectedConsumer(w http.ResponseWriter, r *http.Request) {
	userDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok {
		httpError.UnauthorizedError(w, "Protected Route, could not get user details")
		return
	}
	if userDetails.Type != types.CONSUMER {
		httpError.UnauthorizedError(w, "Protected Route, could not get user details")
		return
	}
	var consumer models.Consumer

	row := h.dbSession.QueryRow("SELECT * FROM CONSUMER WHERE email = ?", userDetails.Email)

	err := row.Scan(&consumer.ID, &consumer.FirstName, &consumer.LastName, &consumer.Email, &consumer.Password, &consumer.Cell, &consumer.Landline, &consumer.CreditInfo)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Protected Route, could not get user info from db:\n\t%v", err))
		return
	}

	jsonBytes, err := json.Marshal(consumer)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Protected Route, could not get user info from db:\n\t%v", err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
