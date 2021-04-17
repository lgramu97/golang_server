package golang_server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Writer: responde a los clientes.
// Request: contiene toda la información.
func HandleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World") // Fprintf recibe un escritor y un mensaje-
}

func HandleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Endpoint API")
}

// Parseo datos y los envio.
func PostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var metaData MetaData
	err := decoder.Decode(&metaData)
	if err != nil {
		fmt.Fprintf(w, "Error : %v \n", err)
		return
	}
	fmt.Fprintf(w, "Payload %v \n", metaData)
}

func UserPostRequest(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var user User
	err := decoder.Decode(&user)
	if err != nil {
		fmt.Fprintf(w, "Error : %v \n", err)
		return
	}
	response, err := user.ToJson()
	if err != nil {
		// Envio mediante el writer una respuesta al usuario por medio del header.
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	//fmt.Fprintf(w, "Payload %v \n", user)
}
