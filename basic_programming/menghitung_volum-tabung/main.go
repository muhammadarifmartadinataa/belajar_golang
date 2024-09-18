package main

import "fmt"

func main() {
	var r float64 = 14
	var t float64 = 20
	const phi float64 = 3.14

	hitungVolum := phi * r * r * t
	fmt.Print("Hasil Hitung Volume Tabung adalah : ", hitungVolum)

}
