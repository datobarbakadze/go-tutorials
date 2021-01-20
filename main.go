package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func singleType(a, b, d int) int {
	return a + b + d
}

func main() {
	result := plus(1, 2)
	result2 := singleType(1, 2, 3)

	fmt.Println(result, result2)
}
