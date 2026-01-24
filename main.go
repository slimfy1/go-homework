package main

import "fmt"

func readInputData() (float64, string, string) {
	var (
		value          float64
		firstCurrency  string
		secondCurrency string
	)
	// First step - enter source currency
	for {
		fmt.Println("Enter your currency: EUR, RUB, USD")
		fmt.Scanln(&firstCurrency)
		if firstCurrency == "EUR" || firstCurrency == "RUB" || firstCurrency == "USD" {
			break
		}
		fmt.Println("Invalid currency! Please enter EUR, RUB, or USD")
	}

	// Second step - enter amount
	fmt.Println("Enter amount of currency")
	fmt.Scanln(&value)

	// Third step - enter out currency
	for {
		switch firstCurrency {
		case "EUR":
			fmt.Println("Enter your currency: RUB, USD")
		case "RUB":
			fmt.Println("Enter your currency: EUR, USD")
		case "USD":
			fmt.Println("Enter your currency: EUR, RUB")
		}
		fmt.Scanln(&secondCurrency)

		// Validate second currency
		if secondCurrency != firstCurrency && (secondCurrency == "EUR" || secondCurrency == "RUB" || secondCurrency == "USD") {
			break
		}
		fmt.Println("Invalid currency or same as source currency!")
	}

	//fmt.Scanln(&value, &firstCurrency, &secondCurrency)
	return value, firstCurrency, secondCurrency
}

func currencyConvert(value float64, firstCurrency string, secondCurrency string) float64 {
	const usdToEur float64 = 0.85
	const usdToRub float64 = 76.0382
	const eurToRub = usdToRub / usdToEur
	const rubToEur = 1.0 / eurToRub
	const eurToUsd = 1.0 / usdToEur
	const rubToUsd = 1.0 / usdToRub

	var convert = 0.0

	switch firstCurrency {
	case "EUR":
		if secondCurrency == "RUB" {
			convert := value * eurToRub
			return convert
		} else if secondCurrency == "USD" {
			convert = value * eurToUsd

		}
	case "RUB":
		if secondCurrency == "EUR" {
			convert := value * rubToEur
			return convert
		} else if secondCurrency == "USD" {
			convert = value * rubToUsd
		}
	case "USD":
		if secondCurrency == "EUR" {
			convert = value * usdToEur
		} else if secondCurrency == "RUB" {
			convert = value * usdToRub
		}
	}
	return convert
}
func main() {
	// Main body
	// Current price

	value, firstCurrency, secondCurrency := readInputData()
	convertValue := currencyConvert(value, firstCurrency, secondCurrency)

	fmt.Printf("Your money: %.2f", convertValue)
}
