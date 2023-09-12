package main

import "fmt"

func main() {
	fmt.Println("Soma de 1 + 1 = ", Soma(1, 1))
}

// Soma retorna a soma de dois inteiros
func Soma(a, b int) int {
	return a + b + 10
}

func Divisao(a, b int) int {
	return a / b
}
