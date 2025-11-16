package main

import "fmt"

func main() {
	// Создать новый проект
	// Объявить пакет main
	// Объявить функцию main
	// Создать константы конвертации: Из USD в EUR, из USD в RUB, рассчитать EUR в RUB на основании первых двух
	const (
		usdToEur = 0.86  // ~ курс USD в EUR
		usdTuRub = 80.88 // ~ курс USD в RUB
	)
	eurTuRub := usdTuRub / usdToEur // рассчитываем EUR в RUB на основании первых двух
	fmt.Printf("Курс EUR в RUB: %.2f\n", eurTuRub)
}
