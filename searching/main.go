package main

import (
	"fmt"
	"sort"
)

func main() {
	items := []int{4, 5, 32, 2, 1, 2, 3}
	fmt.Println(linearSearch(items, 100))

	elements := []int{1, 5, 4, 2, 8, 9, 0}
	sort.Ints(elements)
	target := 9
	index := sort.SearchInts(elements, target)
	fmt.Println(index)

}
func linearSearch(items []int, target int) int {
	for i := 0; i < len(items); i++ {
		if items[i] == target {
			return i
		}
	}
	return -1
}
