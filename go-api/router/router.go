package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/TasosFrago/epms/api/auth"
	"github.com/TasosFrago/epms/api/consumer"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   DBStore
}

type DBStore struct {
	Conn *sql.DB
}

func NewServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   DBStore{db},
	}
}

func (a *APIServer) Run() error {
	router := mux.NewRouter().StrictSlash(true)

	mainRouter := router.PathPrefix("/api/v1/").Subrouter()

	mainRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GET Request Received"))
	}).Methods("GET")

	authEndpoint.AddAuthSubRouter(mainRouter, a.db.Conn)
	consumerEndpoint.AddConsumerSubRouter(mainRouter, a.db.Conn)

	LogAvailableEndpoints(router)

	loggedRouter := LoggingMiddleware(router)

	loggingColor := color.New(color.FgCyan).SprintFunc()
	log.Printf("%s\n", loggingColor("Starting server on "+a.addr+"..."))
	return http.ListenAndServe(a.addr, loggedRouter)
}
