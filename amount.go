package currency

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//Amount is an amount object of currency
type Amount struct {
	curreny        Currency
	basicUnitValue string //the value of amount in currency's basic unit
	minorUnitValue int64  //the value of amount in currency's minor unit
}

//newZeroAmount create a new Amount object with Currency property, but zero value
func newZeroAmount(curreny Currency) Amount {
	return Amount{curreny, "0", 0}
}

//setBasicUnitValue set the value of amount in currency's basic unit (e.g: USD，1.5 dollar or 1.50 dollar)
func (amount *Amount) setBasicUnitValue(floatValue float64) {
	amount.roundBasicUnitValue(floatValue)
}

//setMinorUnitValue set the value of amount in currency's minor unit(e.g: USD, 150 cent)
func (amount *Amount) setMinorUnitValue(value int64) {
	basicUnitFloatValue := float64(value) / math.Pow10(int(amount.curreny.MinorUnitDigits()))
	amount.setBasicUnitValue(basicUnitFloatValue)
}

//BasicUnitValue returns the value of amount in currency's basic unit
func (amount Amount) BasicUnitValue() string {
	return amount.basicUnitValue
}

//MinorUnitValue returns the value of amount in currency's minor unit
func (amount Amount) MinorUnitValue() int64 {
	return amount.minorUnitValue
}

//CurrencyCode returns the currency code (three-letter alphabetic code) of amount
func (amount Amount) CurrencyCode() string {
	return amount.curreny.Code()
}

//Add return (amount + other)
//return error if the currency of two amount are not same
func (amount Amount) Add(other Amount) (Amount, error) {
	if amount.curreny.Code() != other.curreny.Code() {
		return Amount{}, errors.New("Amount add fail: curreny are not same")
	}

	totalValue := amount.minorUnitValue + other.minorUnitValue
	result := newZeroAmount(amount.curreny)
	result.setMinorUnitValue(totalValue)
	return result, nil
}

//Minus return (amount - other)
//return error if the currency of two amount are not same
func (amount Amount) Minus(other Amount) (Amount, error) {
	if amount.curreny.Code() != other.curreny.Code() {
		return Amount{}, errors.New("Amount minus fail: curreny are not same")
	}

	totalValue := amount.minorUnitValue - other.minorUnitValue
	result := newZeroAmount(amount.curreny)
	result.setMinorUnitValue(totalValue)
	return result, nil
}

//Multiply return (amount * factor)
func (amount Amount) Multiply(factor float64) Amount {
	basicUnitFloatValue, _ := strconv.ParseFloat(amount.basicUnitValue, 64)
	basicUnitFloatValue = basicUnitFloatValue * factor
	result := newZeroAmount(amount.curreny)
	result.roundBasicUnitValue(basicUnitFloatValue)
	return result
}

//Divide return (amount / factor)
//return error if factor is 0
func (amount Amount) Divide(factor float64) (Amount, error) {
	if factor == 0 {
		return Amount{}, errors.New("Amount divide fail: factor can not be 0")
	}

	basicUnitFloatValue, _ := strconv.ParseFloat(amount.basicUnitValue, 64)
	basicUnitFloatValue = basicUnitFloatValue / factor
	result := newZeroAmount(amount.curreny)
	result.roundBasicUnitValue(basicUnitFloatValue)
	return result, nil
}

//Fx foreign exchange
//return error if targetCurrencyCode is not three-letter alphabetic code
//return error if targetCurrencyCode is not managed by factory
//return error if rate=0
func (amount Amount) Fx(targetCurrencyCode string, rate float64) (Amount, error) {
	targetCurrencyCode = strings.ToUpper(strings.TrimSpace(targetCurrencyCode))
	if targetCurrencyCode == amount.curreny.Code() {
		return amount, nil
	}
	if rate == 0 {
		return Amount{}, errors.New("fx rate can't be 0")
	}
	targetCurrency, err := Factory.GetCurrencyByCode(targetCurrencyCode)
	if err != nil {
		return Amount{}, err
	}

	minorUnitDigits := int(amount.curreny.MinorUnitDigits())
	basicUnitFloatValue := float64(amount.minorUnitValue) / math.Pow10(minorUnitDigits)
	floatValue := basicUnitFloatValue * rate
	result := newZeroAmount(targetCurrency)
	result.roundBasicUnitValue(floatValue)
	return result, nil
}

//IsEquals return true if the currency and value are same, otherwise return false
func (amount Amount) IsEquals(other Amount) bool {
	return amount.curreny.Code() == other.curreny.Code() && amount.minorUnitValue == other.minorUnitValue
}

//IsGreatThan return true if amount > other, otherwise return false,
//return error if the currency of two amount are not same
func (amount Amount) IsGreatThan(other Amount) (bool, error) {
	if amount.curreny.Code() != other.curreny.Code() {
		return false, errors.New("curreny are not same")
	}
	result := amount.minorUnitValue > other.minorUnitValue
	return result, nil
}

//String returns default format string of amount(e.g.: USD 1.00)
func (amount Amount) String() string {
	return fmt.Sprintf("%s %s", amount.curreny.Code(), amount.basicUnitValue)
}

//roundBasicUnitValue using banker's rounding algorithm
func (amount *Amount) roundBasicUnitValue(floatValue float64) {
	formatStr := fmt.Sprintf("%%0.%df", amount.curreny.MinorUnitDigits())
	stringValue := fmt.Sprintf(formatStr, floatValue)
	amount.basicUnitValue = stringValue
	stringValue = strings.Replace(stringValue, ".", "", 1)
	intValue, _ := strconv.Atoi(stringValue)
	amount.minorUnitValue = int64(intValue)
}
