package main

import (
	"fmt"
	"reflect"
)

func main() {
	var primes [5]int
	var countries [5]string

	//package reflect untuk mengecek type dari valuenya
	fmt.Println(reflect.ValueOf(primes).Kind())
	fmt.Println(reflect.ValueOf(countries).Kind())

	//deklarasikan array
	x := [5]int{1, 2, 9, 10, 11}
	var y [5]int = [5]int{0, 2, 34}

	//looping array cara pertama
	for indexX := 0; indexX < len(x); indexX++ {
		fmt.Print(x[indexX])
	}
	fmt.Println(" ")

	//looping array cara kedua
	for indexY, element := range y {
		fmt.Println(indexY, "=> ", element)
	}

	//looping array cara ketiga
	index := 0
	for range x {
		fmt.Println(x[index])
		index++
	}

}
