package main

import (
	"fmt"
	"hackerrank/robot-en-marte/di"
)

/*
QUESTION DESCRIPTION La NASA nos ha contratado y como empresa llevaremos nuestro primer robot a Marte.
Nuestro robot se mueve 1 metro en cada dirección con el comando Izquierda (L) Derecha (R) Arriba (U)
Abajo (D).
La NASA prepara una lista de indicaciones para el movimiento del robot desde la base de carga en el
ejemplo marcada como punto 0.
Sin embargo están preocupados porque en caso de una emergencia el robot pueda regresar a tiempo a la
base de carga y quieren que evaluemos los planes de movimiento en un simulador, y les digamos la
cantidad de instrucciones máximas que deberíamos enviar al robot cuando se encuentre en su punto más
lejano para que pueda retornar a la base.
Calcule cuál es el número máximo de instrucciones que debería enviarse al robot para que en algún punto
del recorrido regrese a la base.
Function Description
Complete la función abajo para completar la tarea requerida, la función tendrá una lista de planes a
ejecutar, evalúe cada uno y retorne una lista con el numero máximo de instrucciones\
*/
func main() {

	handler, err := di.Initialize()
	if err != nil {
		panic("fatal err: " + err.Error())
	}
	instrucciones := []string{"U", "UUU", "RLUUULDDDRRRLLLLLLLUUUUURRRRR"}
	movimientos, err := handler.CalcularMaximoRetorno(instrucciones)

	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	for _, movimiento := range movimientos {
		fmt.Printf("%s", handler.Imprimir(movimiento))
	}
}
