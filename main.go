package main

import "fmt"

func readInputData() (float64, string, string) {
	var (
		value          float64
		firstCurrency  string
		secondCurrency string
	)
	fmt.Scanln(&value, &firstCurrency, &secondCurrency)
	return value, firstCurrency, secondCurrency
}

func currencyConvert(value int, firstCurrency string, secondCurrency string) float64 {

}
func main() {
	// Main body
	// Current price
	const usdToEur float64 = 0.85
	const usdToRub float64 = 76.0382

	fmt.Println("Enter the amount EUR you want to convert to RUB:")
	var amount float64
	fmt.Scanln(&amount)

	const eurToRub = usdToRub / usdToEur
	fmt.Printf("Your money: %v", amount*eurToRub)
}
