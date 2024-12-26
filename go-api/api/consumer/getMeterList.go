package consumerEndpoint

import (
	"net/http"
	"database/sql"
	"context"
	"fmt"
	"strconv"
	"encoding/json"

	"github.com/TasosFrago/epms/utls/types"
	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/models"

	"github.com/gorilla/mux"
)

func (h ConsumerHandler) GetMeterList(w http.ResponseWriter, r *http.Request) {
	consumerDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok || consumerDetails.Type != types.CONSUMER {
		httpError.UnauthorizedError(w, "Get Meter List, unauthorized user.")
		return
	}
	
	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Meter List, failed to convert string to int:\n\v%v", err))
		return
	}
	if consumerDetails.ID != user_id {
		httpError.UnauthorizedError(w, "Get Meter List, anauthorized user.")
	}

	meters, err := meterList(h.dbSession, r.Context(), user_id)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Meter List, failed to get meters:\n\t%v", err))
		return
	}

	jsonBytes, err := json.Marshal(meters)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Meter List, failed to marshal json:\n\t%v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func meterList(dbSession *sql.DB, ctx context.Context, user_id int) ([]models.Meter, error) {
	var meters []models.Meter
	rows, err := dbSession.QueryContext(
		ctx,
		`
		SELECT supply_id, status, address
		FROM METER
		WHERE owner = ?;`,
		user_id,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var meter models.Meter
		err := rows.Scan(
			&meter.ID,
			&meter.Status,
			&meter.Address,
		)
		if err != nil {
			return nil, err
		}
		meters = append(meters, meter)
	}
	
	return meters, nil
}