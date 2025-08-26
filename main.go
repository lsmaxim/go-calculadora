package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Calculadora simples em Go")
	fmt.Println("-------------------------")

	// Ler o primeiro número
	fmt.Print("Digite o primeiro número: ")
	input1, _ := reader.ReadString('\n')
	input1 = strings.TrimSpace(input1)
	num1, err := strconv.ParseFloat(input1, 64)
	if err != nil {
		fmt.Println("Número inválido!")
		return
	}

	// Ler o segundo número
	fmt.Print("Digite o segundo número: ")
	input2, _ := reader.ReadString('\n')
	input2 = strings.TrimSpace(input2)
	num2, err := strconv.ParseFloat(input2, 64)
	if err != nil {
		fmt.Println("Número inválido!")
		return
	}

	// Ler a operação
	fmt.Print("Digite a operação (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	var resultado float64
	switch op {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Erro: divisão por zero")
			return
		}
		resultado = num1 / num2
	default:
		fmt.Println("Operação inválida!")
		return
	}

	fmt.Printf("Resultado: %.2f\n", resultado)
}
