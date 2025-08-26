package main

import (
	"fmt"
	"os"
	"strconv"

)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Uso: calc <num1> <num2> <operação>")
		return
	}

	// Lê os argumentos
	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	num2, err2 := strconv.ParseFloat(os.Args[2], 64)
	op := os.Args[3]

	if err1 != nil || err2 != nil {
		fmt.Println("Erro: números inválidos")
		return
	}

	// Faz o cálculo com base na operação
	switch op {
	case "+":
		fmt.Printf("%.2f\n", num1+num2)
	case "-":
		fmt.Printf("%.2f\n", num1-num2)
	case "*":
		fmt.Printf("%.2f\n", num1*num2)
	case "/":
		if num2 != 0 {
			fmt.Printf("%.2f\n", num1/num2)
		} else {
			fmt.Println("Erro: divisão por zero")
		}
	default:
		fmt.Println("Operação inválida. Use +, -, * ou /")
	}
}
