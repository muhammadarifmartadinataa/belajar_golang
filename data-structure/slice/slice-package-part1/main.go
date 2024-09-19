package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func main() {

	numbers := []int{5, 6, 7, 8, 9}

	//binary search
	index, isFound := slices.BinarySearch[[]int](numbers, 9)
	fmt.Println("Found at", index)
	fmt.Println("is Found ? ", isFound)

	type Player struct {
		Name  string
		Score int
	}

	players := []Player{
		{"Arif", 100},
		{"Budi", 80},
		{"Doni", 400},
	}
	target := Player{"Budi", 100}

	idx, isFoundd := slices.BinarySearchFunc[[]Player](players, target, func(p1, p2 Player) int {
		return cmp.Compare[string](p1.Name, p2.Name)
	})

	fmt.Println("found at", isFoundd)
	fmt.Println("Index ke ", idx)

	cars := []string{
		"Subaru WrX",
		"Pajero Sport",
		"Pajero Dakar",
	}
	newSlice := slices.Clone[[]string](cars)
	fmt.Println(newSlice)

	//compact yaitu digunakan untuk menghapus duplikat dengan syrat datanya harus terurut teerlebih dahulu
	number := []int{
		1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5, 5, 7,
	}
	result := slices.Compact[[]int](number)
	fmt.Println(result)

	//compactfunction
	items := []string{"burger", "BURGER", "Kebab", "kebab", "Sate", "SATE"}
	result2 := slices.CompactFunc[[]string](items, func(item1, item2 string) bool {
		return strings.EqualFold(item1, item2)
	})
	fmt.Println(result2)

}
