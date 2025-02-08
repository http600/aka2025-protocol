package protocol

type Item struct {
    Key       string      `json:"key,omitempty"`
    Title     string      `json:"title,omitempty"`
    CoverImg  string      `json:"cover_img"`
    PriceSale PriceEntity `json:"price_sale,omitempty"`
    PriceRent []PriceRent `json:"price_rent,omitempty"`
}
