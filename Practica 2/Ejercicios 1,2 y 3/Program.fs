module Ejemplos2 =

   
    //=========Ejercicios de repositorio===================
     
    //1)	Haciendo uso de la función filter, implemente una función que a partir
    //de una lista de cadenas representando frases de parámetro, filtre aquellas que
    //contengan una palabra que el usuario indique en otro argumento (palabra completa contenida). 
    let filtrarPalabras (pal:string) (lista:string list) =
        lista
        |> List.filter (fun str -> str.Contains(pal))
  
    //2)	Construya una función que se llame general-sec A B. Esta función genera una
    //de números en una lista de A hasta B, donde  A y B son números enteros
    let rec generar_sec A B =
        if A > B then []
        else A :: generar_sec (A+1) B
    //3)Migrar el ejercicio realizado en Go sobre facturas con listas de productos haciendo
    //énfasis en la implementación de las funciones para calcular el monto a pagar por la factura completa
    //y el monto a pagar por concepto del 13% de impuesto de venta para aquellos productos que deban
    //pagarlo según criterio de selección.
    open System.Reflection

    type producto = { nombre : string; descripcion : string; montoVenta : int32 }
    type Productos = producto list

    let mutable factura = []

    let rangoPagoImpuestos = 20000
    let porcentajeImpuesto = 0.13

    let agregarProducto (nom : string) (desc : string) (pre : int32) =
        let idx = factura |> List.tryFindIndex (fun p -> p.nombre = nom)
        match idx with
        | Some _ -> printfn "cant add existing product to the list"
        | None -> factura <- factura @ [{ nombre = nom; descripcion = desc; montoVenta = pre }]

    let calcularMontoFactura () =
        factura |> List.map (fun p -> p.montoVenta) |> List.sum

    let calcularMontoFactura2 () =
        factura |> List.fold (fun acc p -> acc + p.montoVenta) 0

    let calcularImpuestoFactura () =
        let lista = factura |> List.filter (fun p -> p.montoVenta > rangoPagoImpuestos)
        let lista2 = lista |> List.map (fun p -> int32(float p.montoVenta * float porcentajeImpuesto) |> int32)

        lista2 |> List.sum




   
open Ejemplos2

let frases = ["la casa";"el perro";"pintando la cerca"]
printfn "%s" "\nRespuesta de ejercicio 1"
printfn "%A" (filtrarPalabras "la" frases)
printfn "%s" "\nRespuesta de ejercicio 2"
printfn "%A" (generar_sec 0 5)
// Ejemplo de uso
agregarProducto "tarjeta madre" "Asus" 54200
agregarProducto "mouse" "alámbrico" 15000
agregarProducto "teclado" "gamer con luces" 30000
agregarProducto "memoria ssd" "2 gb" 65000
agregarProducto "cable video" "display port 1m" 18000
printfn "%s" "\nRespuesta de ejercicio 3"
printfn "%s" "Monto factura"
printfn "%d" (calcularMontoFactura ())
printfn "%s" "Impuesto"
printfn "%d" (calcularImpuestoFactura ())