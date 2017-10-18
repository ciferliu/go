package currency

//Currency is an ISO 4217 currency, maintains three-letter alphabetic code and fraction digits of minor currency unit
type Currency struct {
	code            string //ISO 4217 three-letter alphabetic code
	minorUnitDigits uint8  //the fraction digits of minor currency unit
}

//Code returns the ISO 4217 three-letter alphabetic code
func (currency Currency) Code() string {
	return currency.code
}

//MinorUnitDigits returns the fraction digits of minor currency unit
func (currency Currency) MinorUnitDigits() uint8 {
	return currency.minorUnitDigits
}
