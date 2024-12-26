package middleware

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/TasosFrago/epms/utls/httpError"
	"github.com/TasosFrago/epms/utls/security"
	"github.com/TasosFrago/epms/utls/types"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			httpError.UnauthorizedError(w, "Unauthorized user, no token passed.")
			return
		}

		// Remove Bearer string from token
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)

		claims, err := security.VerifyToken(tokenString)
		if err != nil {
			httpError.UnauthorizedError(w, fmt.Sprintf("Unauthorized user, invalid token:\n\t%v", err))
			return
		}

		authDetails := types.AuthDetails{
			ID: claims.ID,
			Email: claims.Email,
			Type:  types.UsrType(claims.UsrType),
		}

		ctx := context.WithValue(r.Context(), types.AuthDetailsKey, authDetails)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}