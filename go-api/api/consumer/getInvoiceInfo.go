package consumerEndpoint

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TasosFrago/epms/models"
	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/utls/types"

	"github.com/gorilla/mux"
)

func (h ConsumerHandler) GetInvoiceInfo(w http.ResponseWriter, r *http.Request) {
	consumerDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok || consumerDetails.Type != types.CONSUMER {
		httpError.UnauthorizedError(w, "Get Invoice Info, unauthorized user.")
		return
	}

	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Invoice Info, failed to convert user_id string to int:\n\t%v", err))
		return
	}
	if consumerDetails.ID != user_id {
		httpError.UnauthorizedError(w, "Get Invoice Info, unauthorized user.")
		return
	}

	invoice_id, err := strconv.Atoi(mux.Vars(r)["supply_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Invoice Info, failed to convert invoice_id string to int:\n\t%v", err))
		return
	}

	invoice, err := invoiceInfo(h.dbSession, r.Context(), user_id, invoice_id)
	if err != nil {
		if errors.Is(err, errUnauthorized) {
			httpError.UnauthorizedError(w, "Get Invoice Info, access denied to user")
		} else {
			httpError.InternalServerError(w, fmt.Sprintf("Get Invoice Info, failed to get invoice:\n\t%v", err))
		}
		return
	}

	jsonBytes, err := json.Marshal(invoice)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Invoice Info, failed to marshal json:\n\t%v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func invoiceInfo(dbSession *sql.DB, ctx context.Context, user_id int, invoice_id int) (models.Invoice, error) {
	var invoice models.Invoice
	row := dbSession.QueryRowContext(
		ctx,
		`SELECT invoice_id, provider, meter, current_cost, total, name, month, year, receiver
		FROM METER, PLAN
		WHERE invoice_id = ? AND plan_id = plan;`,
		invoice_id,
	)

	err := row.Scan(
		&invoice.ID,
		&invoice.Provider,
		&invoice.Meter,
		&invoice.CurrentCost,
		&invoice.Total,
		&invoice.PlanName,
		&invoice.Month,
		&invoice.Year,
		&invoice.Receiver,
	)
	if err != nil {
		return models.Invoice{}, err
	}

	if invoice.Receiver != user_id {
		return models.Invoice{}, errUnauthorized
	}

	return invoice, nil
}
