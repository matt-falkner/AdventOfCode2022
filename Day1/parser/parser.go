package parser

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"unicode"
)

type Elf struct {
	index         int
	totalCalories int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isInt(s string) bool {
	for _, c := range s {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func MostCaloriesElf(Elfs []Elf) Elf {
	sort.Slice(Elfs[:], func(i, j int) bool {
		return Elfs[i].totalCalories > Elfs[j].totalCalories
	})

	return Elfs[0]
}

func CaloriesOfTop3Elfs(Elfs []Elf) int {
	sort.Slice(Elfs[:], func(i, j int) bool {
		return Elfs[i].totalCalories > Elfs[j].totalCalories
	})

	var calories int = Elfs[0].totalCalories + Elfs[1].totalCalories + Elfs[2].totalCalories

	return calories
}

func ParseElfLedger(filename string) []Elf {
	var Elfs []Elf

	/* Read in File */
	dat, err := os.ReadFile(filename)
	check(err)
	var lineNumber int = 0
	var currElfIndex int = 1

	fmt.Print(len(dat))
	msg := make([]byte, 0)
	var currentElf Elf = Elf{}

	for i := 0; i < len(dat); {

		/* If the character is an Integer type, save it to data buffer to help build whole number of line*/
		if isInt(string(dat[i])) {
			msg = append(msg, dat[i])
		}

		/* End of an Line containing either a number or the end of an elf*/
		if dat[i] == '\n' {
			lineNumber += 1

			/* Reset Found Number */

			/* Add the current number saved in message since the line is now over
			to the current elf, once the elf its done, the total calories will be appended with his object */
			if string(msg) != "" {
				// string to int
				i, err := strconv.Atoi(string(msg))
				if err != nil {
					// ... handle error
					panic(err)
				}

				//fmt.Printf("Found=%v\n", i)

				currentElf.totalCalories += i
			}
			var next_char_index = i + 1
			//fmt.Printf("%v >= %v\n", next_char_index, len(dat))
			if next_char_index >= len(dat) {
				/* Looking beyond the end of the file, don't look!*/
				//fmt.Println("Looking Beyond edge of file SAVE ME")
			} else {
				/* Next character should exist*/

				if dat[next_char_index] == '\n' {
					//fmt.Printf("Space after Space found")

					currentElf.index = currElfIndex
					Elfs = append(Elfs, currentElf)

					/* Clear Curr Temp Elf for next elf*/
					currentElf = Elf{}
					currElfIndex += 1
				}
			}
			/* RESET MSG: Number captured has been consumed. Reset it before loading it again....*/
			msg = make([]byte, 0)
		}
		/* End of the Last Elf*/
		if i+1 == len(dat) {
			//fmt.Println("End of Last Elf")

			if string(msg) != "" {
				// string to int
				i, err := strconv.Atoi(string(msg))
				if err != nil {
					// its not a INT, WHY?!
					panic(err)
				}

				fmt.Printf("Found=%v\n", i)

				currentElf.totalCalories += i

				currentElf.index = currElfIndex
				Elfs = append(Elfs, currentElf)

			}

		}

		i++
	}

	fmt.Println("Returning Elfs")
	return Elfs
}
