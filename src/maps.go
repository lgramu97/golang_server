package main

import "fmt"

func main() {
	m1 := make(map[string]int) // Key: string, Value: int
	m1["a"] = 8
	m1["b"] = 3
	fmt.Println(m1["a"]) // Imprimo solo el valor
}
