package apiHelper

import (
	"errors"
)

type InvoicePayment struct {
	ID          int      `json:"invoice_id,omitempty"`
	Total       *float32 `json:"total,omitempty"`
	CurrentCost float32  `json:"current_cost,omitempty"`
	Receiver    int      `json:"receiver,omitempty"`
	Meter       int      `json:"meter,omitempty"`
	Provider    string   `json:"provider,omitempty"`
	Plan        int      `json:"plan,omitempty"`

	Type     string  `json:"type,omitempty"`
	Price    float32 `json:"price,omitempty"`
	Name     *string `json:"name,omitempty"`
	Month    string  `json:"month,omitempty"`
	Year     int     `json:"year,omitempty"`
	Duration int     `json:"duration,omitempty"`
}

type InvoiceStatus struct {
	ID          int     `json:"invoice_id,omitempty"`
	Provider    string  `json:"provider,omitempty"`
	CurrentCost float32 `json:"current_cost,omitempty"`
	IssueDate   string  `json:"issue_date,omitempty"`
	ExpiryDate  string  `json:"expiry_date,omitempty"`
	IsPaid      bool    `json:"is_paid,omitempty"`
}

var (
	ErrUnauthorized    = errors.New("unauthorized access")
	ErrNotAbleToChoose = errors.New("already committed to plan")
	ErrEmptyQuery      = errors.New("empty set")
	ErrBadReq          = errors.New("value can't be empty")
)
