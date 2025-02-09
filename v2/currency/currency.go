package currency

// Currency represents a currency with its details
type Currency struct {
    ISOCode     string
    Symbol      string
    NumericCode int
    UnitFactor  int // Number of units in one main unit (e.g., 100 for USD, 1 for JPY)
}

// Predefined currencies as constants
const (
    USD = "USD"
    EUR = "EUR"
    JPY = "JPY"
    GBP = "GBP"
    CNY = "CNY"
    AUD = "AUD"
    CAD = "CAD"
    CHF = "CHF"
    INR = "INR"
    BRL = "BRL"
    RUB = "RUB"
    KRW = "KRW"
    SGD = "SGD"
    MXN = "MXN"
    NZD = "NZD"
    SEK = "SEK"
    HKD = "HKD"
    NOK = "NOK"
    TRY = "TRY"
    ZAR = "ZAR"
)

// currencies holds all predefined currency details as a slice
var currencies = []Currency{
    {USD, "$", 840, 100},
    {EUR, "€", 978, 100},
    {JPY, "¥", 392, 1},
    {GBP, "£", 826, 100},
    {CNY, "¥", 156, 100},
    {AUD, "A$", 36, 100},
    {CAD, "C$", 124, 100},
    {CHF, "CHF", 756, 100},
    {INR, "₹", 356, 100},
    {BRL, "R$", 986, 100},
    {RUB, "₽", 643, 100},
    {KRW, "₩", 410, 1},
    {SGD, "S$", 702, 100},
    {MXN, "Mex$", 484, 100},
    {NZD, "NZ$", 554, 100},
    {SEK, "kr", 752, 100},
    {HKD, "HK$", 344, 100},
    {NOK, "kr", 578, 100},
    {TRY, "₺", 949, 100},
    {ZAR, "R", 710, 100},
}

// GetCurrency fetches a Currency by its ISO code
func GetCurrency(code string) (Currency, bool) {
    for _, c := range currencies {
        if c.ISOCode == code {
            return c, true
        }
    }
    return Currency{}, false
}
