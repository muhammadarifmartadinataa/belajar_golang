package main

import "fmt"

func main() {
	fmt.Println(isPrime(13))
}
func isPrime(n int64) int32 {
	for i := 2; i < int(n); i++ {
		if int(n)%i == 0 {
			return int32(i)
		}
	}
	return 1
}
