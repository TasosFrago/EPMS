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

func (h InvoiceHandler) GetInvoiceList(w http.ResponseWriter, r *http.Request) {
	consumerDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok || consumerDetails.Type != types.CONSUMER {
		httpError.UnauthorizedError(w, "Get Invoice List, unauthorized user.")
		return
	}

	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Meter List, failed to convert string to int:\n\t%v", err))
		return
	}
	if consumerDetails.ID != user_id {
		httpError.UnauthorizedError(w, "Get Invoice List, unauthorized user.")
		return
	}

	// Check if there is a supply_id in the route
	var supply_id *int
	supply_id_str, exists := mux.Vars(r)["supply_id"]
	if exists {
		supply_id = new(int)
		*supply_id, err = strconv.Atoi(supply_id_str)
		if err != nil {
			httpError.InternalServerError(w, fmt.Sprintf("Get Unpaid Invoice List, failed to convert string to int:\n\t%v", err))
			return
		}
	}

	invoices, err := invoiceList(h.dbSession, r.Context(), user_id, supply_id)
	if err != nil {
		if errors.Is(err, apiHelper.ErrUnauthorized) {
			httpError.UnauthorizedError(w, "Get Invoice List, unauthorized user.")
		} else {
			httpError.InternalServerError(w, fmt.Sprintf("Get Invoice List, failed to get invoices:\n\t%v", err))
		}
		return
	}

	jsonBytes, err := json.Marshal(invoices)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Invoice List, failed to marshal json:\n\t%v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func invoiceList(dbSession *sql.DB, ctx context.Context, user_id int, supply_id *int) ([]apiHelper.InvoicePayment, error) {
	invoices := []apiHelper.InvoicePayment{}

	var (
		rows *sql.Rows
		err  error
	)
	if supply_id == nil {
		rows, err = dbSession.QueryContext(
			ctx,
			`
			SELECT invoice_id, INVOICE.provider, current_cost, month, year
			FROM INVOICE, PLAN
			WHERE receiver = ? AND plan = plan_id;`,
			user_id,
		)
	} else {
		rows, err = dbSession.QueryContext(
			ctx,
			`
			SELECT invoice_id, INVOICE.provider, current_cost, month, year
			FROM INVOICE, PLAN
			WHERE receiver = ? AND plan = plan_id;`,
			user_id,
		)
	}
	if err == sql.ErrNoRows && supply_id != nil {
		return nil, apiHelper.ErrUnauthorized
	} else if err == sql.ErrNoRows {
		return invoices, nil
	} else if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var invoice apiHelper.InvoicePayment
		err := rows.Scan(
			&invoice.ID,
			&invoice.Provider,
			&invoice.CurrentCost,
			&invoice.Month,
			&invoice.Year,
		)
		if err != nil {
			return nil, err
		}
		invoices = append(invoices, invoice)
	}

	return invoices, nil
}
