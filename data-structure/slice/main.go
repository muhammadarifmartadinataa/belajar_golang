package main

import "fmt"

func main() {

	//APPEND & COPPY
	var colors = []string{"red", "green", "yellow"}
	colors = append(colors, "purpple")

	copiedColors := make([]string, 10)
	copy(copiedColors, colors)

	fmt.Println(copiedColors)
}
