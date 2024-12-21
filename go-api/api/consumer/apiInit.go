package consumerEndpoint

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type ConsumerHandler struct {
	db *sql.DB
}

func NewConsumerHandler(db *sql.DB) *ConsumerHandler {
	return &ConsumerHandler{
		db: db,
	}
}

func AddConsumerSubRouter(router *mux.Router, db *sql.DB) error {
	// Define endpoints subrouter
	subRouter := router.PathPrefix("/consumer/").Subrouter()

	consHandl := NewConsumerHandler(db)

	// Define are endpoints for /consumer
	subRouter.HandleFunc("/", consHandl.GetConsumer).Methods("GET")

	return nil
}
