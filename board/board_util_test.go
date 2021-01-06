package board

import "testing"

func TestRowColToAlgebra(t *testing.T) {
	if RowColToAlgebra(0, 0) != "a1" {
		t.Errorf("invalid output:%s\n", RowColToAlgebra(0, 0))
	}
	if RowColToAlgebra(7, 7) != "h8" {
		t.Errorf("invalid output:%s\n", RowColToAlgebra(7, 7))
	}
}

func TestAlgebraToRowCol(t *testing.T) {
	for _, i := range AlgebraToRowCol("a1") {
		if i != 0 {
			t.Errorf("invalid conversion of a1: %d\n", i)
		}
	}

	for _, i := range AlgebraToRowCol("h8") {
		if i != 7 {
			t.Errorf("invalid conversion of h8: %d\n", i)
		}
	}
}
