package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	// fmt.Println("Hello, World!")
	// //Difference between Print and Println:
	// // Print does not add a newline at the end, while Println does.
	// fmt.Print("This is a simple Go program.")
	// fmt.Print("You can run this program using the command: go run main.go")



	// var whatTosay string = "Hello, World again!"
	// // whatTosay := "Hello, World again!" 
	// // // This is a shorthand for declaring and initializing a variable in Go.

	// sayHelloWorld(whatTosay)


	reader := bufio.NewReader(os.Stdin)
	whatTosay := doctor.Intro()
	fmt.Println(whatTosay)

	for {
		fmt.Print("You: ")
		// Read user input until a newline character is encountered
		userInput, _ := reader.ReadString('\n')
		
		userInput = strings.Replace(userInput, "\r\n", "", -1) // Remove the newline character from the input
		userInput = strings.Replace(userInput, "\n", "", -1) // Remove the newline character from the input

		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
		
	}
	
}

// func sayHelloWorld(whatTosay string) {
// 	fmt.Println(whatTosay)
// }


