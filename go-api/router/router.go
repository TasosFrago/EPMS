package router

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/TasosFrago/epms/api/auth"
	"github.com/TasosFrago/epms/api/consumer"
	"github.com/TasosFrago/epms/api/invoice"
	"github.com/TasosFrago/epms/api/meter"
	"github.com/TasosFrago/epms/api/pays"
	"github.com/TasosFrago/epms/api/plan"
	"github.com/TasosFrago/epms/api/provider"

	"github.com/fatih/color"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type APIServer struct {
	addr string
	db   *DBStore
}

type DBStore struct {
	Conn *sql.DB
}

func NewServer(addr string) *APIServer {
	return &APIServer{
		addr: addr,
		db:   nil,
	}
}
func (a *APIServer) SetDB(db *sql.DB) {
	a.db = &DBStore{db}
}

func (a *APIServer) RunWithTemporaryHandlers(server *http.Server) error {
	router := mux.NewRouter().StrictSlash(true)

	// Temporary health check route to say we're waiting for the DB
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Service is starting, waiting for DB connection..."))
	}).Methods("GET")

	server.Handler = router

	// Start the server immediately
	fmt.Printf("Listening on %s...\n", a.addr)
	return server.ListenAndServe()
}

func (a *APIServer) Run() error {
	if a.db == nil || a.db.Conn == nil {
		log.Fatal("Database connection is not set")
	}
	router := mux.NewRouter().StrictSlash(true)

	mainRouter := router.PathPrefix("/api/v1/").Subrouter()

	mainRouter.Methods("OPTIONS").HandlerFunc(SetOptions)

	mainRouter.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("GET Request Received"))
	}).Methods("GET")

	authEndpoint.AddAuthSubRouter(mainRouter, a.db.Conn)
	consumerRouter := consumerEndpoint.AddConsumerSubRouter(mainRouter, a.db.Conn)
	meterRouter := meterEndpoint.AddMeterSubRouter(consumerRouter, a.db.Conn)

	invoiceEndpoint.AddInvoiceConsumerMeterSubRouter(consumerRouter, a.db.Conn)
	invoiceEndpoint.AddInvoiceConsumerMeterSubRouter(meterRouter, a.db.Conn)

	invoiceEndpoint.AddInvoiceSubRouter(consumerRouter, a.db.Conn)

	planEndpoint.AddPlanSubRouter(mainRouter, a.db.Conn)

	providerEndpoint.AddProviderHandler(mainRouter, a.db.Conn)

	paysEndpoint.AddPaysSubRouter(mainRouter, a.db.Conn)

	// Configure CORS
	corsHandler := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:5173", "https://epms-six.vercel.app/", "https://epms-tasosfragos-projects.vercel.app/", "https://epms-git-feature-fe-tasosfragos-projects.vercel.app/"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	LogAvailableEndpoints(router)

	loggedRouter := LoggingMiddleware(router)

	loggingColor := color.New(color.FgCyan).SprintFunc()
	log.Printf("%s\n", loggingColor("Starting server on "+a.addr+"..."))
	return http.ListenAndServe(a.addr, corsHandler(loggedRouter))
}

func SetOptions(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	log.Printf("Received OPTIONS request from origin: %s", origin)

	if origin == "" {
		origin = "*"
	}

	// Allowed origins for credentials
	allowedOrigins := []string{
		"http://localhost:5173",
		"https://epms-six.vercel.app",
		"https://epms-tasosfragos-projects.vercel.app",
		"https://epms-git-feature-fe-tasosfragos-projects.vercel.app",
	}

	// Check if the origin matches an allowed one
	matched := false
	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			matched = true
			w.Header().Set("Access-Control-Allow-Origin", origin)
			break
		}
	}

	// If no match, set default allowed origin
	if !matched {
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	}

	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	log.Printf("CORS Headers set: %+v", w.Header()) // Log the CORS headers
	w.WriteHeader(http.StatusNoContent)             // Respond with 204 No Content
}
