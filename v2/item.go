package protocol

type Item struct {
    Key         string         `json:"key"`
    Title       string         `json:"title"`
    CoverImg    string         `json:"cover_img"`
    PriceToBuy  PriceEntity    `json:"price_to_buy"`
    PriceToRent PriceOfRenting `json:"price_to_rent"`
}
