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
func (s *Server) Handle(method string, path string, handler http.HandlerFunc) {
	// A cada regla (path) se le asigna el handler que recive.
	_, exist := s.router.rules[path]
	if !exist {
		s.router.rules[path] = make(map[string]http.HandlerFunc)
	}
	s.router.rules[path][method] = handler
}

// Encadenamiento de middleware, puedo ejecutar el handler si es que paso todos los middleware.
//permite cadena de middlewares{
func (s *Server) AddMiddleware(f http.HandlerFunc, middlewares ...MiddleWare) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
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
