package in

type Point struct {
    Holder   string `json:"holder,omitempty"`
    Item     string `json:"item,omitempty"`
    Offer    uint8  `json:"offer,omitempty"`
    PriceKey int    `json:"price_key,omitempty"`
    Term     int    `json:"term,omitempty"`
}
