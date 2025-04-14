package main

import "time"

var chan1 = make(chan string)
var chan2 = make(chan string)

func task1() {
	time.Sleep(1 * time.Second)
	chan1 <- "Task 1 completed"
}

func task2() {
	time.Sleep(2 * time.Second)
	chan1 <- "Task 2 completed"
}

func main() {
	go task1()
	go task2()

	//select is used to wait for multiple channel operations
	// It will block until one of the channels is ready to send or receive data
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-chan1:
			println("received", msg1)
		case msg2 := <-chan2:
			println("received", msg2)
		}
	}

	println("All tasks completed")
}