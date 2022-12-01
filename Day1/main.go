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

	top3CaloriesTotal := parser.CaloriesOfTop3Elfs(Elfs)
	fmt.Printf("\nCalories Of Top3: %v\n", top3CaloriesTotal)

}
