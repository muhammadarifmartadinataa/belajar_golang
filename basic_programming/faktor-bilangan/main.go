package main

import "fmt"

func main() {
	nilai := 100
	if nilai < 0 || nilai > 100 {
		fmt.Print("Invalid")
	} else if nilai <= 39 {
		fmt.Print("E")
	} else if nilai <= 50 {
		fmt.Print("D")
	} else if nilai <= 69 {
		fmt.Print("C")
	} else if nilai <= 84 {
		fmt.Print("B")
	} else if nilai <= 100 {
		fmt.Print("A")
	}
}
