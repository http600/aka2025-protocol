package protocol

// Currency Define a Currency type as a struct
type Currency struct {
    ISOCode     string
    Symbol      string
    NumericCode int
    UnitFactor  int // Number of units in one unit (e.g., 100 for USD, 1 for JPY)
}

// Use constants to define specific currencies
var (
    USD = Currency{ISOCode: "USD", Symbol: "$", NumericCode: 840, UnitFactor: 100} // 1 USD = 100 cents
    EUR = Currency{ISOCode: "EUR", Symbol: "€", NumericCode: 978, UnitFactor: 100} // 1 EUR = 100 cents
    GBP = Currency{ISOCode: "GBP", Symbol: "£", NumericCode: 826, UnitFactor: 100} // 1 GBP = 100 pence
    JPY = Currency{ISOCode: "JPY", Symbol: "¥", NumericCode: 392, UnitFactor: 1}   // 1 JPY = 1 yen (no subunit)
    CNY = Currency{ISOCode: "CNY", Symbol: "¥", NumericCode: 156, UnitFactor: 100} // 1 CNY = 100 fen
)
