//  2. Escriba el programa más eficiente que pueda para imprimir en pantalla un rombo

// Para dicho fin, escriba y use una función que reciba de parámetro la cantidad de elementos de la línea del centro,
// la cual debe ser impar positiva.
package main

import (
	"fmt"
	"strings"
)

func main() {
	n := 5 // Cantidad de elementos de la línea del centro
	Rombo(n)
}

func Rombo(n int) {
	if n%2 == 0 || n <= 0 {
		fmt.Println("El número debe ser impar y positivo")
		return
	}
	for i := 1; i <= n; i += 2 {
		fmt.Print(strings.Repeat("   ", (n-i)/2))
		fmt.Println(strings.Repeat(" * ", i))
	}
	for i := n - 2; i >= 1; i -= 2 {
		fmt.Print(strings.Repeat("   ", (n-i)/2))
		fmt.Println(strings.Repeat(" * ", i))
	}
}
