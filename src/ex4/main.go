package main

import (
	"fmt"
	"sort"
)

func main() {
	irregularMatrix := [][]int{{1, 2, 3, 4},
		{5, 5, 7, 8},
		{9, 11, 11},
		{12, 13, 15, 14},
		{16, 17, 18, 19, 20}}

	slice := UniqueInts(irregularMatrix)
	fmt.Printf("1x%d: %v\n", len(slice), slice)

	slice = Flatten(irregularMatrix)
	fmt.Printf("1x%d: %v\n", len(slice), slice)

	// Printing 2D matrix
	fmt.Printf("3x%d: %v\n", ComputeRow(slice, 3), Make2D(slice, 3))
	fmt.Printf("4x%d: %v\n", ComputeRow(slice, 4), Make2D(slice, 4))
	fmt.Printf("5x%d: %v\n", ComputeRow(slice, 5), Make2D(slice, 5))
	fmt.Printf("6x%d: %v\n", ComputeRow(slice, 6), Make2D(slice, 6))

}

func Print2D(matrix [][]int) {
	fmt.Printf("1x%d: %v\n", len(matrix), matrix)
}

func ComputeRow(matrix []int, column int) int {
	row := len(matrix) / column
	col := len(matrix) % column
	// fmt.Println("row -> col\n", row, col)
	if col > 0 {
		row++
	}
	return row
}

func Make2D(matrix []int, column int) [][]int {
	mMatrix := make([][]int, ComputeRow(matrix, column))
	//	j := column
	//	k := 0

	for i, val := range matrix {
		row := i / column
		col := i % column
		//fmt.Printf("i:%d -> row:%d => col:%d\n", i, row, col)
		// we allocate when column is starting (i.e. col is a mod of column)
		if col == 0 {
			mMatrix[row] = make([]int, column)
		}
		mMatrix[row][col] = val
	}
	return mMatrix
}

func UniqueInts(matrix [][]int) []int {
	slice := make([]int, 0, len(matrix)+len(matrix[0]))
	slMap := make(map[int]int)
	for _, innerMatrix := range matrix {
		for _, val := range innerMatrix {
			if _, found := slMap[val]; !found {
				slice = append(slice, val)
				slMap[val]++
			}
		}
	}
	sort.Ints(slice)
	return slice
}

// Minimum result length is len(matrix) + len(matrix[0]); by using append()
// we can cope with irregular matrices whose inner slices are of different
// lengths.
func Flatten(matrix [][]int) []int {
	slice := make([]int, 0, len(matrix)+len(matrix[0]))
	for _, innerMatrix := range matrix {
		for _, value := range innerMatrix {
			slice = append(slice, value)
		}
	}
	return slice
}
