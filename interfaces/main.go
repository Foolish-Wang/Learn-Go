package main

import "fmt"

//interface type
type Animal interface {
	Says() string
	HowManyLegs() int
}

type Dog struct {
	name         string
	Sound        string
	NumberOfLegs int
}

func (d *Dog) Says() string {
	return d.Sound
}

func (d *Dog) HowManyLegs() int {
	return d.NumberOfLegs
}

type Cat struct {
	name         string
	Sound        string
	NumberOfLegs int
	HasTail      bool
}

func (d *Cat) Says() string {
	return d.Sound
}

func (d *Cat) HowManyLegs() int {
	return d.NumberOfLegs
}

func main() {
	dog := Dog{
		name:         "Dog",
		Sound:        "Woof",
		NumberOfLegs: 4,
	}
	cat := Cat{
		name:         "Cat",
		Sound:        "Meow",
		NumberOfLegs: 4,
		HasTail:      true,
	}

	Riddle(&dog)
	Riddle(&cat)
	
}

func Riddle(a Animal){
	riddle := fmt.Sprintf(`This animal says "%s" and has %d legs`, a.Says(), a.HowManyLegs())
	fmt.Println(riddle)
}