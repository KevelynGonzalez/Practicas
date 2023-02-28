package main

import "fmt"

func rotate(secuencia *[]string, n int, direccion int) {
	length := len(*secuencia)
	if length == 0 || n%length == 0 {
		return
	}
	if n < 0 {
		n = length - (-n % length)
	} else {
		n = n % length
	}
	if direccion == 0 { // hace que rote a la izq
		*secuencia = append((*secuencia)[n:], (*secuencia)[:n]...)
	} else { // hace que rote a la derecha
		*secuencia = append((*secuencia)[length-n:], (*secuencia)[:length-n]...)
	}
}

func main() {
	secuencia := []string{"a", "b", "c", "d", "e", "f", "g", "h"}
	fmt.Println("Secuencia inicial:", secuencia)
	rotate(&secuencia, 3, 0)

	fmt.Println("rotacion a la izquierda:", secuencia)
	rotate(&secuencia, 2, 1)
	fmt.Println("rotacion a la derecha:", secuencia)
}
