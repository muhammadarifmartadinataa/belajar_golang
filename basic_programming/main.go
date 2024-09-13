package main

import "fmt"

func main() {

	/*
	 *****
	 */
	for i := 0; i <= 4; i++ {
		for j := 1; j <= 5; j++ { // Diperbaiki '=' yang salah
			//fmt.Print(" ")

			if j >= 5-i {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println() // Tambahkan newline agar setiap baris tidak disambung
	}
}
