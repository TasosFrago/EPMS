package models

type Pays struct {
	ID       int64    `json:"payment_id,omitempty"`
	User     int      `json:"user,omitempty"`
	Provider string   `json:"provider,omitempty"`
	SupplyID int64    `json:"supply_id,omitempty"`
	Amount   *float32 `json:"amount,omitempty"`
}
