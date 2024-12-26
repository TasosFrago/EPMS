package models

type Employee struct {
	BadgeID    int64   `json:"badge,omitempty"`
	FirstName  string  `json:"first_name,omitempty"`
	LastName   string  `json:"last_name,omitempty"`
	Email      string  `json:"email,omitempty"`
	Password   string  `json:"password,omitempty"`
	Phone      string  `json:"phone,omitempty"`
	Salary     float32 `json:"salary,omitempty"`
	Department string  `json:"department,omitempty"`
}
