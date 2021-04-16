package main

import (
	"fmt"
	"net/http"
)

// Writer: responde a los clientes.
// Request: contiene toda la información.
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World") // Fprintf recibe un escritor y un mensaje-
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint API")
}
