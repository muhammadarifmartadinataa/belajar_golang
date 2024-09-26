package main

import (
	"fmt"
	"slices"
)

func main() {
	//function contains adalah sebuah function yang bisa kita gunkana untuk mengecek apakah sebuah item di slice ada atau tidak

	item := []string{"apple", "milk", "coffe", "banana"}

	current := "milk"

	isFound := slices.Contains[[]string](item, current)

	fmt.Println(isFound)

	//function slice  delete

	persons := []string{"arif", "faiz", "doni", "tirta"}

	result := slices.Delete(persons, 0, 1)
	fmt.Println(result)

	//DeleteFunc Slice

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	hasil := slices.DeleteFunc[[]int](numbers, func(number int) bool {
		return number%2 == 0
	})
	fmt.Println(hasil)

	//function
}
