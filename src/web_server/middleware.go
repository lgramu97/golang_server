package main

import (
	"fmt"
	"net/http"
)

// MiddleWare para chequear si esta o no autenticado.
func CheckAuth() MiddleWare {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			flag := true
			fmt.Println("Checking authentication")
			if flag { // si pasa el middleware llama al siguiente handler(f).
				f(w, r)
			} else {
				return
			}
		} // retorno un handler.
	}
}
