package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, valor := range servidores {
		revisarServidor(valor)
	}

	fin := time.Since(inicio)

	fmt.Printf("Finalizado %s \n", fin)

}

func revisarServidor(servidor string) {
	_, error := http.Get(servidor)
	if error != nil {
		fmt.Println(servidor, "El servidor esta caido.")
	} else {
		fmt.Println(servidor, "El servidor esta funcionando correctamente.")
	}
}
