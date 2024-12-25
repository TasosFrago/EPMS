package authEndpoint

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type AuthHandler struct {
	dbSession *sql.DB
}

func NewAuthHandler(db *sql.DB) *AuthHandler {
	return &AuthHandler{
		dbSession: db,
	}
}

func AddAuthSubRouter(router *mux.Router, db *sql.DB) error {
	subRouter := router.PathPrefix("/auth").Subrouter()

	authHandl := NewAuthHandler(db)

	subRouter.HandleFunc("/signup/consumer", authHandl.SignUpCons).Methods("POST")
	subRouter.HandleFunc("/login/consumer", authHandl.LogInConsumer).Methods("POST")
	// subRouter.HandleFunc("/login/provider", authHandl.LogIn).Methods("POST")
	// subRouter.HandleFunc("/login/employee", authHandl.LogIn).Methods("POST")

	return nil
}
