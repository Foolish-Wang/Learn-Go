package main

import (
	"log"
	"myapp/staff"
)
var employees = []staff.Employee{
	{FirstName: "John", LastName: "Doe", Salary: 50000, FullTime: true},
	{FirstName: "Jane", LastName: "Smith", Salary: 60000, FullTime: false},
	{FirstName: "Alice", LastName: "Johnson", Salary: 70000, FullTime: true},
	{FirstName: "Bob", LastName: "Brown", Salary: 80000, FullTime: false},
}
func main() {
	mystaff := staff.Office{AllStaff: employees}

	log.Println(mystaff.All())

	log.Println(mystaff.Overpaid())
	
	log.Println(mystaff.Underpaid())

	// mystaff.notVisible() 
	// This will cause a compilation error because notVisible is not exported

}

