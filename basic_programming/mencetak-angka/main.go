package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%4 == 0 && i%7 == 0 {
			fmt.Print("Buzz")
		} else if i%4 == 0 {
			fmt.Print(i * i)
		} else if i%7 == 0 {
			fmt.Print(i * i * i)
		} else {
			fmt.Print(i)
		}
	}

}
