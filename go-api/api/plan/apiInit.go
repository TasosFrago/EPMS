package planEndpoint

import (
	"database/sql"

	"github.com/TasosFrago/epms/router/middleware"

	"github.com/gorilla/mux"
)

type PlanHandler struct {
	dbSession *sql.DB
}

func NewPlanHandler(db *sql.DB) *PlanHandler {
	return &PlanHandler{
		dbSession: db,
	}
}

func AddPlanSubRouter(router *mux.Router, db *sql.DB) *mux.Router {
	subRouter := router.PathPrefix("/consumer/{user_id}/meters/{supply_id}").Subrouter()

	planHandl := NewPlanHandler(db)

	privateRouter := subRouter.PathPrefix("/").Subrouter()
	privateRouter.Use(middleware.AuthMiddleware)

	privateRouter.HandleFunc("/plan", planHandl.GetPlans).Methods("GET")

	return router
}
