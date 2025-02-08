package in

type Point struct {
    Item     string `json:"item,omitempty"`
    PriceKey string `json:"price_key,omitempty"`
    Term     int    `json:"term,omitempty"`
}
