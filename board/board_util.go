package board

import (
	"strconv"
	"strings"
)

func RowColToAlgebra(row, col int) string {
	rowChar := rune('a' + row)
	col = col + 1

	return string(rowChar) + strconv.Itoa(col)
}

func AlgebraToRowCol(c string) []int {
	cSplit := strings.Split(c, "")
	rowChar := cSplit[0]
	col, _ := strconv.Atoi(cSplit[1])

	row := int([]rune(rowChar)[0] - 'a')

	return []int{row, col - 1}
}
