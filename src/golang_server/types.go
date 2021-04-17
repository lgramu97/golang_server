package golang_server

import (
	"encoding/json"
	"net/http"
)

// Creamos un tipo middleware para utilizar como cadenas.
// Sirven para encadenar todo antes y dar respuesta si se ejecuta algo o no.
type MiddleWare func(http.HandlerFunc) http.HandlerFunc

// 'json:"name"' para indicar al json como deserializar.
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// Parsea los datos que recive del server a un json.
func (u *User) ToJson() ([]byte, error) {
	return json.Marshal(u) // magia de go.
}

// Interfaz que sirve para ser enviada como parametro al Handler PostRequest.
type MetaData interface{}
