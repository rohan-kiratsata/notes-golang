package main

import (
	"fmt"
	"sort"
)

func main() {

	var testArr [5]string

	testArr[0] = "a"
	testArr[1] = "b"
	// testArr[2] = "c"
	testArr[3] = "d"
	testArr[4] = "e"

	fmt.Println("arr:", testArr)
	fmt.Println("length:", len(testArr))

	var testArr2 = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr2:", testArr2)
	fmt.Println("length:", len(testArr2))

	// array is value type not reference type

	var testArr3 = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("arr3:", testArr3)
	fmt.Println("length:", len(testArr3))

	// 2D array
	var testArr4 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("arr4:", testArr4)
	fmt.Println("length:", len(testArr4))

	// SLICES

	var testSlice = []int{1, 2, 3, 4}
	fmt.Println("slice:", testSlice)

	// append - add to slice
	testSlice = append(testSlice, 5)
	fmt.Println("slice appended:", testSlice)

	// slice with range
	testSlice = append(testSlice[1:])
	fmt.Println("slice with range:", testSlice)

	// last range is non-inclusive
	testSlice = append(testSlice[1:3], 10)
	fmt.Println("slice with range and new value:", testSlice)

	highScore := make([]int, 5)
	highScore[0] = 100
	highScore[1] = 341
	highScore[2] = 870
	highScore[3] = 450
	highScore[4] = 422

	fmt.Println("highScore:", highScore)

	sort.Ints(highScore)
	fmt.Println("highScore sorted:", highScore)

	fmt.Println(sort.IntsAreSorted(highScore))

	var fruits = []string{"apple", "banana", "grape", "mango"}
	fmt.Println("fruits:", fruits)

	var index int = 2
	fruits = append(fruits[:index], fruits[index+1:]...)
	fmt.Println("fruits after append:", fruits)

}
