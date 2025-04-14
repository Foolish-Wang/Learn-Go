package main

import "fmt"

func main() {
	age := 10
	name := "John"
	rightHanded := true

	fmt.Println("Age:", age)
	fmt.Println("Name:", name)
	fmt.Println("Right Handed:", rightHanded)

	ageInTenYears := age + 10
	fmt.Println("Age in 10 years:", ageInTenYears)
	isTeenager := age >= 13 && age <= 19
	fmt.Println("Is Teenager:", isTeenager)
	
}