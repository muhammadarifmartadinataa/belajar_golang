package main

import (
	"fmt"
	"maps"
)

func main() {

	//clone
	courses := map[string]string{
		"COO1": "Algorithm",
		"COO2": "Data Science",
		"COO3": "Computer Network",
		"COO4": "Design Grafis",
	}
	cloned := map[string]string(courses)
	fmt.Println(cloned)

	//delete
	students := map[string]int{
		"Alex": 90,
		"Arif": 100,
		"Devi": 95,
		"Glo":  85,
		"Dika": 70,
	}
	maps.DeleteFunc[map[string]int](students, func(name string, score int) bool {
		return score <= 80
	})

	fmt.Println(students)
}
