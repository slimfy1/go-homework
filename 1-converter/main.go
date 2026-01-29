package main

import (
	"fmt"
	"strings"
)

type CurrencyMap = map[string]float64

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
	// Map с парами валют и их курсами
	conversionRates := CurrencyMap{
		"EUR_RUB": 89.47, // 1 EUR = 89.47 RUB
		"EUR_USD": 1.18,  // 1 EUR = 1.18 USD
		"RUB_EUR": 0.011, // 1 RUB = 0.011 EUR
		"RUB_USD": 0.013, // 1 RUB = 0.013 USD
		"USD_EUR": 0.85,  // 1 USD = 0.85 EUR
		"USD_RUB": 76.04, // 1 USD = 76.04 RUB
	}

	// Формируем ключ из пары валют
	key := firstCurrency + "_" + secondCurrency

	// Получаем курс из map и умножаем на сумму
	if rate, exists := conversionRates[key]; exists {
		return value * rate
	}

	return 0.0 // если пара не найдена
}
func main() {
	// Main body
	// Current price

	value, firstCurrency, secondCurrency := readInputData()
	convertValue := currencyConvert(value, firstCurrency, secondCurrency)

	fmt.Printf("Your money: %.2f", convertValue)
}
