package main

import "net/http"

type Server struct {
	port   string
	router *Router
}

// Creacion de servidor.
func NewServer(port string) *Server {
	return &Server{
		port:   port,
		router: NewRouter(),
	}
}

// Asocia el server a la ruta (path)
func (s *Server) Handle(path string, handler http.HandlerFunc) {
	// A cada regla (path) se le asigna el handler que recive.
	s.router.rules[path] = handler
}

// Agregamos a la parte de Listen del servidor el manejo de handlers mediante el router.
func (s *Server) Listen() error {
	http.Handle("/", s.router) // Endpoint principal.
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		return err
	} else {
		return nil
	}
}
