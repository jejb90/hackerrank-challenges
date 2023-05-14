package ctx

import "hackerrank/repositories"

type Intruccion interface {
	Imprimir(intruccion repositories.Intrucciones) string
	CalcularMaximoRetorno(intrucciones []string) ([]repositories.Intrucciones, error)
}
type Handler struct {
	Intruccion Intruccion
}

func (h *Handler) CalcularMaximoRetorno(instrucciones []string) ([]repositories.Intrucciones, error) {
	return h.Intruccion.CalcularMaximoRetorno(instrucciones)
}
func (h *Handler) Imprimir(i repositories.Intrucciones) string {
	return h.Intruccion.Imprimir(i)
}

func NewHandler(intruccion Intruccion) *Handler {
	return &Handler{
		Intruccion: intruccion,
	}
}
