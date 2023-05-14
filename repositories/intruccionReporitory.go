package repositories

import (
	"fmt"
	"math"
)

type Intrucciones struct {
	Intrucciones  string
	Recorrido     [][]int
	MaximoRetorno int
}

type Intruccion struct{}

func (i *Intruccion) Imprimir(ins Intrucciones) string {
	return fmt.Sprintf("Intruccion %s \n Recorrido %v \n MaximoRetorno %d \n", ins.Intrucciones, ins.Recorrido, ins.MaximoRetorno)
}

func (i *Intruccion) CalcularMaximoRetorno(instrucciones []string) ([]Intrucciones, error) {
	movimientos := make([]Intrucciones, 0, len(instrucciones))

	for _, instruccion := range instrucciones {
		x := 0
		y := 0
		maximoRecorrido := 0
		movimientosAct := make([][]int, 0, len(instruccion))

		for _, movimiento := range instruccion {
			switch string(movimiento) {
			case "R":
				x++
			case "L":
				x--
			case "U":
				y++
			case "D":
				y--
			}
			recorridoActual := int(math.Abs(float64(x)) + math.Abs(float64(y)))
			if recorridoActual > maximoRecorrido {
				maximoRecorrido = recorridoActual
			}
			movimientosAct = append(movimientosAct, []int{x, y})
		}
		movimientos = append(movimientos, Intrucciones{
			Intrucciones:  instruccion,
			Recorrido:     movimientosAct,
			MaximoRetorno: maximoRecorrido,
		})
	}

	return movimientos, nil
}

func NewIntruccionRepository() *Intruccion {
	return &Intruccion{}
}
