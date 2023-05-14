package ctx

type CleanList interface {
	NumElementos(lista []int) int
}

type Handler struct {
	CleanList CleanList
}

func (h *Handler) NumElementos(lista []int) int {
	return h.CleanList.NumElementos(lista)
}

func NewHandler(cleanList CleanList) *Handler {
	return &Handler{
		CleanList: cleanList,
	}
}
