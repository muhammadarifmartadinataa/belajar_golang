package main

import "fmt"

func main() {
	//fibonanci (5) = fibo (4) + fibo(3)
	fmt.Println(fibo(5))
}
func fibo(n int) int {
	if n <= 1 {
		return n
	}
	return fibo(n-1) + fibo(n-2)
}
