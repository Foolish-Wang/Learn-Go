package main

import "fmt"

type Vehicle struct {
	NumberOfWheels     int
	NumberOfPassengers int
}

type Car struct {
	Make       string
	Model      string
	Year       int
	IsElectric bool
	IsHybrid   bool
	Vehicle    Vehicle
}

func (v Vehicle) ShowDetails() {
	fmt.Println("Number of wheels:", v.NumberOfWheels)
	fmt.Println("Number of passengers:", v.NumberOfPassengers)
}

func (c Car) Show() {
	fmt.Println("Make:", c.Make)
	fmt.Println("Model:", c.Model)
	fmt.Println("Year:", c.Year)
	fmt.Println("Is Electric:", c.IsElectric)
	fmt.Println("Is Hybrid:", c.IsHybrid)
	c.Vehicle.ShowDetails()
}

func main() {
	suv := Vehicle{
			NumberOfWheels:     4,
			NumberOfPassengers: 5,
	}

	volvoXC90 := Car{
		Make:       "Volvo",
		Model:      "XC90",
		Year:       2023,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle:    suv,
	}

	volvoXC90.Show()

	fmt.Println()

	teslaModelX := Car{
		Make:       "Tesla",
		Model:      "Model X",
		Year:       2023,
		IsElectric: true,
		IsHybrid:   false,
		Vehicle: Vehicle{
			NumberOfWheels:     4,
			NumberOfPassengers: 7,
		},
	}
	teslaModelX.Show()
	
}