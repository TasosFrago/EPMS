package router

import (
	"database/sql"
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
	consumerRouter := consumerEndpoint.AddConsumerSubRouter(mainRouter, a.db.Conn)
	meterRouter := meterEndpoint.AddMeterSubRouter(consumerRouter, a.db.Conn)

	invoiceEndpoint.AddInvoiceConsumerMeterSubRouter(consumerRouter, a.db.Conn)
	invoiceEndpoint.AddInvoiceConsumerMeterSubRouter(meterRouter, a.db.Conn)

	invoiceEndpoint.AddInvoiceSubRouter(consumerRouter, a.db.Conn)

	planEndpoint.AddPlanSubRouter(mainRouter, a.db.Conn)

	providerEndpoint.AddProviderHandler(mainRouter, a.db.Conn)

	paysEndpoint.AddPaysSubRouter(mainRouter, a.db.Conn)

	LogAvailableEndpoints(router)

	loggedRouter := LoggingMiddleware(router)

	loggingColor := color.New(color.FgCyan).SprintFunc()
	log.Printf("%s\n", loggingColor("Starting server on "+a.addr+"..."))
	return http.ListenAndServe(a.addr, loggedRouter)
}
