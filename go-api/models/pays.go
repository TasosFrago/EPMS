package models

type Pays struct {
    ID int64            `json:"payment_id"`
    User int64          `json:"user"`
    Provider string     `json:"provider"`
    SupplyID int64      `json:"supply_id"`
    Amount *float32     `json:"amount"`
}
