package protocol

// Currency defines a type for handling currency information
type Currency struct {
    ISOCode     string
    Symbol      string
    NumericCode int
    UnitFactor  int // Number of units in one main unit (e.g., 100 for USD, 1 for JPY)
}

// Predefined currencies including top 20 by popularity
var currencies = map[string]Currency{
    "USD": {"USD", "$", 840, 100},
    "EUR": {"EUR", "€", 978, 100},
    "JPY": {"JPY", "¥", 392, 1},
    "GBP": {"GBP", "£", 826, 100},
    "CNY": {"CNY", "¥", 156, 100},
    "AUD": {"AUD", "A$", 36, 100},
    "CAD": {"CAD", "C$", 124, 100},
    "CHF": {"CHF", "CHF", 756, 100},
    "INR": {"INR", "₹", 356, 100},
    "BRL": {"BRL", "R$", 986, 100},
    "RUB": {"RUB", "₽", 643, 100},
    "KRW": {"KRW", "₩", 410, 1},
    "SGD": {"SGD", "S$", 702, 100},
    "MXN": {"MXN", "Mex$", 484, 100},
    "NZD": {"NZD", "NZ$", 554, 100},
    "SEK": {"SEK", "kr", 752, 100},
    "HKD": {"HKD", "HK$", 344, 100},
    "NOK": {"NOK", "kr", 578, 100},
    "TRY": {"TRY", "₺", 949, 100},
    "ZAR": {"ZAR", "R", 710, 100},
}

// GetCurrency fetches a Currency by its ISO code
func GetCurrency(code string) (Currency, bool) {
    currency, exists := currencies[code]
    return currency, exists
}
