package main

import (
	"day1/parser"
	"fmt"
)

func main() {
	Elfs := parser.ParseElfLedger("elf_calories_ledger.txt")
	fmt.Printf("\n\n\n%v\n", Elfs)
	elfWithMostCalories := parser.MostCaloriesElf(Elfs)
	fmt.Printf("\n\n\n%v\n", elfWithMostCalories)
}
