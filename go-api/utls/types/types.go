package types

type UsrType int

const (
	CONSUMER UsrType = iota
	PROVIDER
	EMPLOYEE
)

type AuthDetails struct {
	ID    int
	Email string
	Type  UsrType
}

type contextKey string

const AuthDetailsKey contextKey = "authDetails"
