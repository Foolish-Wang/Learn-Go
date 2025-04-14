package main

import (
	"fmt"
	"sort"
)

//basic types (numeric types, string, boolean, etc.)
// var myInt int = 10
// var myUint uint = 20
// var myInt8 int8 = -128
// var myUint8 uint8 = 255
// var myInt16 int16 = -32768
// var myUint16 uint16 = 65535
// var myInt32 int32 = -2147483648
// var myUint32 uint32 = 4294967295
// var myInt64 int64 = -9223372036854775808
// var myUint64 uint64 = 18446744073709551615

// var myString string = "Hello, World!" //string is unchangeable
// var myBool bool = true

// var myByte byte = 'A'
// var myFloat float64 = 20.5
// var myComplex complex128 = 1 + 2i

// aggregation types (array, slice, map, struct, etc.)

// var myArray [5]int = [5]int{1, 2, 3, 4, 5}
// var mySlice []int = []int{1, 2, 3, 4, 5}
// var myMap map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}

type Car struct {
	Make  string
	Model string
	Year  int
	luxury bool
}

func (a *Car) String() {
	fmt.Printf("Car: %s %s %d", a.Make, a.Model, a.Year)
}

func main() {

	mycar := Car{
		Make:  "Toyota",
		Model: "Corolla",
		Year:  2020,	
		luxury: false,
	}
	fmt.Println(mycar)
	mycar.String() // prints the car details

	//Pointers
	x := 10
	myFirstPointer := &x // myFirstPointer is a pointer to x
	fmt.Println("Value of x:", x) // prints 10
	fmt.Println("Address of x:", &x) // prints the address of x
	fmt.Println("Value of myFirstPointer:", myFirstPointer) // prints the address of x

	changeValueOfPointer(myFirstPointer) // pass the pointer to the function
	fmt.Println("Value of x after changeValueOfPointer:", x) // prints 20


	//slices
	var animals []string // declare a slice of strings
	animals = append(animals, "dog") 
	animals = append(animals, "cat")
	animals = append(animals, "bird")	
	animals = append(animals, "fish")
	animals = append(animals, "lizard")
	animals = append(animals, "snake")
	animals = append(animals, "turtle")
	animals = append(animals, "frog")
	animals = append(animals, "hamster")

	fmt.Println("Animals slice:", animals) // prints the slice of animals

	for _, x := range animals {
		fmt.Println(x) // prints each animal in the slice
		// _ is the index, x is the value
	}

	fmt.Println("First two animals are", animals[0:2]) 
	fmt.Println(len(animals)) // prints the length of the slice
	fmt.Println(cap(animals)) // prints the capacity of the slice

	fmt.Println(sort.StringsAreSorted(animals)) // prints false, because the slice is not sorted

	sort.Strings((animals)) // sorts the slice of animals
	fmt.Println("Sorted animals slice:", animals) // prints the sorted slice of animals
	fmt.Println(sort.StringsAreSorted(animals)) // prints true, because the slice is sorted

	DeleteFromSlice(animals, 2) // delete the animal at index 2
	fmt.Println("Animals slice after deletion:", animals) // prints the slice of animals after deletion


	intMap := make(map[string]int) 
	intMap["one"] = 1 // add a key-value pair to the map
	intMap["two"] = 2 // add another key-value pair to the map
	intMap["three"] = 3 // add another key-value pair to the map
	intMap["four"] = 4 // add another key-value pair to the map

	//random access to map
	for key, value := range intMap {
		fmt.Println("Key:", key, "Value:", value) // prints each key-value pair in the map
	}
	delete(intMap, "two") // delete the key "two" from the map

	for key, value := range intMap {
		fmt.Println("Afrer deletion:") 
		fmt.Println("Key:", key, "Value:", value) // prints each key-value pair in the map
	}

	// check if a key exists in the map
	// if the key exists, ok will be true, and el will be the value of the key
	el, ok := intMap["one"] // check if the key "one" exists in the map
	if ok {
		fmt.Println("Key 'one' exists with value:", el) // prints the value of the key "one"
	} else {
		fmt.Println("Key 'one' does not exist")
	}

	fmt.Println(sumMany(1,1,1,2,2,3,4,5,6,7,8,9,10)) // call the variadic function with multiple arguments
}

func changeValueOfPointer(p *int) {
	*p = 20 // change the value at the address pointed to by p
}

func DeleteFromSlice(slice []string, index int) []string {
	slice[index] = slice[len(slice)-1] // replace the element at index with the last element
	slice = slice[:len(slice)-1] // remove the last element from the slice
	return slice // return the modified slice
}

func sumMany(nums ...int) int { // variadic function that takes a variable number of int arguments
	sum := 0
	for _, num := range nums {
		sum += num // add each number to the sum
	}
	return sum // return the sum
}