package main

import "fmt"

// Creamos una interfaz
type animal interface {
	mover() string
}

type perro struct {
}

func (p perro) mover() string {
	return "Soy un perro y camino"
}

type pez struct {
}

func (p pez) mover() string {
	return "Soy un pez y nado"
}

type pajaro struct {
}

func (p pajaro) mover() string {
	return "Soy un ave y vuelo"
}

// Tengo una unica función para todos.
func moverAnimal(a animal) {
	fmt.Println(a.mover())
}

func main() {
	// Codigo limpio y eficiente.
	perro := perro{}
	moverAnimal(perro)
	pez := pez{}
	moverAnimal(pez)
	pajaro := pajaro{}
	moverAnimal(pajaro)

	// No hay un token que indique que se esta implementando la interfaz,
	// simplemente se implementa la función con el mismo nombre.

}
