package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("inputs/06.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines = []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err = scanner.Err(); err != nil {
		panic(err)
	}

	result, err := solutionD6P2(lines)
	if err != nil {
		panic(err)
	}

	fmt.Printf("--> %d\n", result)
}
