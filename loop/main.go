package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"myapp/mylogger"
	"os"
	"time"
)

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}

	j := 1000

	for j > 100 {
		j = rand.Intn(1000) + 1
		fmt.Println(j)

	}

	fmt.Println("Got", j, "Loop finished")

	reader := bufio.NewReader(os.Stdin)

	ch := make(chan string)
	go mylogger.ListenForLog(ch)
	fmt.Println("Enter a message:")

	for i := 0; i < 5; i++ {
		fmt.Println("->")
		input, _ := reader.ReadString('\n')
		ch <- input
		time.Sleep(time.Second)
	}

}