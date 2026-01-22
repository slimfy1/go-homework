package main

import "fmt"

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
