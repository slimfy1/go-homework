package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
1. Принимает операцию (AVG - среднее, SUM - сумму, MED - медиану)
2. Принимает неограниченное число чисел через запятую (2, 10, 9)
3. Разбивает строку чисел по запятым и затем делает расчёт в зависимости от операции выводя результат
*/
func main() {
	operation, numbers := readInputData()
	result := calculate(operation, numbers)
	fmt.Println(result)
}

func stringToInt(s string) ([]int, error) {
	var convertNumbers = []int{}
	// Read numbers
	fmt.Println("Enter the numbers:")

	// Making slice of numbers
	numbersSlice := strings.SplitSeq(s, ",")

	for i := range numbersSlice {
		j, err := strconv.Atoi(i)
		if err != nil {
			fmt.Println("Invalid numbers")
			return nil, err
		}
		convertNumbers = append(convertNumbers, j)
	}

	return convertNumbers, nil
}

func readInputData() (string, []int) {

	var operation string
	var numbers string

	// Read operation
	for {
		fmt.Println("Enter the operation (AVG, SUM, MED):")
		fmt.Scanln(&operation)
		if operation == "AVG" || operation == "SUM" || operation == "MED" {
			break
		}
		fmt.Println("Invalid operation")
	}
numbersLoop:
	for {
		fmt.Println("Enter the numbers:")
		fmt.Scanln(&numbers)

		convertNumbers, err := stringToInt(numbers)
		if err != nil {
			continue numbersLoop
		}
		return operation, convertNumbers
	}
}

func Median(values []float64) float64 {
	n := len(values)
	if n == 0 {
		return 0 // или panic / error
	}

	// копия, чтобы не портить исходный слайс
	sorted := append([]float64(nil), values...)
	sort.Float64s(sorted)

	mid := n / 2
	if n%2 == 1 {
		return sorted[mid]
	}
	return (sorted[mid-1] + sorted[mid]) / 2.0
}

func calculate(operation string, numbers []int) float64 {
	switch operation {
	case "AVG":
		var sum float64
		for _, number := range numbers {
			sum += float64(number)
		}
		if len(numbers) == 0 {
			return 0
		}
		return float64(sum / float64(len(numbers)))
	case "SUM":
		sum := 0.0
		for _, number := range numbers {
			sum += float64(number)
		}
		return sum
	case "MED":
		var floatNumbers []float64
		for _, number := range numbers {
			floatNumbers = append(floatNumbers, float64(number))
		}
		return Median(floatNumbers)
	}
	return 0
}
