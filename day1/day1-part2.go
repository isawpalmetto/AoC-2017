package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// takes arg 1 as input
	if len(os.Args) != 2 {
		fmt.Println("Must provide input")
		os.Exit(0)
	}
	input := os.Args[1]
	var (
		sum  int
		next int
	)

	steps := len(input) / 2

	for i, _ := range input {
		val := getInt(input[i])

		// next = i + steps
		if i+steps > len(input)-1 {
			next = getInt(input[i+steps-len(input)])
		} else {
			next = getInt(input[i+steps])
		}

		if val == next {
			sum += val
		}
	}
	fmt.Printf("sum = %+v\n", sum)
}

func getInt(b byte) int {
	val, _ := strconv.Atoi(string(b))
	return val
}
