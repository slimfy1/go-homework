package main

import (
	"fmt"
	"strings"
)

func validateCurrency(prompt string, excludeCurrency string) string {
	validCurrencies := []string{"EUR", "RUB", "USD"}

	for {
		fmt.Println(prompt)
		var currency string
		fmt.Scanln(&currency)

		// Проверка на валидность и отличие от исключаемой валюты
		for _, valid := range validCurrencies {
			if currency == valid && currency != excludeCurrency {
				return currency
			}
		}

		if excludeCurrency != "" {
			fmt.Println("Invalid currency or same as source currency!")
		} else {
			fmt.Println("Invalid currency! Please enter EUR, RUB, or USD")
		}
	}
}

func validateAmount() float64 {
	for {
		fmt.Println("Enter amount of currency")
		var value float64
		_, err := fmt.Scanln(&value)

		if err == nil && value > 0 {
			return value
		}

		fmt.Println("Invalid amount! Please enter a positive number")
		// Очистка буфера от невалидного значения
		var discard string
		fmt.Scanln(&discard)
	}
}
func readInputData() (float64, string, string) {
	// Шаг 1: Ввод исходной валюты
	firstCurrency := validateCurrency("Enter your currency: EUR, RUB, USD", "")

	// Шаг 2: Ввод суммы
	value := validateAmount()

	// Шаг 3: Ввод целевой валюты (исключаем исходную)
	prompt := fmt.Sprintf("Enter target currency (available: %s)", getAvailableCurrencies(firstCurrency))
	secondCurrency := validateCurrency(prompt, firstCurrency)

	return value, firstCurrency, secondCurrency
}

// Вспомогательная функция для формирования списка доступных валют
func getAvailableCurrencies(exclude string) string {
	currencies := []string{"EUR", "RUB", "USD"}
	available := []string{}

	for _, curr := range currencies {
		if curr != exclude {
			available = append(available, curr)
		}
	}

	return strings.Join(available, ", ")
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
