package matrix

func CreateMatrix(matrixLength int) [][]int8 {
	x := matrixLength
	y := matrixLength
	table := make([][]int8, x)
	for i := range table {
		table[i] = make([]int8, y)
	}
	return table

}
