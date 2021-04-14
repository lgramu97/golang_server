package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	canal := make(chan string) // establezco nuevo canal para enviar datos y leer.

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for {
		if i > 2 {
			break
		}
		for _, valor := range servidores {
			go revisarServidor(valor, canal)
		}
		time.Sleep(1 + time.Second)
		fmt.Println()
		i++
	}

	fin := time.Since(inicio)
	fmt.Printf("Finalizado request %s \n", fin)

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	fin = time.Since(inicio)
	fmt.Printf("Finalizado response %s \n", fin)

}

func revisarServidor(servidor string, canal chan string) {
	_, error := http.Get(servidor)
	if error != nil {
		fmt.Println(servidor, "El servidor esta caido.")
		canal <- servidor + " No se encuentra disponible"
	} else {
		fmt.Println(servidor, "El servidor esta funcionando correctamente.")
		canal <- servidor + " Servidor funcionando."
	}
}
