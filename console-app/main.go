package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/eiannone/keyboard"
)

func main() {
	err := keyboard.Open()
	if err != nil {
		log.Fatal(err) //log the error and exit the program
	}

	//defer the closing of the keyboard to ensure it is closed when the program exits
	//anymous function 
	defer func(){
		_ = keyboard.Close()
	}()
	
	coffees := make(map[int]string)
	coffees[1] = "Cappuccino"
	coffees[2] = "Latte"
	coffees[3] = "Espresso"
	coffees[4] = "Americano"
	coffees[5] = "Mocha"
	coffees[6] = "Macchiato"


	fmt.Println("MENU")
	fmt.Println("----------------------------------------")
	fmt.Println("1 - Cappuccino")
	fmt.Println("2 - Latte")
	fmt.Println("3 - Espresso")
	fmt.Println("4 - Americano")
	fmt.Println("5 - Mocha")
	fmt.Println("6 - Macchiato")
	fmt.Println("Q - Quit")
	fmt.Println("----------------------------------------")

	for {
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			log.Fatal(err) //log the error and exit the program
		}

		if char == 'q' || char == 'Q' {
			fmt.Println("Exiting program...")	
			break
		}

		i,_ := strconv.Atoi(string(char)) //convert the character to an integer
		fmt.Printf("You chose %s\n", coffees[i])//format the string to include the coffee name

		
	
	}

}