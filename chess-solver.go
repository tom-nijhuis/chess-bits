package main

import (
	"fmt"
	"sort"
)

// todo: use profiler. https://go.dev/blog/pprof

const USE_SYMMETRY bool = false // symmetry does not work with sorted queen positions.
const FIND_ALL bool = true

type PiecePlacement struct {
	mapping map[string][]int
	hashVal uint64
}

func (piecePlaces *PiecePlacement) hash() uint64 {
	hashMagics := make(map[string]uint64)
	hashMagics["Q"] = 0x3EE136815DD17C6C
	hashMagics["K"] = 0x499250A6ECBC86F1
	hashMagics["R"] = 0x676451E4C0A59376

	if piecePlaces.hashVal == 0 {
		for piece, location := range piecePlaces.mapping {
			piecePlaces.hashVal ^= hashMagics[piece] // xor with the seed for this piece
			for _, ind := range location {
				piecePlaces.hashVal ^= hashMagics[piece]<<ind | hashMagics[piece]>>(64-ind) // bit shift with roll around
			}
		}
	}

	// Todo: Capture symmetry in the hashes only.

	return piecePlaces.hashVal
}

func (piecePlaces *PiecePlacement) deepCopy() PiecePlacement {
	newPiecePlacement := PiecePlacement{
		mapping: make(map[string][]int),
		hashVal: 0,
	}

	for k, srcIntSlice := range piecePlaces.mapping {
		dstIntSlice := make([]int, len(srcIntSlice))
		copy(dstIntSlice, srcIntSlice)
		newPiecePlacement.mapping[k] = dstIntSlice
	}
	return newPiecePlacement
}

func (piecePlaces *PiecePlacement) insert(piece string, position int) {
	piecePlaces.mapping[piece] = append(piecePlaces.mapping[piece], position) // not sorted.
	piecePlaces.hashVal = 0

}

type SolutionSet struct {
	MapList        []*PiecePlacement
	RecursiveCalls int
}

func (solSet *SolutionSet) addSolution(solution *PiecePlacement) {
	solSet.MapList = append(solSet.MapList, solution)
}

func flipHorizontalAxis(x uint64) uint64 {
	var y uint64 = 0
	for i := 0; i < 8; i++ {
		row := (x >> (i * 8)) & 0xff //and with the top row
		y += row << ((8 - i - 1) * 8)
	}
	return y
}

func flipVerticalAxis(x uint64) uint64 {
	var y uint64 = 0
	for i := 0; i < 8; i++ {
		col := x >> i & 0x0101010101010101 // and with the least significant column
		y += col << (8 - i - 1)
	}
	return y
}

func flipDiagonalMain(x uint64) uint64 {

	var y uint64 = 0
	// Todo: Make this a loop of 8 instead of 64?
	for bitindex := 0; bitindex < 64; bitindex++ {
		var bitvalue uint64 = (x >> bitindex) & 1
		newx := bitindex / 8 // newx = y
		newy := bitindex % 8 // newy = x
		newbitindex := 8*newy + newx
		y = y + (bitvalue << newbitindex)
	}
	return y
}

func flipDiagonalSub(x uint64) uint64 {

	var y uint64 = 0
	for bitindex := 0; bitindex < 64; bitindex++ {
		var bitvalue uint64 = (x >> bitindex) & 1
		newx := 7 - bitindex/8 // newx = 7-y
		newy := 7 - bitindex%8 // newy = 7-x
		newbitindex := 8*newy + newx
		y = y + (bitvalue << newbitindex)
	}
	return y
}

func printUint(x uint64) {
	fmt.Printf("+--------+(0x%016x)", x)
	str := "\n|"
	for row := 0; row < 8; row++ {
		for col := 0; col < 8; col++ {
			if x&1 == 1 {
				str += "█"
			} else {
				str += "·"
			}
			x = x >> 1
		}
		str += "|\n"
		if row < 7 {
			str += "|"
		}
	}
	str += "+--------+\n"
	fmt.Printf("%v", str)
}

func copyPositionMap(positionMap map[string][]int) map[string][]int {
	newPositionMap := make(map[string][]int)
	for k, srcIntSlice := range positionMap {
		dstIntSlice := make([]int, len(srcIntSlice))
		copy(dstIntSlice, srcIntSlice)
		newPositionMap[k] = dstIntSlice
	}
	return newPositionMap
}

