package models

type Invoice struct {
    ID int64                `json:"invoice_id"`
    Total *float32          `json:"total"`
    CurrentCost float32     `json:"current_cost"`
    Receiver int64          `json:"receiver"`
    Meter int64             `json:"meter"`
    Provider string         `json:"provider"`
    Plan int64              `json:"plan"`
}
