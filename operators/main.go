package main

import "math"

func main() {

	answer := 7 + 3*2 - 1/5
	println(answer)

	answer2 := 7 + 3*2 - 1%5// % is the modulus operator
	println(answer2)

	answer3 := 7^2 + 3*2 - 1%5// ^ is the bitwise XOR operator
	// ^ is not the power operator in Go, it is the bitwise XOR operator
	println(answer3)

	answer4 := math.Pow(7, 2) + 3*2 - 1%5
	println(answer4)

	area := math.Pi * 5 * 5
	println(area)

	answer5 := 7 + 3*2 - 1/5.0
	println(answer5)

	second := 31
	minute := 1

	if minute < 59 && second + 1 > 59 {
		minute++
		second = 0
	}
	println(minute, second)
}