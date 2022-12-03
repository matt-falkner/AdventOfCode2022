package main

import (
	"P1/parser"
)

func main() {

	rucksacks := parser.ParseRucksacksFromInput("input.txt")
	parser.FindDuplicatesCountScore(rucksacks)

	return
}
