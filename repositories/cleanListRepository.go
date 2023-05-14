package repositories

type CleanList struct {
}

func (c *CleanList) NumElementos(lista []int) int {
	sinRepetir := make(map[int]bool)

	for _, numero := range lista {
		sinRepetir[numero] = true
	}

	return len(sinRepetir)
}

func NewCleanList() *CleanList {
	return &CleanList{}
}
