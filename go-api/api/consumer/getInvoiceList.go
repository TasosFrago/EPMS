package consumerEndpoint

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TasosFrago/epms/models"
	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/utls/types"

	"github.com/gorilla/mux"
)

func (h ConsumerHandler) GetInvoiceList(w http.ResponseWriter, r *http.Request) {
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

	invoices, err := invoiceList(h.dbSession, r.Context(), user_id)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Invoice List, failed to get invoices:\n\t%v", err))
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

func invoiceList(dbSession *sql.DB, ctx context.Context, user_id int) ([]models.Invoice, error) {
	var invoices []models.Invoice
	rows, err := dbSession.QueryContext(
		ctx,
		`
		SELECT invoice_id, provider, current_cost, month, year
		FROM INVOICE, PLAN
		WHERE receiver = ? AND plan = plan_id;`,
		user_id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var invoice models.Invoice
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
