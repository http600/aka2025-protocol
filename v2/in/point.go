package in

type Point struct {
    Item     string `json:"item,omitempty"`
    Offer    uint8  `json:"offer,omitempty"`
    PriceKey string `json:"price_key,omitempty"`
    Term     int    `json:"term,omitempty"`
}
