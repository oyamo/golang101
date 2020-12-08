package main
import "fmt"

func main() {
	Marks := make(map[string]int)
	Marks["Chemistry"] 		= 10
	Marks["Mathematics"] 	= 60
	Marks["Biology"] 		= 56
	Marks["Physics"] 		= 90
	Marks["ComputerScience"]= 90


	for subject, Mark := range Marks {
		fmt.Println(subject, Mark)
	}

	// Cheking if a value exists
	if Mark, Available := Marks["Maths"];  !Available{
		fmt.Println("It Does not exist")
	} else {
		fmt.Println("He scored",Mark,"in Maths")
	}

	// Accessing map values
	fmt.Println(Marks["Biology"])
 }
