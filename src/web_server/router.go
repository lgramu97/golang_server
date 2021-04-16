package main

import (
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
	// El router intercepta el msj y lo escribe
	// De request extraigo el path y todo para mapear.
	handler, exist := r.FindHandler(request.URL.Path)
	if !exist {
		w.WriteHeader(http.StatusNotFound) // Status de response (404).
		return
	} else {
		handler(w, request)
	}
}

// Busca la ruta y retorna el handler y si existe o no.
func (r *Router) FindHandler(path string) (http.HandlerFunc, bool) {
	handler, exist := r.rules[path]
	return handler, exist
}
