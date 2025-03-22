package main

import (
	"math/rand/v2"
	"reflect"
	"testing"
)

const FUZZCOUNT int = 40

// Todo: Implement single bit x,y tests

func TestFlipHorizontalAxis(t *testing.T) {

	if flipHorizontalAxis(boardmaskTop) != boardmaskBottom {
		t.Fatalf("flipHorizontalAxis(boardmaskTop) != boardmaskBottom")
	}
	if flipHorizontalAxis(boardmaskLeft) != boardmaskLeft {
		t.Fatalf("flipHorizontalAxis(boardmaskLeft) != boardmaskLeft")
	}
	// Fuzzing
	for i := 0; i < FUZZCOUNT; i++ {
		rndm := randUint64()
		if flipHorizontalAxis(flipHorizontalAxis(rndm)) != rndm {
			t.Fatalf("flipHorizontalAxisf(lipHorizontalAxis(rndm)) != rndm")
		}
	}
}

func TestFlipVerticalAxis(t *testing.T) {

	if flipVerticalAxis(boardmaskLeft) != boardmaskRight {
		t.Fatalf("flipVerticalAxis(boardmaskTop) != boardmaskBottom")
	}
	if flipVerticalAxis(boardmaskTop) != boardmaskTop {
		t.Fatalf("flipVerticalAxis(boardmaskLeft) != boardmaskLeft")
	}
	// Fuzzing
	for i := 0; i < FUZZCOUNT; i++ {
		rndm := randUint64()
		if flipVerticalAxis(flipVerticalAxis(rndm)) != rndm {
			t.Fatalf("flipVerticalAxisf(lipVerticalAxis(rndm)) != rndm")
		}
	}
}

func TestFlipDiagonalMain(t *testing.T) {

	if flipDiagonalMain(boardmaskLeft) != boardmaskTop {
		t.Fatalf("flipMainDiagonal(boardmaskTop) != boardmaskBottom")
	}
	if flipDiagonalMain(boardmaskRight) != boardmaskBottom {
		t.Fatalf("flipMainDiagonal(boardmaskLeft) != boardmaskLeft")
	}
	if flipDiagonalMain(boardmaskTriangleSubTop) != boardmaskTriangleSubTop {
		t.Fatalf("flipMainDiagonal(boardmaskLeft) != boardmaskLeft")
	}
	// Fuzzing
	for i := 0; i < FUZZCOUNT; i++ {
		rndm := randUint64()
		if flipDiagonalMain(flipDiagonalMain(rndm)) != rndm {
			t.Fatalf("flipMainDiagonalf(lipVerticalAxis(rndm)) != rndm")
		}
		x, y := randTuple(8)
		flipx, flipy := y, x // flip on sub diagonal
		// Todo: Fix this test,.
		var org uint64 = uint64(1) << (y*8 + x)
		var flipped uint64 = uint64(1) << (flipy*8 + flipx)
		if flipDiagonalMain(org) != flipped {
			t.Fatalf("flipDiagonalMain(pos) != flipped_pos, %v != %v (x = %v, y=%v)", org, flipped, x, y)
		}
	}
}

func TestFlipDiagonalSub(t *testing.T) {

	if flipDiagonalSub(boardmaskLeft) != boardmaskBottom {
		t.Fatalf("flipMainDiagonal(boardmaskTop) != boardmaskBottom")
	}
	if flipDiagonalSub(boardmaskRight) != boardmaskTop {
		t.Fatalf("flipMainDiagonal(boardmaskLeft) != boardmaskLeft")
	}
	if flipDiagonalSub(boardmaskTriangleMainTop) != boardmaskTriangleMainTop {
		t.Fatalf("flipMainDiagonal(boardmaskLeft) != boardmaskLeft")
	}
	// Fuzzing
	for i := 0; i < FUZZCOUNT; i++ {
		rndm := randUint64()
		if flipDiagonalSub(flipDiagonalSub(rndm)) != rndm {
			t.Fatalf("flipMainDiagonalf(flipDiagonalSub(rndm)) != rndm")
		}
		x, y := randTuple(8)
		flipx, flipy := 7-y, 7-x // flip on sub diagonal
		// Todo: Fix this test,.
		var org uint64 = uint64(1) << (y*8 + x)
		var flipped uint64 = uint64(1) << (flipy*8 + flipx)
		if flipDiagonalSub(org) != flipped {
			t.Fatalf("flipDiagonalSub(pos) != flipped_pos, %v != %v (x = %v, y=%v)", org, flipped, x, y)
		}
	}
}

func TestInsertSorted(t *testing.T) {

	sortedSlice := []int{1, 3, 5, 7, 9}
	newElement := 4
	sortedSlice = insertSorted(sortedSlice, newElement)
	if !reflect.DeepEqual(sortedSlice, []int{1, 3, 4, 5, 7, 9}) {
		t.Fatalf("Not sorted: %v", sortedSlice)
	}

	sortedSlice = []int{} // test empty slice
	newElement = 8
	sortedSlice = insertSorted(sortedSlice, newElement)
	if !reflect.DeepEqual(sortedSlice, []int{8}) {
		t.Fatalf("Not sorted: %v", sortedSlice)
	}
}

func TestSolutionSets(t *testing.T) {

	var sol PiecePlacement
	sol.mapping = make(map[string][]int)
	sol.mapping["Q"] = []int{1, 2, 3}

	if sol.hash() == 0 {
		t.Fatalf("Hash does not work")
	}

	firstHash := sol.hash()
	sol.insert("Q", 4)
	if sol.hash() == firstHash {
		t.Fatalf("Hash does not update: %v == %v", firstHash, sol.hash())
	}
}

func randUint64() uint64 {
	rndm := uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
	return rndm
}

func randTuple(n int) (int, int) {
	return rand.IntN(n), rand.IntN(n)
}
