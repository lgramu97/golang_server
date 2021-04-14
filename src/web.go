package main

import (
	"fmt"
	"io"
	"net/http"
)

type escritorWeb struct{}

func (e escritorWeb) Write(p []byte) (int, error) {
	fmt.Println(string(p))
	return len(p), nil
}

func main() {
	// Protocolo de comunicación con la web.
	response, err := http.Get("http://google.com")
	if err == nil {
		fmt.Println(response.Body)
	}
	// Creo un escritor. Necesito un struct que implemente Write.
	e := escritorWeb{}
	io.Copy(e, response.Body)
}
