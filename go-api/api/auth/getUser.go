package authEndpoint

import (
	"encoding/json"
	"net/http"

	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/utls/types"
)

func (h AuthHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	userDetails, ok := r.Context().Value(types.AuthDetailsKey).(types.AuthDetails)
	if !ok {
		httpError.UnauthorizedError(w, "Get User info, unauthorized user.")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"user_id":   userDetails.ID,
		"user_type": int(userDetails.Type),
	})
}
