package models

type Employee struct {
    BadgeID int64       `json:"badge"`
    FirstName string    `json:"first_name"`
    LastName string     `json:"last_name"`
    Email string        `json:"email"`
    Password string     `json:"password"`
    Phone string        `json:"phone"`
    Salary float32      `json:"salary"`
}