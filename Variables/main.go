package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
)

const prompt = " and don't type your number in, just press enter when ready."
func main() {

	//one way -declare, then assign
	var firstNumber int
	firstNumber = rand.Intn(8) + 2 //random number between 2 and 8

	//another way - declare and assign in one line
	var secondNumber = rand.Intn(8) + 2 //random number between 2 and 8

	//one step variable:declare and assign value and let go figure out the type
	subtraction := rand.Intn(8) + 2 //random number between 2 and 8
	answer := (firstNumber * secondNumber) - subtraction

	playTheGame(firstNumber, secondNumber, subtraction, answer)

	
}

func playTheGame(firstNumber,secondNumber,subtraction,answer int)  {
	reader := bufio.NewReader(os.Stdin)

	//display a welcome message
	fmt.Println("Welcome to the Guess Number game!")
	fmt.Println("------------------------------------------------------")
	fmt.Println("Think of a number between 1 and 10.", prompt)
	reader.ReadString('\n')
	fmt.Println("------------------------------------------------------")
	//take them through the game
	fmt.Print("Multipy your number by ", firstNumber, prompt)
	reader.ReadString('\n')

	fmt.Print("Now multipy the result by ", secondNumber, prompt)
	reader.ReadString('\n')

	fmt.Print("Now divide the result by the number you originally thought of",  prompt)
	reader.ReadString('\n')

	fmt.Print("Now subtract ", subtraction, prompt)
	reader.ReadString('\n')

	//give them the answer
	fmt.Println("------------------------------------------------------")
	fmt.Println("The answer is: ", answer)
}