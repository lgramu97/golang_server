package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t task) String() string {
	return fmt.Sprintf("La tarea con nombre %s dice: %s", t.nombre, t.descripcion)
}

func (t task) marcarCompletaCopia() {
	t.completado = true
}

func (t task) actualizarDescripcionCopia(desc string) {
	t.descripcion = desc
}

func (t task) actualizarNombreCopia(name string) {
	t.nombre = name
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(desc string) {
	t.descripcion = desc
}

func (t *task) actualizarNombre(name string) {
	t.nombre = name
}

func main() {
	t := task{
		nombre:      "Mi curso Go",
		descripcion: "Completar curso de go implementando lista tareas",
	}

	fmt.Println(t)

	t.actualizarDescripcionCopia("Nueva descripcion")
	fmt.Println(t)
	// La descripción no cambio! Esto se debe a que se utilizo una copia y no la referencia
	// al struct task declarado en main.
	// Modifiquemos la reciver functions y que se declaren mediant *.
	t.actualizarDescripcion("Se actualizo ahora si")
	fmt.Println(t)

}
