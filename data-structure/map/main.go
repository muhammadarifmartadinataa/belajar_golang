// map memiliki key dan value yang mana setiap keynya bersifat uniq
package main

import "fmt"

func main() {
	//deklarasi map
	harga := map[string]int{"siomay": 10000, "batagor": 15000}
	fmt.Println(harga)

	//deklarasi menggunakan make
	var price = make(map[string]int)
	price["bakwan"] = 7000
	fmt.Print(price)
}
