package main

import (
	"P1/parser"
	"fmt"
)

func main() {

	rounds := parser.ParseInput("input.txt")
	totalScore := parser.TallyScore(rounds)
	fmt.Printf("Total Score: %v\n", totalScore)
	return
}
