package protocol

type PriceEntity struct {
    Price    int      `json:"price"`
    Currency Currency `json:"currency"`
}

type PriceOfRenting struct {
    PriceEntity
    Period string `json:"duration"`
}
