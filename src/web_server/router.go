package golang_server

import (
	"net/http"
)

// Mapea que hanlder va con que codigo.
type Router struct {
	rules map[string]map[string]http.HandlerFunc
}

// Ahora tengo un mapa de path y metodos, con un handler
func NewRouter() *Router {
	return &Router{
		rules: make(map[string]map[string]http.HandlerFunc),
	}
}

// Se encarga de entregar respuestas a las rutas. Implementa INTERFACE para ser capaz de manejar el request.
func (r *Router) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	// El router intercepta el msj y lo escribe
	// De request extraigo el path y todo para mapear.
	handler, method_exist, exist := r.FindHandler(request.URL.Path, request.Method)
	if !exist {
		w.WriteHeader(http.StatusNotFound) // Status de response (404).
		return
	}

	if !method_exist {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	handler(w, request)

}

// Busca la ruta y retorna el handler y si existe o no (asociado al metodo).
func (r *Router) FindHandler(path string, method string) (http.HandlerFunc, bool, bool) {
	_, exist := r.rules[path]
	handler, existMethod := r.rules[path][method]
	return handler, existMethod, exist
}
