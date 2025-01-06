package planEndpoint

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/TasosFrago/epms/api"
	"github.com/TasosFrago/epms/models"
	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/utls/settings"
	"github.com/TasosFrago/epms/utls/types"

	"github.com/gorilla/mux"
)

func (h PlanHandler) GetPlans(w http.ResponseWriter, r *http.Request) {
	consumerDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok && consumerDetails.Type != types.CONSUMER {
		httpError.UnauthorizedError(w, "Get plans list, unauthorized user.")
		return
	}

	user_id, err := strconv.Atoi(mux.Vars(r)["user_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get plans List, failed to convert string to int:\n\t%v", err))
		return
	}
	if consumerDetails.ID != user_id {
		httpError.UnauthorizedError(w, "Get plans list, unauthorized user.")
		return
	}

	supply_id, err := strconv.Atoi(mux.Vars(r)["supply_id"])
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get plans List, failed to convert string to int:\n\t%v", err))
		return
	}

	// Default values
	defaultPage := 1
	defaultLimit := 50

	// Get optional query parameters
	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	// Convert page to integer, use default if empty or invalid
	page := defaultPage
	if pageStr != "" {
		if parsedPage, err := strconv.Atoi(pageStr); err == nil && parsedPage > 0 {
			page = parsedPage
		}
	}

	// Convert limit to integer, use default if empty or invalid
	limit := defaultLimit
	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil && parsedLimit > 0 {
			limit = parsedLimit
		}
	}

	plans, err := getPlans(h.dbSession, r.Context(), user_id, supply_id, page, limit)
	if err != nil {
		if errors.Is(err, apiHelper.ErrUnauthorized) {
			httpError.UnauthorizedError(w, "Get plans List, Unauthorized user")
		} else {
			httpError.InternalServerError(w, fmt.Sprintf("Get plans List, internal server error:\n\t%v", err))
		}
		return
	}

	jsonBytes, err := json.Marshal(plans)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Get plans List, internal server error:\n\t%v", err))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)
}

func getPlans(dbSession *sql.DB, ctx context.Context, user_id int, supply_id int, page int, limit int) ([]models.Plan, error) {
	plans := []models.Plan{}
	plan := models.Plan{}

	err := dbSession.QueryRowContext(
		ctx,
		`
		SELECT
		plan_id, type, price, name, provider, STR_TO_DATE(CONCAT('1-', month, '-', year), '%d-%M-%Y') AS issue_date, duration
		FROM PLAN, METER
		WHERE plan = plan_id AND supply_id = ? AND owner = ?
		`,
		supply_id,
		user_id,
	).Scan(
		&plan.ID,
		&plan.Type,
		&plan.Price,
		&plan.Name,
		&plan.Provider,
		&plan.IssueDate,
		&plan.Duration,
	)

	if err == sql.ErrNoRows {
		row2 := dbSession.QueryRowContext(
			ctx,
			`
			SELECT * FROM METER WHERE owner = ? AND supply_id = ?; 
			`,
			user_id,
			supply_id,
		)
		if row2.Err() == sql.ErrNoRows {
			return nil, apiHelper.ErrUnauthorized
		} else if row2.Err() != nil {
			return nil, row2.Err()
		}
		// Else we need to select a plan
	} else if err != nil {
		return nil, err
	}

	now := settings.GetCurrentDate()
	currentYear, currentMonth, _ := now.Date()

	const layout = "2006-01-02"
	parseDate, err := time.Parse(layout, *plan.IssueDate)
	if err != nil {
		// Internal error
		return nil, err
	}

	issueYear, issueMonth, _ := parseDate.Date()
	if issueYear == currentYear && issueMonth == currentMonth {
		// NOTE: Has selected plan so empty list
		fmt.Println("Years and months match")
		return plans, nil
	}
	if parseDate.After(now) {
		// plan after this month should not be selected
		return nil, fmt.Errorf("unexpected error, plan %d cannot be selected", plan.ID)
	}
	if parseDate.Before(now) {
		expiryDate := parseDate.AddDate(0, plan.Duration, 0)

		if expiryDate.After(now) {
			// NOTE: Has plan with duration exceeding this month so empty list
			fmt.Println("Does not expire until")
			return plans, err
		}
	}

	offset := (page - 1) * limit
	rows, err := dbSession.QueryContext(
		ctx,
		`
		SELECT
		plan_id, type, price, name, provider, STR_TO_DATE(CONCAT('1-', month, '-', year), '%d-%M-%Y') AS issue_date, duration
		FROM PLAN
		WHERE month = ?
		LIMIT ? OFFSET ?
		`,
		currentMonth.String(),
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var tmpPlan models.Plan
		err := rows.Scan(
			&tmpPlan.ID,
			&tmpPlan.Type,
			&tmpPlan.Price,
			&tmpPlan.Name,
			&tmpPlan.Provider,
			&tmpPlan.IssueDate,
			&tmpPlan.Duration,
		)
		if err != nil {
			return nil, err
		}
		plans = append(plans, tmpPlan)
	}
	return plans, nil
}
