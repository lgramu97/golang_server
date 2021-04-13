package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

// Reciver function de calc.
func (calc) operate(entrada, operador string) int {
	entradaLimpia := strings.Split(entrada, operador) // Lista con operandos.
	operador1 := parsear(entradaLimpia[0])
	operador2 := parsear(entradaLimpia[1])

	switch operador {
	case "+":
		fmt.Println(operador1 + operador2)
		return operador1 + operador2
	case "-":
		fmt.Println(operador1 - operador2)
		return operador1 - operador2
	case "*":
		fmt.Println(operador1 * operador2)
		return operador1 * operador2
	case "/":
		fmt.Println(operador1 / operador2)
		return operador1 / operador2
	default:
		fmt.Println("Operación no valida")
		return 0
	}
}

func parsear(x string) int {
	operador, _ := strconv.Atoi(x)
	return operador
}

func leerEntrada() string {
	// Instanciamos un escaner para leer datos.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// OBtengo un string por consola
	return scanner.Text()
}

func main() {
	fmt.Println("Inserte operación:")
	entrada := leerEntrada()
	fmt.Println("La entrada es ", entrada)
	fmt.Println("Inserte operador:")
	operador := leerEntrada()
	fmt.Println("El operador es ", operador)
	c := calc{}.operate(entrada, operador)
	fmt.Printf("El resultado es: %d \n", c)

}
