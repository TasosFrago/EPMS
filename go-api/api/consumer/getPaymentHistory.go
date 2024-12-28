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

func (h ConsumerHandler) GetPaymentHistory(w http.ResponseWriter, r *http.Request) {
	consumerDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok || consumerDetails.Type != types.CONSUMER {
		httpError.UnauthorizedError(w, "Get Payment History, unauthorized user.")
		return
	}

	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Payment History, failed to convert string to int:\n\t%v", err))
		return
	}
	if consumerDetails.ID != user_id {
		httpError.UnauthorizedError(w, "Get Payment History, unauthorized user.")
		return
	}

	payments, err := paymentHistory(h.dbSession, r.Context(), user_id)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Payment History, failed to get payments:\n\t%v", err))
		return
	}

	jsonBytes, err := json.Marshal(payments)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Payment History, failed to marshal json:\n\t%v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func paymentHistory(dbSession *sql.DB, ctx context.Context, user_id int) ([]models.Pays, error) {
	var payments []models.Pays
	rows, err := dbSession.QueryContext(
		ctx,
		`
		SELECT payment_id, supply_id, provider, amount
		FROM PAYS
		WHERE user = ?;`,
		user_id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var payment models.Pays
		err := rows.Scan(
			&payment.ID,
			&payment.SupplyID,
			&payment.Provider,
			&payment.Amount,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}
