package main

import "fmt"

func main() {
	fmt.Println("Hi there")

	// types and variable declarations
	fmt.Println("1+1 = ", 7.0/3.0, " 2+2 = ", 4)

	var a int = 3
	var b, c int = 3.0, 4

	var t, g float32 = 3.5, 3.9
	var k string
	fmt.Println(a)
	fmt.Println(b, c)
	fmt.Println(t, g)
	fmt.Println(k)

	f := "apple"

	fmt.Println(f)

	// for loop
	for i := 0; i < 20; i++ {

		if i == 10 {
			fmt.Println("skipped")
			continue
		}
		fmt.Println("dato")

		switch i {
		case 3:
			fmt.Println("We were here")
			break
		case 7:
			fmt.Println("again?")
			break

		}
	}

	// arrays

	ar := [5]int{5, 4, 3, 2, 1}
	i := 0
	for i <= 4 {
		ar[i] = i
		i++
	}

	fmt.Println(ar)

	// slices
	sl := make([]string, 3)
	sl[0] = "a"
	sl[1] = "b"
	sl[2] = "c"

	sl2 := make([]string, len(sl))
	fmt.Println(sl2)
	copy(sl2, sl)
	fmt.Println(sl2)
	sl2 = append(sl2, "e")
	fmt.Print(sl2)

	sliced := sl2[1:3]
	fmt.Println(sliced)
	sliced2 := sl2[:3]
	fmt.Println(sliced2)

	// maps
	map1 := make(map[string]int)
	map1["hello"] = 6
	map1["number"] = 3

	fmt.Println(map1)
	delete(map1, "hello")
	fmt.Println(map1)

	value, _ := map1["number"]
	fmt.Println(value)
}
