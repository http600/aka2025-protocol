package protocol

import "aka2025-protocol.go/aka2025/protocol/v2/currency"

type PriceEntity struct {
    Price    int               `json:"price,omitempty"`
    Currency currency.Currency `json:"currency,omitempty"`
}

type PriceOffer struct {
    PriceEntity
    PriceEvery int `json:"price_every,omitempty"`
    MinTerm    int `json:"min_term,omitempty"`
}
