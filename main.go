package main

import "fmt"

func main() {
	fmt.Println("hello world")

	nums := []int{1, 2, 3}
	sum := 0

	// iterate through slice and grab values only
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	// iterate throug the slice and get indexes
	for i, num := range nums {
		if num == 2 {
			fmt.Println("index", i)
		}
	}

	// itarate in  map and grab only values
	mymap := map[string]string{"first": "mercedes", "second": "BMW"}
	for _, car := range mymap {
		fmt.Println(car)
	}

	// iterate in map and grab only key
	for key := range mymap {
		fmt.Println(key)
	}

	// iterate through the strings
	for index, letterUnicode := range "some strubg" {
		fmt.Println(index, letterUnicode)
	}
}
