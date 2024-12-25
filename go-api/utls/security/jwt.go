package security

import (
	"fmt"
	"os"
	"time"

	"github.com/TasosFrago/epms/utls/types"

	"github.com/golang-jwt/jwt/v5"
)

type CustomJWTClaims struct {
	Email   string `json:"email"`
	UsrType int    `json:"usrT"`
	jwt.RegisteredClaims
}

func CreateToken(email string, usr types.UsrType, expirationTime *time.Time) (string, error) {
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		return "", fmt.Errorf("JWT_KEY env variable is not set")
	}

	var exp time.Time
	if expirationTime != nil {
		exp = *expirationTime
	} else {
		exp = time.Now().Add(24 * time.Hour)
	}

	claims := CustomJWTClaims{
		Email:   email,
		UsrType: int(usr),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "epms",
		},
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	tokenString, err := token.SignedString([]byte(jwtKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func VerifyToken(tokenString string) (*CustomJWTClaims, error) {
	jwtKey := os.Getenv("JWT_KEY")
	if jwtKey == "" {
		return nil, fmt.Errorf("JWT_KEY env variable is not set")
	}

	claims := &CustomJWTClaims{}

	token, err := jwt.ParseWithClaims(
		tokenString,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
