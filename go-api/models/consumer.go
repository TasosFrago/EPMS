package models

type Consumer struct {
	ID         int64   `json:"user_id"`
	FirstName  string  `json:"first_name"`
	LastName   string  `json:"last_name"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
	Cell       string  `json:"cell"`
	Landline   *string `json:"landline"`
	CreditInfo *int    `json:"credit_info"`
}
