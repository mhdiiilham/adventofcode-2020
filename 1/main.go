package main

import "fmt"

func main() {
	input := readInputFile("input.txt")
	result := reportRepair(*input)

	fmt.Println("Result: ", result)

}
