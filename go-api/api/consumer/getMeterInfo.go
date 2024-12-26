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

var errUnauthorized = errors.New("unauthorized access")

func (h ConsumerHandler) GetMeterInfo(w http.ResponseWriter, r *http.Request) {
	consumerDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok || consumerDetails.Type != types.CONSUMER {
		httpError.UnauthorizedError(w, "Get Meter Info, unauthorized user.")
		return
	}

	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Meter Info, failed to convert user_id string to int:\n\t%v", err))
		return
	}
	if consumerDetails.ID != user_id {
		httpError.UnauthorizedError(w, "Get Meter Info, unauthorized user.")
		return
	}

	supply_id, err := strconv.Atoi(mux.Vars(r)["supply_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Meter Info, failed to convert supply_id string to int:\n\t%v", err))
		return
	}

	meter, err := meterInfo(h.dbSession, r.Context(), user_id, supply_id)
	if err != nil {
		if errors.Is(err, errUnauthorized) {
			httpError.UnauthorizedError(w, "Get Meter Info, access denied to user")
		} else {
			httpError.InternalServerError(w, fmt.Sprintf("Get Meter Info, failed to get meter:\n\t%v", err))
		}
		return
	}

	jsonBytes, err := json.Marshal(meter)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Meter Info, failed to marshal json:\n\t%v", err))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func meterInfo(dbSession *sql.DB, ctx context.Context, user_id int, supply_id int) (models.Meter, error) {
	var meter models.Meter
	row := dbSession.QueryRowContext(
		ctx,
		`SELECT plan, status, kWh, address, rated_power, department, owner
		FROM METER
		WHERE supply_id = ?;`,
		supply_id,
	)

	err := row.Scan(
		&meter.Plan,
		&meter.Status,
		&meter.KWH,
		&meter.Address,
		&meter.RatedPower,
		&meter.Department,
		&meter.Owner,
	)
	if err != nil {
		return models.Meter{}, err
	}

	if meter.Owner != user_id {
		return models.Meter{}, errUnauthorized
	}

	return meter, nil
}
