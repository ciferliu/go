package currency

import "testing"
import "strconv"

func init() {
	Factory.NewCurrency("USD", 2)
	Factory.NewCurrency("CNY", 2)
}

func TestAdd(t *testing.T) {
	//case 1:
	usdAmount1, _ := Factory.NewAmountInBasicUnit("usd", "1.567")
	usdAmount2, _ := Factory.NewAmountInBasicUnit("usd", "0.43")
	got, _ := usdAmount1.Add(usdAmount2)
	want, _ := Factory.NewAmountInBasicUnit("usd", "2")
	if !want.IsEquals(got) {
		t.Errorf("%s Add(%s) == %s, want %s", usdAmount1.String(), usdAmount2.String(), got.String(), want.String())
	}

	//case 2:
	cnyAmount, _ := Factory.NewAmountInBasicUnit("cny", "0.43")
	got, err := usdAmount1.Add(cnyAmount)
	if err == nil {
		t.Errorf("%s Add(%s) should be return an error, but no error return", usdAmount1.String(), cnyAmount.String())
	}
}

func TestMinus(t *testing.T) {
	//case 1:
	usdAmount1, _ := Factory.NewAmountInBasicUnit("usd", "2")
	usdAmount2, _ := Factory.NewAmountInBasicUnit("usd", "1.567")
	got, _ := usdAmount1.Minus(usdAmount2)
	want, _ := Factory.NewAmountInBasicUnit("usd", "0.43")
	if !want.IsEquals(got) {
		t.Errorf("%s Minus(%s) == %s, want %s", usdAmount1.String(), usdAmount2.String(), got.String(), want.String())
	}

	//case 2:
	cnyAmount, _ := Factory.NewAmountInBasicUnit("cny", "1.567")
	got, err := usdAmount1.Minus(cnyAmount)
	if err == nil {
		t.Errorf("%s Minus(%s) should be return an error, but no error return", usdAmount1.String(), cnyAmount.String())
	}
}

func TestMultiply(t *testing.T) {
	//case 1:
	usdAmount1, _ := Factory.NewAmountInBasicUnit("usd", "2")
	factorStr := "0"
	factor, _ := strconv.ParseFloat(factorStr, 10)
	got := usdAmount1.Multiply(factor)
	want, _ := Factory.NewAmountInBasicUnit("usd", "0")
	if !want.IsEquals(got) {
		t.Errorf("%s Multiply(%s) == %s, want %s", usdAmount1.String(), factorStr, got.String(), want.String())
	}

	//case 2:
	factorStr = "1.0"
	factor, _ = strconv.ParseFloat(factorStr, 10)
	got = usdAmount1.Multiply(factor)
	want, _ = Factory.NewAmountInBasicUnit("usd", "2")
	if !want.IsEquals(got) {
		t.Errorf("%s Multiply(%s) == %s, want %s", usdAmount1.String(), factorStr, got.String(), want.String())
	}

	//case 3:
	factorStr = "0.436"
	factor, _ = strconv.ParseFloat(factorStr, 10)
	got = usdAmount1.Multiply(factor)
	want, _ = Factory.NewAmountInBasicUnit("usd", "0.87")
	if !want.IsEquals(got) {
		t.Errorf("%s Multiply(%s) == %s, want %s", usdAmount1.String(), factorStr, got.String(), want.String())
	}
}

func TestDivide(t *testing.T) {
	//case 1:
	usdAmount1, _ := Factory.NewAmountInBasicUnit("usd", "2")
	factorStr := "0"
	factor, _ := strconv.ParseFloat(factorStr, 10)
	got, err := usdAmount1.Divide(factor)
	if err == nil {
		t.Errorf("%s Divide(0) should be return an error, but no error return", usdAmount1.String())
	}

	//case 2:
	factorStr = "1.0"
	factor, _ = strconv.ParseFloat(factorStr, 10)
	got, _ = usdAmount1.Divide(factor)
	want, _ := Factory.NewAmountInBasicUnit("usd", "2")
	if !want.IsEquals(got) {
		t.Errorf("%s Divide(%s) == %s, want %s", usdAmount1.String(), factorStr, got.String(), want.String())
	}

	//case 3:
	factorStr = "0.436"
	factor, _ = strconv.ParseFloat(factorStr, 10)
	got, _ = usdAmount1.Divide(factor)
	want, _ = Factory.NewAmountInBasicUnit("usd", "4.59")
	if !want.IsEquals(got) {
		t.Errorf("%s Divide(%s) == %s, want %s", usdAmount1.String(), factorStr, got.String(), want.String())
	}
}

