package main

import (
	"fmt"
	"golang.org/x/exp/slices"
	"sort"
)

type producto struct {
	nombre         string
	descripcion    string
	montoVenta     int32
	existencias    int32 // Nuevo campo agregado
	cantidadMinima int32
}

type Productos []producto

var factura Productos
var productos Productos
var listaMinimos Productos

const rangoPagoImpuestos = 20000
const porcentajeImpuesto = 0.13

func (f *Productos) agregarProducto(nom string, desc string, pre int32, exist int32, minimo int32) {
	idx := slices.IndexFunc(*f, func(p producto) bool { return p.nombre == nom })
	if idx == -1 {
		*f = append(*f, producto{nom, desc, pre, exist, minimo})
	} else {
		fmt.Println("no se puede añadir productos a la lista")
	}
}

func (f *Productos) calcularImpuestoFactura() int32 {
	lista := filter1(*f, func(p producto) bool {
		return p.montoVenta > rangoPagoImpuestos
	})
	lista2 := map1(lista, func(p producto) int32 {
		return int32(float64(p.montoVenta) * porcentajeImpuesto)
	})
	return reduce(lista2)
}

func (f *Productos) calcularMontoFactura() int32 {
	lista := map1(*f, func(p producto) int32 {
		return p.montoVenta
	})
	return reduce(lista)
}
func (f *Productos) calcularMontoFactura2() int32 {
	lista := map1(*f, func(p producto) int32 {
		return p.montoVenta
	})
	return reduce(lista)
}

func filter1[T any](slice []T, f func(T) bool) []T {
	elementoFiltrado := make([]T, 0)
	for _, element := range slice {
		if f(element) {
			elementoFiltrado = append(elementoFiltrado, element)
		}
	}
	return elementoFiltrado
}

// Map para slices de cualquier tipo
func map1[T, R any](slice []T, f func(T) R) []R {
	mapp := make([]R, len(slice))
	for i, e := range slice {
		mapp[i] = f(e)
	}
	return mapp
}
func reduce(list []int32) int32 {
	var result int32 = 0
	for _, v := range list {
		result += v
	}
	return result
}
func (l *Productos) aumentarInventario(listaMinimos []producto) {
	for _, p := range listaMinimos {
		unidadesFaltantes := p.cantidadMinima - p.existencias
		if unidadesFaltantes > 0 {
			p.existencias += unidadesFaltantes
			for i, prod := range *l {
				if prod.nombre == p.nombre {
					(*l)[i] = p
					break
				}
			}
		}
	}
}
func (l *Productos) imprimirProductos() {
	fmt.Println("Lista de productos actualizada:")
	for _, p := range *l {
		fmt.Printf("Nombre: %s\nDescripción: %s\nMonto de venta: %d\nExistencias: %d\nCantidad mínima: %d\n\n", p.nombre, p.descripcion, p.montoVenta, p.existencias, p.cantidadMinima)
	}
}
func (l *Productos) productosOrdenados() {
	sort.Slice(*l, func(i, j int) bool {
		return (*l)[i].nombre < (*l)[j].nombre
	})
}

func main() {
	factura.agregarProducto("tarjeta madre", "Asus", 54200, 10000, 20000)
	factura.agregarProducto("mouse", "alámbrico", 15000, 1500, 10)
	factura.agregarProducto("teclado", "gamer con luces", 30000, 20, 2)
	factura.agregarProducto("memoria ssd", "2 gb", 65000, 80000, 90000)
	factura.agregarProducto("cable video", "display port 1m", 18000, 54, 100)
	println(factura.calcularMontoFactura())
	println(factura.calcularImpuestoFactura())

	///////////////////////////////////////////////////////////////////////////////////////////////////////////////

	productos.agregarProducto("tarjeta madre", "Asus", 54200, 10000, 20000)
	productos.agregarProducto("mouse", "alámbrico", 15000, 1500, 10)
	productos.agregarProducto("teclado", "gamer con luces", 30000, 20, 2)
	productos.agregarProducto("memoria ssd", "2 gb", 65000, 80000, 90000)
	productos.agregarProducto("cable video", "display port 1m", 18000, 54, 100)
	fmt.Println("---------------------------Ejercicio de repositorio---------------------")
	minimos := filter1(factura, func(p producto) bool {
		return p.existencias < p.cantidadMinima
	})

	productos.aumentarInventario(minimos)
	nombreProducto := map1(minimos, func(p producto) string {
		return p.nombre
	})
	fmt.Println("------lista de productos con mínimas existencias de inventario-------")
	fmt.Println(nombreProducto)

	fmt.Println("\n\n--------Lista actualizada de productos-----")
	fmt.Println(productos)
	fmt.Println("\n\n-----------------Lista de productos ordenada alfabeticamente-------------------- ")
	productos.productosOrdenados()
	productos.imprimirProductos()

}
