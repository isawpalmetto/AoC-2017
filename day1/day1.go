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
		next int
		sum  int
	)

	for i, _ := range input {
		val := getInt(input[i])

		if i == len(input)-1 {
			next = getInt(input[0])
		} else {
			next = getInt(input[i+1])
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
