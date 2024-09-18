package main

import (
	"fmt"
	"math"
)

//func tanpa menggunakan paramater

func sayHello() {
	fmt.Println("say hello")
}

// function menggunakan paramater
func greeting(hour int) {
	if hour < 12 {
		fmt.Println("Selamat Siang")
	} else if hour < 18 {
		fmt.Print("Selamat Sore")
	} else {
		fmt.Println("Selamat Malam")
	}
}

// function  single return
func calculateSquare(side int) int {
	return side * side
}

// function multiple return
func calculateCircle(diameter float64) (float64, float64) {
	var keliling = math.Pi * math.Pow(diameter/2, 2)
	var luas = math.Pi * diameter
	return keliling, luas
}

func main() {
	hour := 10
	sayHello()
	greeting(hour)

	var side = 15
	wide := calculateSquare(side)
	fmt.Printf("Luas Persegi :  %d \n\n", wide)

	var diameter float64 = 15
	keliling, luas := calculateCircle(diameter)

	fmt.Printf("luas Lingkaran: %.2f \n", keliling)
	fmt.Printf("Keliling lingkaran : %.2f \n\n", luas)

}
