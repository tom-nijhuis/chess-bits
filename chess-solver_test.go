package main

import (
	"reflect"
	"testing"
)

func TestPiecePlacement_hash(t *testing.T) {
	type fields struct {
		mapping map[string][]int
		hashVal uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			piecePlaces := &PiecePlacement{
				mapping: tt.fields.mapping,
				hashVal: tt.fields.hashVal,
			}
			if got := piecePlaces.hash(); got != tt.want {
				t.Errorf("PiecePlacement.hash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPiecePlacement_deepCopy(t *testing.T) {
	type fields struct {
		mapping map[string][]int
		hashVal uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   PiecePlacement
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			piecePlaces := &PiecePlacement{
				mapping: tt.fields.mapping,
				hashVal: tt.fields.hashVal,
			}
			if got := piecePlaces.deepCopy(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PiecePlacement.deepCopy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPiecePlacement_insert(t *testing.T) {
	type fields struct {
		mapping map[string][]int
		hashVal uint64
	}
	type args struct {
		piece    string
		position int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			piecePlaces := &PiecePlacement{
				mapping: tt.fields.mapping,
				hashVal: tt.fields.hashVal,
			}
			piecePlaces.insert(tt.args.piece, tt.args.position)
		})
	}
}

func TestSolutionSet_addSolution(t *testing.T) {
	type fields struct {
		MapList        []*PiecePlacement
		RecursiveCalls int
	}
	type args struct {
		solution *PiecePlacement
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			solSet := &SolutionSet{
				MapList:        tt.fields.MapList,
				RecursiveCalls: tt.fields.RecursiveCalls,
			}
			solSet.addSolution(tt.args.solution)
		})
	}
}

func Test_flipHorizontalAxis(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipHorizontalAxis(tt.args.x); got != tt.want {
				t.Errorf("flipHorizontalAxis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flipVerticalAxis(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipVerticalAxis(tt.args.x); got != tt.want {
				t.Errorf("flipVerticalAxis() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flipDiagonalMain(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipDiagonalMain(tt.args.x); got != tt.want {
				t.Errorf("flipDiagonalMain() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flipDiagonalSub(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipDiagonalSub(tt.args.x); got != tt.want {
				t.Errorf("flipDiagonalSub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_printUint(t *testing.T) {
	type args struct {
		x uint64
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printUint(tt.args.x)
		})
	}
}

func Test_copyPositionMap(t *testing.T) {
	type args struct {
		positionMap map[string][]int
	}
	tests := []struct {
		name string
		args args
		want map[string][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := copyPositionMap(tt.args.positionMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("copyPositionMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_insertSorted(t *testing.T) {
	type args struct {
		slice      []int
		newElement int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertSorted(tt.args.slice, tt.args.newElement); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSolve(t *testing.T) {
	type args struct {
		piecesRemaining []string
		Solutions       *SolutionSet
		positionMap     PiecePlacement
		openBoard       uint64
		skipHashMap     *map[uint64]bool
	}
	tests := []struct {
		name       string
		args       args
		wantSolved bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotSolved := Solve(tt.args.piecesRemaining, tt.args.Solutions, tt.args.positionMap, tt.args.openBoard, tt.args.skipHashMap); gotSolved != tt.wantSolved {
				t.Errorf("Solve() = %v, want %v", gotSolved, tt.wantSolved)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
