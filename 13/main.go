package main

import (
	"io/ioutil"
	"strings"
)

func main() {

}

func readFromFile(fileName string) []string {
	raws, err := ioutil.ReadFile(fileName)
	if err != nil {
		return []string{}
	}

	return strings.Split(string(raws), "\n")
}
