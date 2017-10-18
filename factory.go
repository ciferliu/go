package currency

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

//Factory is currency factory
var Factory = factory{currencyMap: make(map[string]Currency), mapLocker: new(sync.Mutex), initLocker: new(sync.Mutex)}
var currencyCodeReg, _ = regexp.Compile("^[A-Z]{3}$")

type factory struct {
	currencyMap map[string]Currency // key is ISO 4217 three-letter alphabetic code
	mapLocker   *sync.Mutex

	initFlag   bool
	initLocker *sync.Mutex
}

//NewCurrency create a new currency object
//return error if the code is not a three-letter alphabetic code
func (factory *factory) NewCurrency(currencyCode string, minorUnitDigits uint8) (Currency, error) {
	currencyCode = strings.ToUpper(strings.TrimSpace(currencyCode))
	currency, exists := factory.currencyMap[currencyCode]
	if exists {
		return currency, nil
	}

	if !currencyCodeReg.MatchString(currencyCode) {
		return Currency{}, errors.New("the currency code is not three-letter alphabetic code")
	}

	factory.mapLocker.Lock()
	defer factory.mapLocker.Unlock()
	currency, exists = factory.currencyMap[currencyCode] //double check
	if !exists {
		currency = Currency{currencyCode, minorUnitDigits}
		factory.currencyMap[currencyCode] = currency
	}
	return currency, nil
}

//InitFromOnlineIso4217Xml init currencies from online ISON 4217 XML
//URL: https://www.currency-iso.org/dam/downloads/lists/list_one.xml
func (factory *factory) InitFromOnlineIso4217Xml() error {
	if factory.initFlag {
		return nil
	}

	factory.initLocker.Lock()
	defer factory.initLocker.Unlock()

	if factory.initFlag { //double check
		return nil
	}

	resp, err := http.Get("https://www.currency-iso.org/dam/downloads/lists/list_one.xml")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	iso4217XmlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var iso4217Xml iso4217Xml
	err = xml.Unmarshal(iso4217XmlBytes, &iso4217Xml)
	if err != nil {
		return err
	}

	for _, ccyNtry := range iso4217Xml.CcyTbl.CcyNtrys {
		var minorUnitDigits int
		if ccyNtry.CcyMnrUnts == "N.A." {
			minorUnitDigits = 0
		} else {
			minorUnitDigits, _ = strconv.Atoi(ccyNtry.CcyMnrUnts)
		}
		factory.NewCurrency(ccyNtry.Ccy, uint8(minorUnitDigits))
	}
	factory.initFlag = true
	return nil
}

//NewAmountByBasicUnit create a new amount object by using basic unit value
//return error if currencyCode is not a three-letter alphabetic code
//return error if currencyCode is not managed by factory
//return error if basicUnitValue is not a numberic value
func (factory *factory) NewAmountByBasicUnit(currencyCode string, basicUnitValue string) (Amount, error) {
	currencyCode = strings.ToUpper(strings.TrimSpace(currencyCode))
	if !currencyCodeReg.MatchString(currencyCode) {
		return Amount{}, errors.New("the currencyCode is not a three-letter alphabetic code")
	}
	currency, exists := factory.currencyMap[currencyCode]
	if !exists {
		return Amount{}, errors.New("currency code is not found")
	}

	basicUnitValue = strings.TrimSpace(basicUnitValue)
	floatValue, err := strconv.ParseFloat(basicUnitValue, 64)
	if err != nil {
		return Amount{}, errors.New("basicUnitValue is not a numberic value")
	}

	amount := newZeroAmount(currency)
	amount.setBasicUnitValue(floatValue)
	return amount, nil
}

//NewAmountByMinorUnit create a new amount object by using minor unit value
//return error if currencyCode is not a three-letter alphabetic code
//return error if currencyCode is not managed by factory
func (factory *factory) NewAmountByMinorUnit(currencyCode string, minorUnitValue int64) (Amount, error) {
	currencyCode = strings.ToUpper(strings.TrimSpace(currencyCode))
	if !currencyCodeReg.MatchString(currencyCode) {
		return Amount{}, errors.New("the currencyCode is not three-letter alphabetic code")
	}
	currency, exists := factory.currencyMap[currencyCode]
	if !exists {
		return Amount{}, errors.New("currency code is not found")
	}

	amount := newZeroAmount(currency)
	amount.setMinorUnitValue(minorUnitValue)
	return amount, nil
}

//GetCurrencyByCode return a Currency object by using  a three-letter alphabetic code
//return error if currencyCode is not a three-letter alphabetic code
//return error if currencyCode is not managed by factory
func (factory *factory) GetCurrencyByCode(currencyCode string) (Currency, error) {
	currencyCode = strings.ToUpper(strings.TrimSpace(currencyCode))
	if !currencyCodeReg.MatchString(currencyCode) {
		return Currency{}, errors.New("the currencyCode is not three-letter alphabetic code")
	}

	currency, exists := factory.currencyMap[currencyCode]
	if !exists {
		return Currency{}, errors.New("currency code is not found")
	}
	return currency, nil
}
