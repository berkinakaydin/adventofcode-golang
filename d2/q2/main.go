package main

import (
	_ "embed"
	"fmt"
	"strings"
	"time"
)

var (
	//go:embed resources/input.txt
	resourcesData []byte
)

func main() {
	start := time.Now()
	solution(strings.Split(string(resourcesData), "\n"))
	elapsed := time.Since(start)
	fmt.Printf("%s Execution took %s", " --", elapsed)
}

func winningHands(yourHand string, opponentHand string) int {
	var score, hand int

	switch opponentHand {
	case "A":
		if yourHand == "X" {
			hand = 3
			score = 0 + hand
		}
		if yourHand == "Y" {
			hand = 1
			score = 3 + hand
		}
		if yourHand == "Z" {
			hand = 2
			score = 6 + hand
		}
	case "B":
		if yourHand == "X" {
			hand = 1
			score = 0 + hand
		}
		if yourHand == "Y" {
			hand = 2
			score = 3 + hand
		}
		if yourHand == "Z" {
			hand = 3
			score = 6 + hand
		}
	case "C":
		if yourHand == "X" {
			hand = 2
			score = 0 + hand
		}
		if yourHand == "Y" {
			hand = 3
			score = 3 + hand
		}
		if yourHand == "Z" {
			hand = 1
			score = 6 + hand
		}
	}

	return score
}

func solution(input []string) {
	result := 0

	for _, row := range input {
		s := strings.Split(row, " ")
		opponentHand := s[0]
		yourHand := s[1]
		result += winningHands(yourHand, opponentHand)
	}

	fmt.Printf("Output: %d", result)
}
