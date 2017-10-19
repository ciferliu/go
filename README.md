# gocurrency

gocurrency is a currency toolkit for Go

---------------------------------------
  * [Features](#features)
  * [Installation](#installation)
  * [Usage](#usage)

---------------------------------------

## Features
  * [ISO 4217](https://www.currency-iso.org/dam/downloads/lists/list_one.xml "ISO 4217") standard currencies
  * user-defined currencies
  * banker rounding algorithm
  * operations: Add、Minus、Multiply、Divide、Fx、Equals、IsGreatThan

---------------------------------------

## Installation
Simple install the package to your [$GOPATH](https://github.com/golang/go/wiki/GOPATH "GOPATH") with the [go tool](https://golang.org/cmd/go/ "go command") from shell:
```bash
$ go get github.com/ciferliu/gocurrency
```
Make sure [Git is installed](https://git-scm.com/downloads) on your machine and in your system's `PATH`.

## Usage
```go
import "github.com/ciferliu/gocurrency"
import "fmt"

//1st step: new currency object and register to factory
gocurrency.Factory.NewCurrency("usd", 2)
gocurrency.Factory.NewCurrency("CNY", 2)

// or you can use the factory's InitFromOnlineIso4217Xml function to support all of ISO 4217 currencies:

//gocurrency.Factory.InitFromOnlineIso4217Xml()

//2nd step: new amount object with currency property
usdAmount1, _ := gocurrency.Factory.NewAmountInBasicUnit("USD", "1.567") //round to $1.57
usdAmount2, _ := gocurrency.Factory.NewAmountInMinorUnit("USD", 43)//43 cent = $0.43
usdAmount, _ := usdAmount1.Add(usdAmount2)// $1.57 + $0.43 = $2.00
fmt.Println(usdAmount.String()) //USD 2.00
```
