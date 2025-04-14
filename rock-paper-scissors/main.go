package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

const (
	ROCK     = 0
	PAPER    = 1
	SCISSORS = 2
)

func main() {
	playerChoice := ""
	playerValue := -1

	computerValue := rand.Intn(3)

	reader := bufio.NewReader(os.Stdin)

	clearScreen()

	fmt.Print("Please enter rock, paper, or scissors ->")
	playerChoice, _ = reader.ReadString('\n')
	playerChoice = strings.TrimSpace(playerChoice)

	switch playerChoice {
	case "rock":
		playerValue = ROCK
	case "paper":
		playerValue = PAPER

	case "scissors":
		playerValue = SCISSORS

	default:
		fmt.Println("Invalid choice")
	}

	switch computerValue {
	case ROCK:
		fmt.Println("Computer chose rock")
		break
	case PAPER:
		fmt.Println("Computer chose paper")
		break
	case SCISSORS:
		fmt.Println("Computer chose scissors")
		break
	default:
		fmt.Println("Invalid computer choice")
	}
	
	if playerValue == computerValue {
		fmt.Println("It's a tie!")
	}
	if playerValue == ROCK && computerValue == SCISSORS {
		fmt.Println("Rock beats scissors! You win!")
	}
	if playerValue == SCISSORS && computerValue == PAPER {
		fmt.Println("Scissors beats paper! You win!")
	}
	if playerValue == PAPER && computerValue == ROCK {
		fmt.Println("Paper beats rock! You win!")
	}
	if playerValue == SCISSORS && computerValue == ROCK {
		fmt.Println("Rock beats scissors! You lose!")
	}
	if playerValue == PAPER && computerValue == SCISSORS {
		fmt.Println("Scissors beats paper! You lose!")
	}
	if playerValue == ROCK && computerValue == PAPER {
		fmt.Println("Paper beats rock! You lose!")
	}

}

// clearScreen clears the screen
func clearScreen() {
	if strings.Contains(runtime.GOOS, "windows") {
		// windows
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		// linux or mac
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
