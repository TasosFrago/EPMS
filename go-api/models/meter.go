package models

type Meter struct {
	ID         int    `json:"supply_id,omitempty"`
	Plan       *int   `json:"plan,omitempty"`
	Status     *bool  `json:"status,omitempty"`
	KWH        *int   `json:"kWh,omitempty"`
	Address    string `json:"address,omitempty"`
	RatedPower string `json:"rated_power,omitempty"`
	Owner      int    `json:"owner,omitempty"`
	Agent      *int64 `json:"agent,omitempty"`
	Department string `json:"department,omitempty"`
}
