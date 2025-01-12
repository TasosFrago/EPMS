package paysEndpoint

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TasosFrago/epms/router/middleware"
	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/utls/types"

	"github.com/gorilla/mux"
)

type PaysHandler struct {
	dbSession *sql.DB
}

func NewPaysHandler(db *sql.DB) *PaysHandler {
	return &PaysHandler{
		dbSession: db,
	}
}

func AddPaysSubRouter(router *mux.Router, db *sql.DB) *mux.Router {
	// Define endpoints subrouter
	subRouter := router.PathPrefix("/").Subrouter()

	payHandl := NewPaysHandler(db)

	subRouter.Use(middleware.AuthMiddleware)

	subRouter.HandleFunc("/consumer/{user_id}/meters/{supply_id}/pays/", payHandl.PayProvider).Methods("POST")

	return subRouter
}

type PayData struct {
	User     int     `json:"user,omitempty"`
	Provider string  `json:"provider,omitempty"`
	SupplyID int64   `json:"supply_id,omitempty"`
	Amount   float32 `json:"amount,omitempty"`
}

func (h PaysHandler) PayProvider(w http.ResponseWriter, r *http.Request) {
	consumerDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok && consumerDetails.Type != types.CONSUMER {
		httpError.UnauthorizedError(w, "Create meter, unauthorized user, inalid user type")
		return
	}

	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Pay provider, failed to convert user_id string to int:\n\t%v", err))
		return
	}
	if consumerDetails.ID != user_id {
		httpError.UnauthorizedError(w, "Pay provider, unauthorized user.")
		return
	}

	supply_id, err := strconv.Atoi(mux.Vars(r)["supply_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Pay provider, failed to convert string to int:\n\t%v", err))
		return
	}

	var payDetails PayData
	err = json.NewDecoder(r.Body).Decode(&payDetails)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Pay provider, invalid JSON: \n\t%v", err))
		return
	}
	if payDetails.Amount <= 0 {
		httpError.BadRequestError(w, "Pay provider, invalid request")
		return
	}
	payDetails.User = user_id
	payDetails.SupplyID = int64(supply_id)
	err = payProvider(h.dbSession, r.Context(), payDetails)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Pay provider, internal server error: \n\t%v", err))
		return
	}

	httpError.StatusCreated(w, "Payed", nil)
}

func payProvider(dbSession *sql.DB, ctx context.Context, pay PayData) error {
	// amount := sql.NullFloat64{Float64: float64(*pay.Amount), Valid: true}
	_, err := dbSession.ExecContext(
		ctx,
		`
		INSERT INTO PAYS
		(user, provider, supply_id, amount)
		VALUES
		(?, ?, ?, ?)
		`,
		pay.User,
		pay.Provider,
		pay.SupplyID,
		pay.Amount,
	)
	if err != nil {
		return err
	}
	return nil
}
