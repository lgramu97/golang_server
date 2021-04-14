package main

import (
	"fmt"
	"net/http"
)

// Mapea que hanlder va con que codigo.
type Router struct {
	rules map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{
		rules: make(map[string]http.HandlerFunc),
	}
}

// Se encarga de entregar respuestas a las rutas. Implementa INTERFACE para ser capaz de manejar el request.
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(w, "Hello World!") // El router intercepta el msj y lo escribe
}
