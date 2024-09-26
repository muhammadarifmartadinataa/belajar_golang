package main

func main() {

}

func factorialNumber(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * factorialNumber(n-1)
	}

}
