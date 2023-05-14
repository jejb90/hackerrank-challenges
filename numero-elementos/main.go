package main

import (
	"fmt"
	"hackerrank/numero-elementos/di"
)

/*
A tu equipo de desarrollo el cliente les ha pedido realizar un algoritmo de indexación y optimización de
almacenamiento en el proceso de inventario.
Sin embargo el equipo tiene problemas con un método que no saben como realizar de forma óptima y te
han pedido ayuda para resolverlo.
Siguiendo el Principio de responsabilidad única (SRP), vas a construir una función que tiene como única
responsabilidad calcular la cantidad de números diferentes dentro de una lista dada.
Function Description
Complete la función en el editor abajo, la función recibe la lista de tareas, y debe retornar la cantidad de
valores únicos que existen.
*/

func main() {

	handler, err := di.Initialize()
	if err != nil {
		panic("fatal err: " + err.Error())
	}
	listaNumeros := []int{1, 2, 3, 2, 4, 5, 1, 3, 100, 101, 100}
	fmt.Println(handler.NumElementos(listaNumeros))
}
