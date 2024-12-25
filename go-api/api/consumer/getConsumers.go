package consumerEndpoint

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TasosFrago/epms/models"
	"github.com/TasosFrago/epms/utls/httpError"
)

func (h ConsumerHandler) GetConsumer(w http.ResponseWriter, r *http.Request) {
	consumers, err := consumerData(h.dbSession)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Getting consumer data: %v", err))
	}

	jsonBytes, err := json.Marshal(consumers)
	if err != nil {
		httpError.InternalServerError(w, fmt.Sprintf("Error marshaling data: %v", err))
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func consumerData(db *sql.DB) ([]models.Consumer, error) {
	var consumers []models.Consumer

	rows, err := db.Query("SELECT * FROM CONSUMER")
	if err != nil {
		return nil, fmt.Errorf("consumerData: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var cons models.Consumer
		if err := rows.Scan(
			&cons.ID,
			&cons.FirstName,
			&cons.LastName,
			&cons.Email,
			&cons.Password,
			&cons.Cell,
			&cons.Landline,
			&cons.CreditInfo,
		); err != nil {
			return nil, fmt.Errorf("consumerData: %v", err)
		}
		consumers = append(consumers, cons)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("consumerData: %v", err)
	}
	return consumers, nil
}
