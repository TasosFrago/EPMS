package models

type Meter struct {
    ID int              `json:"supply_id"`
    Plan *int           `json:"plan"`
    Status *bool        `json:"status"`
    KWH *int            `json:"kWh"`
    Address string      `json:"address"`
    RatedPower string   `json:"rated_power"`
    Owner *int64        `json:"owner"`
    Adgent *int64       `json:"adgent"`
}
