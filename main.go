// ============================
// structs, methods, interfaces
// ============================

package main

import "fmt"

// structs
type person struct {
	name string
	age  int64
	age1 int64
	age2 int64
	age3 int64
}

// func getPerson(name string) *person {
// 	var newperson = person{name: name}

// 	return &newperson
// }

//interface
type geometry interface {
	area() float64
	perim() float64
}

// methods
type rect struct {
	width, height float64
}

type triangle struct {
	height, bottom, left, right float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (t triangle) area() float64 {
	return t.height * t.bottom / 2
}

func (t triangle) perim() float64 {
	return t.left + t.right + t.bottom
}

func measure(g geometry) {
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main() {
	// fmt.Println(getPerson("Jhon"))
	// s := &person{name: "David", age: 21}
	// fmt.Println(s.name)

	// per := s
	// fmt.Println(per.age)
	// fmt.Println(getPerson("Mike"))
	// d := person{name: "Ann", age: 40, age1: 1, age2: 2, age3: 3}

	// fmt.Println(&d)

	r := rect{width: 10, height: 5}
	t := triangle{height: 4, left: 4, right: 5, bottom: 3}

	measure(r)
	measure(t)
	// r.area()
	// r.perim()
	// fmt.Println("area: ", r.area())
	// fmt.Println("perim:", r.perim())
	// fmt.Println(r)

	// fmt.Println("area: ", rp.area())
	// fmt.Println("perim:", rp.perim())

}
