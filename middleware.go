package _go

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// MiddleWare para chequear si esta o no autenticado.
func CheckAuth() MiddleWare {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Logica del middleware
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

func Login() MiddleWare {
	return func(f http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// Logica del middleware
			start := time.Now()
			defer func() { // Se ejecuta al final-
				log.Println(r.URL.Path, time.Since(start))
			}()
			f(w, r)
		} // retorno un handler.
	}
}
