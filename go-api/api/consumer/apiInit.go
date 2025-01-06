package consumerEndpoint

import (
	"database/sql"

	"github.com/TasosFrago/epms/router/middleware"

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

func AddConsumerSubRouter(router *mux.Router, db *sql.DB) *mux.Router {
	// Define endpoints subrouter
	subRouter := router.PathPrefix("/consumer/{user_id}").Subrouter()

	consHandl := NewConsumerHandler(db)

	privateRouter := subRouter.PathPrefix("/").Subrouter()
	privateRouter.Use(middleware.AuthMiddleware)

	// Defining Protected routes
	privateRouter.HandleFunc("/", consHandl.GetConsumerInfo).Methods("GET")

	privateRouter.HandleFunc("/{user_id}/meters/{supply_id}/plans", consHandl.GetAvailablePlans).Methods("GET")
	// TODO: Move available plans list under /plans. (Does not need to be a protected route)

	privateRouter.HandleFunc("/payments", consHandl.GetPaymentHistory).Methods("GET")

	return subRouter
}
