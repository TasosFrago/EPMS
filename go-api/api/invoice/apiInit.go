package invoiceEndpoint

import (
	"database/sql"

	"github.com/TasosFrago/epms/router/middleware"

	"github.com/gorilla/mux"
)

type InvoiceHandler struct {
	dbSession *sql.DB
}

func NewInvoiceHandler(db *sql.DB) *InvoiceHandler {
	return &InvoiceHandler{
		dbSession: db,
	}
}

func AddInvoiceConsumerMeterSubRouter(router *mux.Router, db *sql.DB) {
	subRouter := router.PathPrefix("/invoices").Subrouter()

	invHandl := NewInvoiceHandler(db)

	subRouter.Use(middleware.AuthMiddleware)

	// You need to get {user_id} and {supply_id}
	subRouter.HandleFunc("/unpaid", invHandl.GetInvoiceUnpaid).Methods("GET")
	subRouter.HandleFunc("/", invHandl.GetInvoiceList).Methods("GET")

}

func AddInvoiceSubRouter(router *mux.Router, db *sql.DB) {
	subRouter := router.PathPrefix("/invoice").Subrouter()

	invHandl := NewInvoiceHandler(db)

	subRouter.Use(middleware.AuthMiddleware)

	subRouter.HandleFunc("/{invoice_id}", invHandl.GetInvoiceInfo).Methods("GET")
}
