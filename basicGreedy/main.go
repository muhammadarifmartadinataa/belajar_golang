package main

import "fmt"

//O(N)
func main() {
	var coin = []int{1, 3, 4}
	target := 6
	i := len(coin) - 1
	result := []int{}
	for target > 0 || i >= 0 {
		if target >= coin[i] {
			target = target - coin[i]
			result = append(result, coin[i])
		} else {
			i--
		}
	}
	fmt.Println(result)

}
