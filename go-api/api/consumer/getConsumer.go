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

func (h ConsumerHandler) GetConsumerInfo(w http.ResponseWriter, r *http.Request) {
	consumerDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok && consumerDetails.Type != types.CONSUMER {
		httpError.UnauthorizedError(w, "Get Consumer info, unauthorized user.")
		return
	}

	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Consumer info, failed to convert string to int:\n\t%v", err))
		return
	}
	if consumerDetails.ID != user_id {
		httpError.UnauthorizedError(w, "Get Consumer info, unauthorized user.")
		return
	}

	consumer, err := consumerInfo(h.dbSession, r.Context(), user_id)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Consumer info, failed to get consumer:\n\t%v", err))
		return
	}

	jsonBytes, err := json.Marshal(consumer)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Consumer info, failed to marshal json:\n\t%v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func consumerInfo(dbSession *sql.DB, ctx context.Context, user_id int) (models.Consumer, error) {
	var consumer models.Consumer
	row := dbSession.QueryRowContext(
		ctx,
		`SELECT first_name, last_name, email, cell, landline
		FROM CONSUMER
		WHERE user_id = ?;`,
		user_id,
	)

	err := row.Scan(
		&consumer.FirstName,
		&consumer.LastName,
		&consumer.Email,
		&consumer.Cell,
		&consumer.Landline,
	)
	if err != nil {
		return models.Consumer{}, err
	}
	return consumer, nil
}
