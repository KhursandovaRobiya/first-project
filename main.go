package main

import (
	"fmt"
)

func main() {
	var sum float64
	fmt.Println("Введите сумму в сомони: ")
	fmt.Scan(&sum)
	const tjsToUsd = 0.0912
	const tjsToRub = 8.64
	var usd, rub float64
	usd = sum * tjsToUsd
	rub = sum * tjsToRub
	fmt.Printf("%.2f сомони = %.2f доллар\n", sum, usd)
	fmt.Printf("%.2f сомони = %.2f рубль", sum, rub)
}
