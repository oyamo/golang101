package main

import "fmt"


func average(values *[]int) (int, float64) {
	var sum = float64(0)
	for _, value := range *values {
		sum += float64(value)
	}

	return int(sum) , sum / float64(len(*values))
}

func main() {
	Array := []int{1,2,3,4,5,6,5,6,777}
	fmt.Println(average(&Array))
}
