package main

import (
	"fmt"
	"time"
)

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		chan1 <- "Hi I am first one"
	}()

	go func() {
		time.Sleep(2 * time.Second)
		chan2 <- "Hi I am second one"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			fmt.Println(msg1)
		case msg2 := <-chan2:
			fmt.Println(msg2)
		}
	}
}
