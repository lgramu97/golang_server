package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Instanciamos un escaner para leer datos.
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	// OBtengo un string por consola
	operacion := scanner.Text()
	fmt.Println(operacion)
	// Separo los datos segun el operador
	operador := "ope" // Si introducimos una operación por input que es distinta a la que se define aca, habra un error.
	valores := strings.Split(operacion, operador)
	fmt.Println(valores)
	// Concateno, no estoy haciendo suma aritmetica.
	fmt.Println(valores[0] + valores[1])

	// Convierto los strings en numeros. El primer valor es el resultado, y el siguiente es el mensaje.
	operador1, err1 := strconv.Atoi(valores[0])
	operador2, err2 := strconv.Atoi(valores[1])

	if err1 == nil && err2 == nil {
		// Ahora si puedo operar.
		switch operador {
		case "+":
			fmt.Println(operador1 + operador2)
		case "-":
			fmt.Println(operador1 - operador2)
		case "*":
			fmt.Println(operador1 * operador2)
		case "/":
			fmt.Println(operador1 / operador2)
		default:
			fmt.Println("Operación no valida")
		}
	} else {
		// Cauando Atoi recibe algo que no es un numero, falla y retorna zero_value (0)
		if err1 != nil {
			fmt.Println(operador1, err1)
		} else {
			fmt.Println(operador2, err2)
		}
		fmt.Println("Hubo un error")
	}

}
