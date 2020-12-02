package main

import "fmt"

func main() {
	input := readInputFile("input.txt")
	result := reportRepair(input)
	threeSumResult := reportRepair3Sum(input)

	fmt.Println("Result: ", result)
	fmt.Println("Three Sum Result: ", threeSumResult)

}
