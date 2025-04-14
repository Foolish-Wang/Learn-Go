package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

var keyPressChan chan rune

func main() {
	// goroutine
	keyPressChan = make(chan rune) // create a channel to listen for key presses
	go listenForKeyPress()         // start a goroutine to listen for key presses

	fmt.Println("Press any key, or q to exit...")
	_ = keyboard.Open() // open the keyboard for input
	defer func(){
		keyboard.Close() // close the keyboard when done
	}()

	for {
		char,_,_ := keyboard.GetSingleKey() // get a single key press
		if char == 'q' || char == 'Q'{ 
			break
		}

		keyPressChan <- char // send the key press to the channel
	}
}

func listenForKeyPress() {
	for {
		key := <-keyPressChan // receive a key press from the channel
		fmt.Printf("Key pressed: %c\n", key) // print the key press
	}
}