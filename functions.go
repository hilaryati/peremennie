package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Функция для чтения строки из stdin
func readInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Пустая функция расчёта
func calculate(amount float64, from string, to string) float64 {
	// Здесь позже будет логика конвертации
	return 0
}

func main() {
	amountStr := readInput("Введите сумму: ")
	from := readInput("Введите исходную валюту: ")
	to := readInput("Введите целевую валюту: ")

	var amount float64
	fmt.Sscanf(amountStr, "%f", &amount)

	result := calculate(amount, from, to)
	fmt.Println("Результат:", result)
}
