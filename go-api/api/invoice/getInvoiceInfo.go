package invoiceEndpoint

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TasosFrago/epms/api"
	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/utls/types"

	"github.com/gorilla/mux"
)

func (h InvoiceHandler) GetInvoiceInfo(w http.ResponseWriter, r *http.Request) {
	consumerDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok && consumerDetails.Type != types.CONSUMER {
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

	invoice_id, err := strconv.Atoi(mux.Vars(r)["invoice_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Invoice Info, failed to convert invoice_id string to int:\n\t%v", err))
		return
	}

	invoice, err := invoiceInfo(h.dbSession, r.Context(), user_id, invoice_id)
	if err != nil {
		if errors.Is(err, apiHelper.ErrUnauthorized) {
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

func invoiceInfo(dbSession *sql.DB, ctx context.Context, user_id int, invoice_id int) (apiHelper.InvoicePayment, error) {
	var invoice apiHelper.InvoicePayment
	row := dbSession.QueryRowContext(
		ctx,
		`SELECT invoice_id, INVOICE.provider, meter, current_cost, total, name, month, year, receiver
		FROM INVOICE, PLAN
		WHERE invoice_id = ? AND plan_id = plan AND receiver = ?;`,
		invoice_id,
		user_id,
	)

	err := row.Scan(
		&invoice.ID,
		&invoice.Provider,
		&invoice.Meter,
		&invoice.CurrentCost,
		&invoice.Total,
		&invoice.Name,
		&invoice.Month,
		&invoice.Year,
		&invoice.Receiver,
	)
	if err == sql.ErrNoRows {
		return apiHelper.InvoicePayment{}, apiHelper.ErrUnauthorized
	}
	if err != nil {
		return apiHelper.InvoicePayment{}, err
	}

	return invoice, nil
}
