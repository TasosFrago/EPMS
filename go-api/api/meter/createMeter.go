package meterEndpoint

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/TasosFrago/epms/api"
	"github.com/TasosFrago/epms/models"
	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/utls/types"

	"github.com/gorilla/mux"
)

// /consumer/{user_id}/meter/
func (h MeterHandler) CreateMeter(w http.ResponseWriter, r *http.Request) {
	var meter models.Meter
	err := json.NewDecoder(r.Body).Decode(&meter)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Create meter, invalid JSON: \n\t%v", err))
		return
	}

	consumerDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok || consumerDetails.Type != types.CONSUMER {
		httpError.UnauthorizedError(w, "Create meter, unauthorized user, inalid user type")
		return
	}

	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Create meter, failed to convert user_id string to int:\n\t%v", err))
		return
	}
	if consumerDetails.ID != user_id {
		httpError.UnauthorizedError(w, "Create meter, unauthorized user.")
		return
	}
	meter.Owner = user_id

	supply_id, err := createMeter(h.dbSession, r.Context(), meter)
	if err != nil {
		if errors.Is(err, apiHelper.ErrBadReq) {
			httpError.BadRequestError(w, "Create meter, invalid Department given")
			return
		} else {
			httpError.InternalServerError(w, fmt.Sprintf("Create meter: \n\t %v", err))
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	httpError.StatusCreated(w, "Created Meter", map[string]int{
		"supply_id": supply_id,
	})
}

func createMeter(dbSession *sql.DB, ctx context.Context, meter models.Meter) (int, error) {
	if meter.Department == "" {
		return 0, apiHelper.ErrBadReq
	}

	var status int
	if meter.Status != nil {
		if *meter.Status {
			status = 1
		} else {
			status = 0
		}
	} else {
		status = 0
	}
	var kWh sql.NullInt64
	if meter.KWH != nil {
		kWh = sql.NullInt64{Int64: int64(*meter.KWH), Valid: true}
	} else {
		kWh = sql.NullInt64{Valid: false}
	}
	fmt.Printf("%v\n", kWh)

	res, err := dbSession.ExecContext(
		ctx,
		`
		INSERT INTO METER
		(status, kWh, address, rated_power, owner, department)
		VALUES
		(?, ?, ?, ?, ?, ?)
		`,
		status,
		kWh,
		meter.Address,
		meter.RatedPower,
		meter.Owner,
		meter.Department,
	)
	if err != nil {
		return 0, err
	}

	supply_id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(supply_id), nil
}
