package models

type Consumer struct {
	ID         int64   `json:"user_id,omitempty"`
	FirstName  string  `json:"first_name,omitempty"`
	LastName   string  `json:"last_name,omitempty"`
	Email      string  `json:"email,omitempty"`
	Password   string  `json:"password,omitempty"`
	Cell       string  `json:"cell,omitempty"`
	Landline   *string `json:"landline,omitempty"`
	CreditInfo *int    `json:"credit_info,omitempty"`
}
