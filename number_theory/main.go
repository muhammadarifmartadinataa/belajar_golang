package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(gcd(30, 12))
}

// chek bilangan prima
func checkPrime(number int) bool {
	if number < 2 {
		return false
	}
	sqrtNumber := int(math.Sqrt(float64(number)))
	for i := 2; i < sqrtNumber; i++ {
		if number%i == 0 {
			return false
		}
	}
	return true
}

// FPB Faktor Persekutuan Terbesar
func gcd(a, b int) int {
	if b == 0 {
		return a
	} else {
		return gcd(b, a%b)
	}
}
