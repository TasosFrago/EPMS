package models

type Invoice struct {
	ID          int      `json:"invoice_id,omitempty"`
	Total       *float32 `json:"total,omitempty"`
	CurrentCost float32  `json:"current_cost,omitempty"`
	Receiver    int      `json:"receiver,omitempty"`
	Meter       int      `json:"meter,omitempty"`
	Provider    string   `json:"provider,omitempty"`
	Plan        int      `json:"plan,omitempty"`
}
