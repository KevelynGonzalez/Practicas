package main

import (
	"fmt"
)

// // Probar su funcionamiento creando una lista/slice de personas y filtrando aquellas personas que sean mayores de edad
// Filter para slices de cualquier tipo
func Filter[T any](slice []T, f func(T) bool) []T {
	elementoFiltrado := make([]T, 0)
	for _, element := range slice {
		if f(element) {
			elementoFiltrado = append(elementoFiltrado, element)
		}
	}
	return elementoFiltrado
}

// Map para slices de cualquier tipo
func Map[T, R any](slice []T, f func(T) R) []R {
	mapp := make([]R, len(slice))
	for i, e := range slice {
		mapp[i] = f(e)
	}
	return mapp
}

// Ejemplo de uso: lista de personas filtrando las mayores de edad usando map/filter
type Personas struct {
	Nombre string
	Edad   int
}

func main() {
	personas := []Personas{
		{Nombre: "Samantha", Edad: 19},
		{Nombre: "Katherine", Edad: 14},
		{Nombre: "Andrey", Edad: 12},
		{Nombre: "Cristhian", Edad: 25},
		{Nombre: "Jonathan", Edad: 10},
		{Nombre: "Jorge", Edad: 40},
	}

	adultos := Filter(personas, func(p Personas) bool {
		return p.Edad >= 18
	})

	names := Map(adultos, func(p Personas) string {
		return p.Nombre
	})

	fmt.Println("Personas que ya son mayores de 18:", adultos)
	fmt.Println("\nNombres de las personas mayores de 18:", names)
}
