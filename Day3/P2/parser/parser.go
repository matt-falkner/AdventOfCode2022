package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/* Type capturing Round data*/
type Group struct {
	first  string
	second string
	third  string
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
		//fmt.Printf("[IsUpper][%v]: %v\n\n", duplicatedCharacter, dupeScore)
		return dupeScore
	} else {
		//fmt.Printf("[IsLower][%v]: %v\n\n", duplicatedCharacter, indexInAlphabet)
		return (indexInAlphabet)
	}
}

func FindDuplicatesCountScore(groups []Group) int {
	totalScore := 0
	for i, rucksack := range groups {
		fmt.Printf("rucksack: (%v||%v) %d\n", rucksack.first, rucksack.second, i)

		for i := 0; i < len(rucksack.first); {
			isDuplicatedSecond := strings.Contains(rucksack.second, string(rucksack.first[i]))
			if isDuplicatedSecond {
				isDuplicatedThird := strings.Contains(rucksack.third, string(rucksack.first[i]))
				if isDuplicatedThird {
					duplicatedChar := string(rucksack.first[i])
					dupeScore := getScore(duplicatedChar)
					fmt.Printf("[Found in all 3 in group: %v][score %v]: \n\n", duplicatedChar, dupeScore)

					totalScore += dupeScore
					break //Ignore duplicates
				}
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
func ParseRucksacksFromInput(filename string) []Group {

	var groups []Group

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
	var group_count = 1
	var currentGroup Group = Group{}

	/* Iterate over each line*/
	for _, each_ln := range text {

		if group_count == 1 {
			currentGroup.first = each_ln
			group_count = 2
		} else if group_count == 2 {
			currentGroup.second = each_ln
			group_count = 3
		} else if group_count == 3 {
			currentGroup.third = each_ln
			groups = append(groups, currentGroup)
			//reset
			group_count = 1
			currentGroup = Group{}
		}
	}
	return groups
}
