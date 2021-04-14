package main

import "fmt"

type taskList struct {
	listaTareas []*task // Agregamos el elemento original, no una copia.
}

func (tl *taskList) agregarLista(t *task) {
	tl.listaTareas = append(tl.listaTareas, t)
}

func (tl *taskList) eliminarTarea(index int) {
	// Los ... indica que envio mas de un elemento a agregar como parametro.
	tl.listaTareas = append(tl.listaTareas[:index], tl.listaTareas[index+1:]...)
}

func (tl *taskList) imprimirTareas() {
	for i := 0; i < len(tl.listaTareas); i++ {
		fmt.Println(tl.listaTareas[i])
	}
}

func (tl *taskList) imprimirTareasRange() {
	for _, value := range tl.listaTareas {
		fmt.Println(value)
		//fmt.Println(value)
	}
}

func (tl *taskList) imprimirListaCompletados() {
	for _, tarea := range tl.listaTareas {
		if tarea.completado {
			fmt.Println(tarea)
		}
	}
}

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
	t1 := &task{
		nombre:      "Mi curso Go",
		descripcion: "Completar curso de go implementando lista tareas",
	}

	/*** PRUEBA PUNTEROS
	fmt.Println(t1)
	t1.actualizarDescripcionCopia("Nueva descripcion")
	fmt.Println(t1)
	// La descripción no cambio! Esto se debe a que se utilizo una copia y no la referencia
	// al struct task declarado en main.
	// Modifiquemos las reciver functions y que se declaren mediante *.
	t1.actualizarDescripcion("Se actualizo ahora si")
	fmt.Println(t1)
	***** FIN PRUEBA PUNTEROS ********/

	t2 := &task{
		nombre:      "Mi Tarea 2 ",
		descripcion: "Segunda tarea",
	}

	t3 := &task{
		nombre:      "Mi Tarea 3 ",
		descripcion: "Tercera tarea",
	}

	// t1, t2 y t3 son etiquetas (contienen la direccion a la tarea) al igual que tl a la lista_tareas.

	tl := &taskList{
		listaTareas: []*task{
			t1, t2,
		},
	}

	tl.agregarLista(t3)
	tl.imprimirTareas()
	fmt.Println("Cantidad de elementos en lista tareas", len(tl.listaTareas))
	tl.eliminarTarea(0)
	fmt.Println("Cantidad de elementos en lista tareas", len(tl.listaTareas))
	tl.imprimirTareasRange()
	tl.listaTareas[0].marcarCompleta()
	fmt.Println("Tareas completadas:")
	tl.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)
	mapaTareas["Lautaro"] = tl

	t4 := &task{
		nombre:      "Mi Tarea 4 ",
		descripcion: "Cuarta tarea",
	}

	t5 := &task{
		nombre:      "Mi Tarea 5 ",
		descripcion: "Quinta tarea",
	}

	lista := &taskList{
		listaTareas: []*task{
			t4, t5,
		},
	}

	mapaTareas["Otra_Persona"] = lista

	fmt.Println("Tareas Lautaro")
	mapaTareas["Otra_Persona"].imprimirTareas()

}
