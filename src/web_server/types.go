package main

import "net/http"

// Creamos un tipo middleware para utilizar como cadenas.
// Sirven para encadenar todo antes y dar respuesta si se ejecuta algo o no.
type MiddleWare func(http.HandlerFunc) http.HandlerFunc
