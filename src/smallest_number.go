package main

import "fmt"

func main() {
	var numbers = [...]int{
		34,67,888,
		34,67,89,
		21,45,78,0x09,
		34,12,45,78,
		0xf,90, 90,
		12,67,89,
	}
	fmt.Println("Before",numbers)
	// Find the smallest number if in that array without using any sort function
	for i := 0; i < len(numbers); i++ {
		for j := 0; j < len(numbers); j++ {
			if numbers[i] < numbers[j] {
				numbers[i] = numbers[i] + numbers[j]
				numbers[j] = numbers[i] - numbers[j]
				numbers[i] = numbers[i] - numbers[j]
			}
		}
	}

	fmt.Println("After", numbers)
	fmt.Println("Smallest Number", numbers[0])

}
