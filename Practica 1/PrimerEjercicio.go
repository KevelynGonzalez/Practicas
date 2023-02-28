// 1)Haga un programa que cuente e indique el número de caracteres, el número de palabras y el
// número de líneas de un texto cualquiera (obtenido de cualquier forma que considere).
// Considere hacer el cálculo de las tres variables en el mismo proceso.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("indique el texto: ")
	reader := bufio.NewReader(os.Stdin)
	texto, _ := reader.ReadString('\n')

	numCaracteres := len(texto) //los espacios tambien son contados

	palabras := strings.Fields(texto)
	numPalabras := len(palabras)

	lineas := bufio.NewScanner(strings.NewReader(texto))
	numLineas := 0
	for lineas.Scan() {
		numLineas++
	}

	fmt.Printf(" Cantidad de caracteres: %d\n Cantidad de palabras: %d\n Número de líneas: %d\n", numCaracteres, numPalabras, numLineas)

}
