package main

import (
	"fmt"
)

func main() {
	defer func() {
		if recover() != nil {
			fmt.Println("Tangkap Eror")
		}
		fmt.Println("Hello World")
	}()
	fmt.Println("Altera")
	panic("eror")
	fmt.Println("Academy")
}
