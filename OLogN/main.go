package main

import "fmt"

func main() {
	n := 100
	result := 0
	for i := n; i > 0; {
		result += i
		i = i / 2
		fmt.Println(i)
	}
}
