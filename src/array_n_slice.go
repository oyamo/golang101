package main

import "fmt"

func main()  {
	// This is an array that we are not aware of its size
	var Names = [...]string{"Oyamo", "Brian", "Otieno", "Mustapha"}
	fmt.Println("My name is", Names[0], Names[1])

	// Using the []type method
	var mutableNames = make([]string,10)
	temp := append(mutableNames, "BrianLEE")
	fmt.Println(mutableNames, temp)


}