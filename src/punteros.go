package main

import "fmt"

func main() {
	x := 25
	fmt.Println("Valor de x antes", x)
	fmt.Println("Direccion de x antes", &x)
	cambiarValor(x)
	fmt.Println("Valor de x despues", x)
	fmt.Println("Direccion de x despues", &x)

	y := &x
	fmt.Println("Direccion de y", y)
	fmt.Println("Valor de y ", *y)

}

func cambiarValor(x int) {
	fmt.Println("Direccion de x dentro de cambiar valor", &x)
	x = 36
}
