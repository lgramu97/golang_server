package main

import "fmt"

type perro struct {
}

func (p perro) caminar() string {
	return "Soy un perro y camino"
}

type pez struct {
}

func (p pez) nada() string {
	return "Soy un pez y nado"
}

type pajaro struct {
}

func (p pajaro) vuela() string {
	return "Soy un ave y vuelo"
}

func moverPerro(p perro) {
	fmt.Println(p.caminar())
}

func moverPez(p pez) {
	fmt.Println(p.nada())
}

func moverPajaro(p pajaro) {
	fmt.Println(p.vuela())
}

func main() {
	// Codigo repetitivo y redundante.
	perro := perro{}
	moverPerro(perro)
	pez := pez{}
	moverPez(pez)
	pajaro := pajaro{}
	moverPajaro(pajaro)
	// MUY MALA PRACTICA
}
