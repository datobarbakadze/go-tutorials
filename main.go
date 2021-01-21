package main

import (
	"errors"
	"fmt"
)

func someLoop(ret int) (int, error) {
	err := errors.New("error here")
	if ret == 2 {
		return ret, err
	} else if ret == 1 {
		return ret, nil
	}
	return ret, errors.New("None of them")
}

func main() {
	v, err := someLoop(1)
	fmt.Println(v)
	fmt.Println(err)
}
