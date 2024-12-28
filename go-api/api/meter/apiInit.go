package meterEndpoint

import (
	"database/sql"

	"github.com/TasosFrago/epms/router/middleware"

	"github.com/gorilla/mux"
)

type MeterHandler struct {
	dbSession *sql.DB
}

func NewMeterHandler(db *sql.DB) *MeterHandler {
	return &MeterHandler{
		dbSession: db,
	}
}

func AddMeterSubRouter(router *mux.Router, db *sql.DB) *mux.Router {
	subRouter := router.PathPrefix("/meters").Subrouter()
	subRouterP := router.PathPrefix("/meter").Subrouter()

	meterHandl := NewMeterHandler(db)

	subRouter.Use(middleware.AuthMiddleware)
	subRouterP.Use(middleware.AuthMiddleware)

	subRouter.HandleFunc("/{supply_id}/", meterHandl.GetMeterInfo).Methods("GET")
	subRouterP.HandleFunc("/", meterHandl.CreateMeter).Methods("POST")
	subRouter.HandleFunc(("/"), meterHandl.GetMeterList).Methods("GET")

	subRouter = subRouter.PathPrefix("/{supply_id}").Subrouter()
	return subRouter
}
