package main

import (
	"fmt"
	"slices"
)

func main() {
	//package slice index untuk mecari index

	fruits := []string{"banana", "mango", "Orange"}
	idx := slices.Index[[]string](fruits, "Orange")

	fmt.Println("found at", idx)

	numbers := []int{12, 14, 15, 20}
	//find the index of number that divisible by 5
	result := slices.IndexFunc[[]int](numbers, func(number int) bool {

		return number%5 == 0
	})
	fmt.Println("found at : ", result)

	//package slice insert the index (keunggulan insert dari append yaitu kita bisa menyisipkan nilai dimana saja berbdea dengan append)
	//insert value at index 1
	numbers = slices.Insert[[]int](numbers, 1, 80)

	//insert value at index 2
	numbers = slices.Insert[[]int](numbers, 2, 90, 80, 60)

	fmt.Println(numbers)

	//package slice max and min
	maxVal := slices.Max[[]int](numbers)
	minVal := slices.Min[[]int](numbers)

	fmt.Println(maxVal)
	fmt.Println(minVal)
}
