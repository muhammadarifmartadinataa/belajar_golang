package main

import "fmt"

func main() {
	fmt.Println(powOptimize(2, 8)) //8
	fmt.Println(powOptimize(5, 3))
	fmt.Println(powOptimize(10, 2))
	fmt.Println(powOptimize(2, 5))
	fmt.Println(powOptimize(7, 3))
}

//OLog(n)
func powOptimize(x, n int) int {
	var result int = 1
	for n > 0 {
		if n%2 == 1 {
			result *= x
		}
		x = x * x
		n = n / 2
	}
	return result * x
}

//O(N)
func pow(x, n int) int {
	var result int = 1
	for i := 1; i <= n; i++ {
		result *= x
	}
	return result
}
