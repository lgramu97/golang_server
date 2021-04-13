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
	// Separo los datos segun el operador "+"
	valores := strings.Split(operacion, "+")
	fmt.Println(valores)
	// Concateno, no estoy haciendo suma aritmetica.
	fmt.Println(valores[0] + valores[1])

	// Convierto los strings en numeros.
	operador1, _ := strconv.Atoi(valores[0])
	operador2, _ := strconv.Atoi(valores[1])
	// Ahora si puedo sumar.
	fmt.Println(operador1 + operador2)

}
