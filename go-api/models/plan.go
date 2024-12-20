package models

type Plan struct {
	ID       int64   `json:"plan_id"`
	Type     string  `json:"type"`
	Price    float32 `json:"price"`
	Name     *string `json:"name"`
	Provider string  `json:"provider"`
	Month    string  `json:"month"`
	Year     int     `json:"year"`
	Duration int     `json:"duration"`
}
