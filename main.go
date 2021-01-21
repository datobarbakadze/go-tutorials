package main

import (
	"fmt"
	"time"
)

// simple routines
func simpleLoop() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
}

// channels
func worker(finish chan<- bool) {
	fmt.Println("Workinbg ....")
	time.Sleep(2 * time.Second)
	fmt.Println("done")

	finish <- true
}

func mainWorker(finish <-chan bool) {
	if <-finish == true {
		fmt.Println("hello world")
	} else {
		fmt.Println("not finished")
	}
}

// func mainWorker(finish chan)
func main() {

	go simpleLoop()

	go func() {
		fmt.Println("hello world")
	}()
	time.Sleep(time.Second)

	messages := make(chan string)
	go func() { messages <- "ping" }()
	nsg := <-messages

	fmt.Println(nsg)

	bufferedChannel := make(chan string, 2)

	bufferedChannel <- "hi there"
	bufferedChannel <- "hi there2"

	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)

	finishControlChannel := make(chan bool)
	go worker(finishControlChannel)
	go mainWorker(finishControlChannel)

	time.Sleep(5 * time.Second)

}
