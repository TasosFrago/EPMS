package models

type Invoice struct {
	ID          int64    `json:"invoice_id,omitempty"`
	Total       *float32 `json:"total,omitempty"`
	CurrentCost float32  `json:"current_cost,omitempty"`
	Receiver    int      `json:"receiver,omitempty"`
	Meter       int64    `json:"meter,omitempty"`
	Provider    string   `json:"provider,omitempty"`
	Plan        int64    `json:"plan,omitempty"`
	PlanName    string   `json:"plan_name,omitempty"`
	Month       string   `json:"month,omitempty"`
	Year        int64    `json:"year,omitempty"`
}
