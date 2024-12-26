package models

type Plan struct {
	ID       int64   `json:"plan_id,omitempty"`
	Type     string  `json:"type,omitempty"`
	Price    float32 `json:"price,omitempty"`
	Name     *string `json:"name,omitempty"`
	Provider string  `json:"provider,omitempty"`
	Month    string  `json:"month,omitempty"`
	Year     int     `json:"year,omitempty"`
	Duration int     `json:"duration,omitempty"`
}
