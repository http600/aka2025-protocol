package protocol

type Item struct {
    Key          string       `json:"key,omitempty"`
    Title        string       `json:"title,omitempty"`
    CoverImg     string       `json:"cover_img"`
    PriceForSale PriceEntity  `json:"price_for_sale,omitempty"`
    PriceToRent  []PriceOffer `json:"price_rent,omitempty"`
}
