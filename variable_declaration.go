package main

import (
	"fmt"
)


func main() {
	// Declare the variables
	var name, school, occupation string
	var age, weight, salary uint32

	// Assign the variables with values
	age = 20
	weight = 70
	school = "Orero High School"
	occupation = "Engineering"
	salary = 340000
	name = "Oyamo Brian"

	// Print the variables to the user
	fmt.Printf("Hello, my name if %10s\n", name)
	fmt.Printf("I learn at %10s\n", school)
	fmt.Printf("My occupation is %10s\n", occupation)
	fmt.Printf("My age is %10d\n", age)
	fmt.Printf("My weight is %10d\n", weight)
	fmt.Printf("I earn %10d\n", salary)

}