func insertSorted(slice []int, newElement int) []int {
	index := sort.Search(len(slice), func(i int) bool { return slice[i] >= newElement })

	slice = append(slice, 0)             // Expand the slice by one element
	copy(slice[index+1:], slice[index:]) // Shift elements to the right
	slice[index] = newElement            // Insert the new element at the correct index

	return slice
}

func Solve(piecesRemaining []string, Solutions *SolutionSet, positionMap PiecePlacement, openBoard uint64, skipHashMap *map[uint64]bool) (
	solved bool) {

	Solutions.RecursiveCalls += 1

	// Todo :  add the Hash to the attemptsList of the Solution set and don't try them again.

	var anySolution bool = false

	stringPiece := map[string][64]uint64{
		"Q":  maskQueen,
		"R":  maskRook,
		"B":  maskBishop,
		"K":  maskKing,
		"Kn": maskKnight,
	}

	if len(piecesRemaining) == 0 {
		// Base case, solved!
		// Add to solutions list
		(*skipHashMap)[positionMap.hash()] = true // mark it as visited
		Solutions.addSolution(&positionMap)

		return true
	}
	if openBoard == 0xffffffffffffffff {
		// the board has no open positions. Close this search branch
		return false
	}

	// Take a piece from the stack
	pieceString := piecesRemaining[0]

	// todo: Identify symmetries and mask of the proper halves.
	var symmetryMask uint64 = 0x0 // todo: Capture symmetry in the hashes only.
	if USE_SYMMETRY {
		// Careful! Does not work with pawns.
		if openBoard|symmetryMask == flipHorizontalAxis(openBoard|symmetryMask) {
			// The board is symmetrical in the horizontal axis. Mask off the bottom.
			symmetryMask |= boardmaskTop
		}
		if openBoard|symmetryMask == flipVerticalAxis(openBoard|symmetryMask) {
			// The board is symmetrical in the vertical axis. Mask off the left.
			symmetryMask |= boardmaskRight
		}
		// Todo: Fix diagonal symmetries
	}

	// Add to the first open position in the open board
	for i := 0; i < 64; i++ {
		if ((openBoard|symmetryMask)>>i)&1 == 0 { // there is an open position on i
			// Add the piece to the new positionMap
			recmap := positionMap.deepCopy()
			recmap.insert(pieceString, i)
			if !(*skipHashMap)[recmap.hash()] { // test if this placement was attempted before
				// Recurse
				// Todo optimisation: Earlier branche cuts: We don't have to continue trying open positions anymore if they didn't lead to any solutions with fewer pieces on the board earlier.
				recsolved := Solve(
					piecesRemaining[1:], // take off the first piece, pass on the remainder
					Solutions,
					recmap,
					openBoard|stringPiece[pieceString][i],
					skipHashMap,
				)
				if recsolved {
					// This branch found a solution downstream.
					anySolution = true
					if !FIND_ALL { // if not trying to add all, return straight away,
						return true
					}
				}
			}
		}
	}

	// Return whether or not a solution lives downstream
	if anySolution {
		return true
	}
	// no solutions downstrean, mark this branch as skippable in the future
	(*skipHashMap)[positionMap.hash()] = true // mark it as visited
	return false

}

func main() {

	func() {
		pieces := []string{}
		for i := 0; i < 9; i++ {
			pieces = append(pieces, "k")
		}
		solutions := SolutionSet{
			MapList:        []*PiecePlacement{},
			RecursiveCalls: 0,
		}
		skipHashMap := make(map[uint64]bool)
		Solve(
			pieces,
			&solutions,
			PiecePlacement{},
			0,
			&skipHashMap,
		)
		fmt.Printf("%v Pieces: solutions: %v, calls: %v\n", len(pieces), len(solutions.MapList), solutions.RecursiveCalls)
		fmt.Printf("Skipmap: %v\n", len(skipHashMap))
		// for k, v := range solutions.MapList {
		// 	fmt.Printf("Solution: %v - %v (%v)\n", k, v, v.hashVal)
		// }
		fmt.Println()
	}()

}
