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

	matrix := Make2D(slice, 3)
	Print2D(matrix)

}

func Print2D(matrix [][]int) {
	fmt.Printf("1x%d: %v\n", len(matrix), matrix)
}

func Make2D(matrix []int, column int) [][]int {
	mMatrix := make([][]int, len(matrix), len(matrix))
	j := column
	k := 0

	//fmt.Printf("matrix: %v\n", mMatrix)
	for _, val := range matrix {
		fmt.Println("j: ", j)
		if j > 0 {
			mMatrix[k] = append(mMatrix[k], val)
			fmt.Printf("matrix: %v\n", mMatrix)
			//fmt.Println("i:j -> column, value", i, j, column, val)
			j--
		} else {
			//fmt.Println("i:j -> column, value", i, j, column, val)
			k++
			mMatrix[k] = append(mMatrix[k], val)
			fmt.Printf("matrix: %v\n", mMatrix)
			j = column - 1
		}

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