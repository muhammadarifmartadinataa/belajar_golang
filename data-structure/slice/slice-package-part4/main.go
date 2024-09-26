package main

import (
	"cmp"
	"fmt"
	"slices"
)

type Product struct {
	Name  string
	Price int
}

func main() {

	products := []Product{
		{"Banana", 12000},
		{"Banana", 12000},
		{"Coffee", 80000},
		{"Cereal", 40000},
	}
	//for duplikat
	slices.SortStableFunc[[]Product](products, func(p1, p2 Product) int {
		return cmp.Compare[int](p1.Price, p2.Price)
	})
	fmt.Println(products)

	items := []string{
		"apple", "coffe", "milk", "banana", "cereal",
	}

	//package slice replace (mengganti)
	items = slices.Replace[[]string](items, 0, 3, "durian", "rambutan", "wafer")
	fmt.Println("before reverse  : ", items)

	//package slice reverse yaitu bertujuan untuk membalik urutan dari sebuah slice
	slices.Reverse[[]string](items)

	fmt.Println("After Reverse : ", items)

	//package slice sort untuk mengurutkan
	numbers := []int{
		-90, 80, 29, 40, 33, 20,
	}
	slices.Sort[[]int](numbers)
	fmt.Println("After Sorted : ", numbers)

}
