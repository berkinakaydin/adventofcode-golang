package main

import (
	"berkin/advent-of-code/utils"
	_ "embed"
	"errors"
	"fmt"
	"strings"
	"time"
)

var (
	//go:embed resources/input.txt
	resourcesData []byte
)

type hand struct {
	code  string
	value int
}

const (
	WIN  = 6
	DRAW = 3
	LOSE = 0
)

const (
	ROCK     = "A"
	PAPER    = "B"
	SCISSORS = "C"
)

func main() {
	start := time.Now()
	solution(strings.Split(string(resourcesData), "\n"))
	elapsed := time.Since(start)
	fmt.Printf("%s Execution took %s", " --", elapsed)
}

func assignHand(cardType string) (hand, error) {
	switch cardType {
	case "A", "X":
		return hand{ROCK, 1}, nil
	case "B", "Y":
		return hand{PAPER, 2}, nil
	case "C", "Z":
		return hand{SCISSORS, 3}, nil
	default:
		return hand{}, errors.New("unknown card type")
	}
}

func calculateScore(yourHand hand, opponentHand hand) int {
	result := yourHand.value

	switch yourHand.code {
	case ROCK:
		switch opponentHand.code {
		case ROCK:
			result += DRAW

		case PAPER:
			result += LOSE

		case SCISSORS:
			result += WIN
		}

	case PAPER:
		switch opponentHand.code {
		case ROCK:
			result += WIN

		case PAPER:
			result += DRAW

		case SCISSORS:
			result += LOSE
		}

	case SCISSORS:
		switch opponentHand.code {
		case ROCK:
			result += LOSE

		case PAPER:
			result += WIN

		case SCISSORS:
			result += DRAW
		}
	}

	return result
}

func solution(input []string) {
	result := 0

	for _, row := range input {
		s := strings.Split(row, " ")
		yourHand, err := assignHand(s[1])
		utils.Check(err)

		opponentHand, err := assignHand(s[0])
		utils.Check(err)

		result += calculateScore(yourHand, opponentHand)
	}

	fmt.Printf("Output: %d", result)
}
