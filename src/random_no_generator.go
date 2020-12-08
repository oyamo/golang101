package main

import (
	"fmt"
	"math/rand"
)

func main() {
	const MAX = 1e8
	var Numbers [MAX]int

	for i := 0; i < MAX; i+=20 {
		Numbers[i] = rand.Intn(1000)
		Numbers[i+1] = rand.Intn(1000)
		Numbers[i+2] = rand.Intn(1000)
		Numbers[i+3] = rand.Intn(1000)
		Numbers[i+4] = rand.Intn(1000)
		Numbers[i+5] = rand.Intn(1000)
		Numbers[i+6] = rand.Intn(1000)
		Numbers[i+7] = rand.Intn(1000)
		Numbers[i+8] = rand.Intn(1000)
		Numbers[i+9] = rand.Intn(1000)
		Numbers[i+10] = rand.Intn(1000)
		Numbers[i+11] = rand.Intn(1000)
		Numbers[i+12] = rand.Intn(1000)
		Numbers[i+13] = rand.Intn(1000)
		Numbers[i+14] = rand.Intn(1000)
		Numbers[i+15] = rand.Intn(1000)
		Numbers[i+16] = rand.Intn(1000)
		Numbers[i+17] = rand.Intn(1000)
		Numbers[i+18] = rand.Intn(1000)
		Numbers[i+19] = rand.Intn(1000)

	}

	fmt.Println(Numbers[:20])
}

