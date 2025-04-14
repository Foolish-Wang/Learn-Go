package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/eiannone/keyboard"
)

//define a struct to hold the user information
type User struct {
	UserName string
	Age int
	FavoriteNumber float64
	OwnsDog bool
}
var reader *bufio.Reader

func main() {
	reader = bufio.NewReader(os.Stdin)
	var user User
	user.UserName = readString("Enter your name:")
	user.Age = readInt("Enter your age:")
	user.FavoriteNumber = readFloat("Enter your favorite number:")
	user.OwnsDog = readBool("Do you own a dog? (y/n):")
	fmt.Printf("Your name is: %s and your age is: %d and your favorite number is: %.2f. Owns a dog: %t.\n", 
	user.UserName, 
	user.Age,
	user.FavoriteNumber,
	user.OwnsDog)

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

func readFloat(s string) float64 {
	for {
		fmt.Println(s)
		prompt()
		userInput, _ := reader.ReadString('\n')

		userInput = strings.TrimSpace(userInput)

		num, err := strconv.ParseFloat(userInput, 64) //convert the string to an integer
		if err != nil {
			fmt.Println("Please enter a whole number")
		} else{
			return num
		}
	}
}

func readBool(s string) bool {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err) //log the error and exit the program
	}
	defer func(){
		_ = keyboard.Close()
	}()
	for {
		fmt.Println(s)
		prompt()
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err) //log the error and exit the program
		}
		if char == 'y' || char == 'Y' {
			return true
		} else if char == 'n' || char == 'N' {
			return false
		} else {
			fmt.Println("Please enter a valid input")
		}
	}

}

		
		