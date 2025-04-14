package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
var reader *bufio.Reader

func main() {


	reader = bufio.NewReader(os.Stdin)
	userName := readString("Enter your name:")
	age := readInt("Enter your age:")
	fmt.Printf("Your name is: %s and your age is: %d\n", userName, age)


}

func prompt() {
	fmt.Print("-> ")
}

func readString(s string) string {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		if userInput == "" {
			fmt.Println("Please enter a valid input")
		} else {
				return userInput
		}
		
	}
	
}

func readInt(s string) int {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')

		userInput = strings.TrimSpace(userInput)

		num, err := strconv.Atoi(userInput) //convert the string to an integer
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else{
			return num
		}
	}
}

		