package router

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
}

func NewServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
	}
}

func (a *APIServer) Run() error {
	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GET Request Received"))
	}).Methods("GET")

	LogAvailableEndpoints(router)

	loggedRouter := LoggingMiddleware(router)

	log.Printf("Starting server on %s...\n", a.addr)
	return http.ListenAndServe(a.addr, loggedRouter)
}
