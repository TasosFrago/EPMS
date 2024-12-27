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

var errNotAbleToChoose = errors.New("already committed to plan")

func (h ConsumerHandler) GetAvailablePlans(w http.ResponseWriter, r *http.Request) {
	consumerDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok || consumerDetails.Type != types.CONSUMER {
		httpError.UnauthorizedError(w, "Get Available Plans, unauthorized user.")
		return
	}

	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Available Plans, failed to convert user_id string to int:\n\t%v", err))
		return
	}
	if consumerDetails.ID != user_id {
		httpError.UnauthorizedError(w, "Get Available Plans, unauthorized user.")
		return
	}

	supply_id, err := strconv.Atoi(mux.Vars(r)["supply_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Available Plans, failed to convert supply_id string to int:\n\t%v", err))
		return
	}

	plans, err := planList(h.dbSession, r.Context(), user_id, supply_id)
	if err != nil {
		if errors.Is(err, errUnauthorized) {
			httpError.UnauthorizedError(w, "Get Available Plans, access denied to user")
		} else if errors.Is(err, errNotAbleToChoose) {
			httpError.UnprocessableEntityError(w, "Get Available Plans, already committed to plan")
		} else {
			httpError.InternalServerError(w, fmt.Sprintf("Get Available Plans, failed to get plans:\n\t%v", err))
		}
		return
	}

	jsonBytes, err := json.Marshal(plans)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get Available Plans, failed to marshal json:\n\t%v", err))
		return
	}

	w.Header().Set("Content_Type", "application/json")
	w.Write(jsonBytes)
}

func planList(dbSession *sql.DB, ctx context.Context, user_id int, supply_id int) ([]models.Plan, error) {
	var meter models.Meter
	row_meter := dbSession.QueryRowContext(
		ctx,
		`SELECT owner
		FROM METER
		WHERE supply_id = ?;`,
		supply_id,
	)

	err_meter := row_meter.Scan(
		&meter.Owner,
	)
	if err_meter != nil {
		return nil, err_meter
	}

	if meter.Owner != user_id {
		return nil, errUnauthorized
	}

	var current_plan models.Plan
	row_plan := dbSession.QueryRowContext(
		ctx,
		`SELECT month, year, duration
		FROM METER, PLAN
		WHERE plan = plan_id AND supply_id = ?;`,
		supply_id,
	)

	err_plan := row_plan.Scan(
		&current_plan.Month,
		&current_plan.Year,
		&current_plan.Duration,
	)
	if err_plan != nil {
		return nil, err_plan
	}

	if (monthNumber(current_plan.Month, current_plan.Year) + current_plan.Duration) > monthNumber("November", 2024) {
		return nil, errNotAbleToChoose
	}

	var plans []models.Plan
	rows, err := dbSession.QueryContext(
		ctx,
		`
		SELECT plan_id, name, type, provider, price, month, year, duration
		FROM PLAN
		WHERE month = November;`,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var plan models.Plan
		err := rows.Scan(
			&plan.ID,
			&plan.Name,
			&plan.Type,
			&plan.Provider,
			&plan.Price,
			&plan.Month,
			&plan.Year,
			&plan.Duration,
		)
		if err != nil {
			return nil, err
		}
		plans = append(plans, plan)
	}

	return plans, nil
}

func monthNumber(month string, year int) (number int) {
	monthMap := map[string]int{
		"January":   1,
		"February":  2,
		"March":     3,
		"April":     4,
		"May":       5,
		"June":      6,
		"July":      7,
		"August":    8,
		"September": 9,
		"October":   10,
		"November":  11,
		"December":  12,
	}

	number = monthMap[month]
	number = number + ((year - 2024) * 12)

	return number
}
