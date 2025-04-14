package main

import "fmt"
type Employee struct {
	Name string
	Age  int
	Salary int
	FullTime bool
}

func main() {
	apple := 18
	banana := 9

	//only same type can be compared
	fmt.Println(apple == banana) // false
	fmt.Println(apple != banana) // true
	fmt.Println(apple > banana)  // true

	jack := Employee{
		Name:    "Jack White",
		Age:     30,
		Salary:  50000,
		FullTime: true,
	}

	joe := Employee{
		Name:    "Joe Black",
		Age:     25,
		Salary:  40000,
		FullTime: false,
	}

	var employees []Employee
	employees = append(employees, jack, joe)
	for _, x := range employees {
		if x.FullTime {
			fmt.Println("Full time employee:", x.Name)
		} else {
			fmt.Println("Part time employee:", x.Name)
		}

		if x.Age > 30 || x.Salary > 50000 {
			fmt.Println("Senior employee:", x.Name)
		} else {
			fmt.Println("Junior employee:", x.Name)
		}
	}
}