package providerEndpoint

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TasosFrago/epms/models"
	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/gorilla/mux"
)

type ProviderHandler struct {
	dbSession *sql.DB
}

func NewProviderHandler(db *sql.DB) *ProviderHandler {
	return &ProviderHandler{
		dbSession: db,
	}
}

func AddProviderHandler(router *mux.Router, db *sql.DB) *mux.Router {
	subRouter := router.PathPrefix("/providers").Subrouter()

	provHandl := NewProviderHandler(db)

	subRouter.HandleFunc("/", provHandl.GetProviders).Methods("GET")
	return subRouter
}

func (h ProviderHandler) GetProviders(w http.ResponseWriter, r *http.Request) {
	providers, err := providers(h.dbSession, r.Context())
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Getting providers: \n\t %v", err))
		return
	}

	jsonBytes, err := json.Marshal(providers)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Getting providers: error marshaling data:\n\t%v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func providers(dbSession *sql.DB, ctx context.Context) ([]models.Provider, error) {
	var providers []models.Provider

	rows, err := dbSession.QueryContext(
		ctx,
		"SELECT name, phone, email FROM PROVIDER",
	)
	if err != nil {
		return nil, fmt.Errorf("provider data: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var prov models.Provider
		if err := rows.Scan(
			&prov.Name,
			&prov.Phone,
			&prov.Email,
		); err != nil {
			return nil, fmt.Errorf("provider data: %w", err)
		}
		providers = append(providers, prov)
	}

	return providers, nil
}
