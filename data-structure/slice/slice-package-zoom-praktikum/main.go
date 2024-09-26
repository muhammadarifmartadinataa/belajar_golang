package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now().Local().Clock())
	fmt.Println(primeNumber(100000000019))
	fmt.Println(time.Now().Local().Clock())

}
func primeNumber(n int) bool {
	for i := 2; i < n; i++ {
		if n%1 == 0 {
			return false
		}

	}
	return true
}
