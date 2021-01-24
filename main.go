package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c1 <- "Hi I am C1"
	}()

	select {
	case msg := <-c1:
		fmt.Println(msg)
	// hah returns the current time via channel :d
	case coolestTimeout := <-time.After(time.Second):
		fmt.Println("time here", coolestTimeout)
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "Legendary channel C2"
	}()
	select {
	case msg := <-c2:
		fmt.Println(msg)
	case <-time.After(3 * time.Second):
		fmt.Println("timout here son of a freak")
	}
}
