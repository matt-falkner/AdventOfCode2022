package main

import (
	"P2/parser"
)

func main() {

	rucksacks := parser.ParseRucksacksFromInput("input.txt")
	parser.FindDuplicatesCountScore(rucksacks)

	return
}
