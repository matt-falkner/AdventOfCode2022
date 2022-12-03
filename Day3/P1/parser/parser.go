package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/* Type capturing Round data*/
type Rucksack struct {
	first  string
	second string
}

func f(in rune) int {
	return int(in - 'A' + 1)
}

func getScore(duplicatedCharacter string) int {
	upper := strings.ToUpper(duplicatedCharacter)
	upperRune := rune(upper[0])
	indexInAlphabet := f(upperRune)

	if strings.ToUpper(duplicatedCharacter) == duplicatedCharacter {
		dupeScore := (indexInAlphabet + 26)
		fmt.Printf("[IsUpper][%v]: %v\n\n", duplicatedCharacter, dupeScore)
		return dupeScore
	} else {
		fmt.Printf("[IsLower][%v]: %v\n\n", duplicatedCharacter, indexInAlphabet)
		return (indexInAlphabet)
	}
}

func FindDuplicatesCountScore(rucksacks []Rucksack) int {
	totalScore := 0
	for i, rucksack := range rucksacks {
		fmt.Printf("rucksack: (%v||%v) %d\n", rucksack.first, rucksack.second, i)

		for i := 0; i < len(rucksack.first); {
			isDuplicated := strings.Contains(rucksack.second, string(rucksack.first[i]))
			if isDuplicated {
				duplicatedChar := string(rucksack.first[i])
				dupeScore := getScore(duplicatedChar)
				totalScore += dupeScore
				break //Ignore duplicates
			}
			i++
		}

	}

	fmt.Printf("Total Score: %v\n", totalScore)
	return totalScore
}

/*
 * Parse file into clean Round object to work with
 */
func ParseRucksacksFromInput(filename string) []Rucksack {

	var rucksacks []Rucksack

	/* Does file exist*/
	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("failed to open")

	}
	/* Scan in file int a string array*/
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	file.Close()

	/* Iterate over each line*/
	for _, each_ln := range text {
		var currentRucksack Rucksack = Rucksack{}

		currentRucksack.first = each_ln[0:(len(each_ln) / 2)]
		currentRucksack.second = each_ln[(len(each_ln) / 2):(len(each_ln))]
		rucksacks = append(rucksacks, currentRucksack)
	}
	return rucksacks
}