func TestFx(t *testing.T) {
	//case 1:
	usdAmount1, _ := Factory.NewAmountInBasicUnit("usd", "2")
	rateStr := "0"
	rate, _ := strconv.ParseFloat(rateStr, 10)
	got, err := usdAmount1.Fx("USD", rate)
	want, _ := Factory.NewAmountInBasicUnit("usd", "2")
	if !want.IsEquals(got) {
		t.Errorf("%s Fx(\"USD\", 0) == %s, want %s", usdAmount1.String(), got.String(), want.String())
	}

	//case 2:
	got, err = usdAmount1.Fx("CNY", rate)
	if err == nil {
		t.Errorf("%s Fx(\"CNY\", 0) should be return an error, but no error return", usdAmount1.String())
	}

	//case 3:
	rateStr = "6.789"
	rate, _ = strconv.ParseFloat(rateStr, 10)
	got, err = usdAmount1.Fx("CN", rate)
	if err == nil {
		t.Errorf("%s Fx(\"CN\", %s) should be return an error, but no error return", usdAmount1.String(), rateStr)
	}

	//case 4:
	got, err = usdAmount1.Fx("CNY", rate)
	want, _ = Factory.NewAmountInBasicUnit("CNY", "13.58")
	if !want.IsEquals(got) {
		t.Errorf("%s Fx(\"CNY\", %s) == %s, want %s", usdAmount1.String(), rateStr, got.String(), want.String())
	}
}

func TestIsEquals(t *testing.T) {
	//case 1:
	usdAmount1, _ := Factory.NewAmountInBasicUnit("usd", "2")
	usdAmount2, _ := Factory.NewAmountInBasicUnit("usd", "2.0")
	equal := usdAmount1.IsEquals(usdAmount2)
	if !equal {
		t.Errorf("%s IsEquals(%s) == %t, want true", usdAmount1.String(), usdAmount2.String(), equal)
	}

	//case 1:
	usdAmount3, _ := Factory.NewAmountInBasicUnit("usd", "3.0")
	equal = usdAmount1.IsEquals(usdAmount3)
	if equal {
		t.Errorf("%s IsEquals(%s) == %t, want false", usdAmount1.String(), usdAmount3.String(), equal)
	}

	//case 1:
	cnyAmount, _ := Factory.NewAmountInBasicUnit("cny", "2")
	equal = usdAmount1.IsEquals(cnyAmount)
	if equal {
		t.Errorf("%s IsEquals(%s) == %t, want false", usdAmount1.String(), cnyAmount.String(), equal)
	}
}

func TestIsGreatThan(t *testing.T) {
	//case 1:
	usdAmount, _ := Factory.NewAmountInBasicUnit("usd", "2")
	cnyAmount, _ := Factory.NewAmountInBasicUnit("cny", "2.0")
	isGreatThan, err := usdAmount.IsGreatThan(cnyAmount)
	if err == nil {
		t.Errorf("%s IsGreatThan(%s), should be return an error, but no error return", usdAmount.String(), cnyAmount.String())
	}

	//case 2:
	usdAmount2, _ := Factory.NewAmountInBasicUnit("usd", "2.0")
	isGreatThan, err = usdAmount.IsGreatThan(cnyAmount)
	if isGreatThan {
		t.Errorf("%s IsGreatThan(%s) == %t, want false", usdAmount.String(), usdAmount2.String(), isGreatThan)
	}

	//case 3:
	usdAmount3, _ := Factory.NewAmountInBasicUnit("usd", "3.0")
	isGreatThan, err = usdAmount3.IsGreatThan(usdAmount)
	if !isGreatThan {
		t.Errorf("%s IsGreatThan(%s) == %t, want true", usdAmount3.String(), usdAmount.String(), isGreatThan)
	}
}

func TestString(t *testing.T) {
	usdAmount, _ := Factory.NewAmountInBasicUnit("usd", "2")
	got := usdAmount.String()
	want := "USD 2.00"
	if want != got {
		t.Errorf("%s String() == %s, want %s", usdAmount.String(), got, want)
	}
}
