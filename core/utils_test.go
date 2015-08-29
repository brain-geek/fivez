package fivez_core

import "testing"
import "reflect"

// import "fmt"

func TestRevertSlice(t *testing.T) {
	if !reflect.DeepEqual(revertSlice([]int{1, 2, 3, 4, 5}), []int{5, 4, 3, 2, 1}) {
		t.Errorf("revertSlice does not work correctly")
	}
}

func TestImpossibleMoveRowDown(t *testing.T) {
	testRow := []int{0, 0, 3, 4}

	newRow := moveRowDown(testRow)

	if newRow != nil {
		t.Errorf("Expected new row to be nil, from base %v, was %v instead", newRow, testRow)
	}
}

func TestSimpleMoveRowDown(t *testing.T) {
	testRow := []int{1, 2, 3, 0}

	newRow := moveRowDown(testRow)
	expectedRow := []int{0, 1, 2, 3}

	if !reflect.DeepEqual(newRow, expectedRow) {
		t.Errorf("Expected new row to be %v, from base %v, was %v instead", expectedRow, testRow, newRow)
	}
}

func TestSmallMergeMoveRowDown(t *testing.T) {
	testRow := []int{1, 2, 2, 3}

	newRow := moveRowDown(testRow)
	expectedRow := []int{0, 1, 2, 5}

	if !reflect.DeepEqual(newRow, expectedRow) {
		t.Errorf("Expected new row to be %v, from base %v, was %v instead", expectedRow, testRow, newRow)
	}
}

func TestBigMergeMoveRowDown(t *testing.T) {
	testRow := []int{1, 5, 5, 3}

	newRow := moveRowDown(testRow)
	expectedRow := []int{0, 1, 10, 3}

	if !reflect.DeepEqual(newRow, expectedRow) {
		t.Errorf("Expected new row to be %v, from base %v, was %v instead", expectedRow, testRow, newRow)
	}
}
