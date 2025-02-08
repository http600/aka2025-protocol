package protocol

type PriceEntity struct {
    Price    int      `json:"price,omitempty"`
    Currency Currency `json:"currency,omitempty"`
}

type PriceRent struct {
    PriceEntity
    PriceEvery int `json:"price_every,omitempty"`
    MinTerm    int `json:"min_term,omitempty"`
}
