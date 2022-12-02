package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

/* Type capturing Round data*/
type Round struct {
	opponentShape OpponentShape
	playerShape   PlayerShape
}

/* Enum Mapping Opponent Encrypted Name to Paper Rock Scissors*/
type OpponentShape string

const (
	P2Paper    OpponentShape = "B"
	P2Rock                   = "A"
	P2Scissors               = "C"
)

/* Enum Mapping Player Encrypted Name to Paper Rock Scissors*/

type PlayerShape string

const (
	P1Paper    PlayerShape = "Y"
	P1Rock                 = "X"
	P1Scissors             = "Z"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getShapeScore(playerShape PlayerShape) int {
	switch playerShape {
	case P1Paper:
		return 2
	case P1Rock:
		return 1
	case P1Scissors:
		return 3
	}
	return -1
}

func getOutcomePlayerScore(round Round) int {
	if round.playerShape == P1Paper {
		switch round.opponentShape {
		case P2Paper:
			return 3
		case P2Rock:
			return 6
		case P2Scissors:
			return 0
		}
	} else if round.playerShape == P1Rock {
		switch round.opponentShape {
		case P2Paper:
			return 0
		case P2Rock:
			return 3
		case P2Scissors:
			return 6
		}

	} else if round.playerShape == P1Scissors {
		switch round.opponentShape {
		case P2Paper:
			return 6
		case P2Rock:
			return 0
		case P2Scissors:
			return 3
		}

	}
	return -1
}

func TallyScore(rounds []Round) int {
	var totalScore int
	for _, round := range rounds {
		totalScore += getOutcomePlayerScore(round) + getShapeScore(round.playerShape)
	}
	return totalScore
}

/*
 * Parse file into clean Round object to work with
 */
func ParseInput(filename string) []Round {

	var rounds []Round

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
		var currentRound Round = Round{}

		/* Scan into temp variables */
		var tempOpponentPlay string
		var tempYourPlay string
		n, err := fmt.Sscanf(each_ln,
			"%s %s", &tempOpponentPlay, &tempYourPlay)
		if err != nil || n < 2 {
			panic(err)
		}

		var opponentShape OpponentShape
		var playerShape PlayerShape

		/* Convert to Enum */
		switch tempOpponentPlay {
		case P2Rock, P2Scissors, string(P2Paper):
			opponentShape = OpponentShape(tempOpponentPlay)
		}
		/* Convert to Enum */
		switch tempYourPlay {
		case P1Rock, P1Scissors, string(P1Paper):
			playerShape = PlayerShape(tempYourPlay)
		}
		currentRound.opponentShape = opponentShape
		currentRound.playerShape = playerShape

		rounds = append(rounds, currentRound)

		fmt.Printf("Opponent: %v, Player: %v\n", currentRound.opponentShape, currentRound.playerShape)

	}
	return rounds
}
