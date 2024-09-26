package main

import "fmt"

func main() {
	//O(n)
	var arr = []int{1, 3, 2, 4, 5}
	fmt.Println(arr)
	var reverseArr []int
	for i := len(arr) - 1; i >= 0; i-- {
		reverseArr = append(reverseArr, arr[i])
	}
	fmt.Println(reverseArr)
}
