package protocol

import "time"

type PriceEntity struct {
    Price    int      `json:"price"`
    Currency Currency `json:"currency"`
}

type PriceOfRenting struct {
    PriceEntity
    Period time.Duration `json:"duration"`
}
