package main

import (
	"fmt"
)

// functions with multiple return values
func retMultVals() (int, int) {
	return 1, 4
}

// variadic functions | the functions which take indefinity arity
func variadic(numbers ...int) {
	fmt.Println(numbers, " ")
	var total int
	for _, num := range numbers {
		total += num
		fmt.Println(num)
	}
	fmt.Println(total)
}

// clouses
func intSeq() func() int {
	total := 1 + 1
	return func() int {
		total++
		return total
	}
}

// recursions
// calc factuorial with recursion
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

// pointers
func normalAssign(val int) {
	val = 0
}

func ponterAssign(val *int) {
	*val = 0
}

func main() {

	// // mult rets
	// val1, val2 := retMultVals()
	// fmt.Println(val1, val2)

	// // variadics
	// variadic(1, 2, 3, 4, 5, 6, 7)

	// numsForVariadic := []int{7, 6, 5, 4, 3, 2}
	// variadic(numsForVariadic...)

	// // closures
	// nextInt := intSeq()
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())
	// fmt.Println(nextInt())

	// newint := intSeq()
	// fmt.Println(newint())
	// fmt.Println(newint())
	// fmt.Println(newint())

	// // recursions
	// // factorial
	// fmt.Println(factorial(3))

	// pointers
	// assignable := 12
	// normalAssign(assignable)
	// fmt.Println("no pinter: ", assignable)

	// ponterAssign(&assignable)
	// fmt.Println("with pointer ", assignable)

}
