package main

import (
	_ "embed"
	"fmt"
	"time"
)

var (
	//go:embed resources/input.txt
	resourcesData []byte
)

func main() {
	start := time.Now()
	solution((string(resourcesData)))
	elapsed := time.Since(start)
	fmt.Printf("%s Execution took %s", " --", elapsed)
	//utils.BuildAndRunPreviousTasks()
}

const SEQUENCE_LENGTH = 14

func isDuplicated(input []rune) bool {
	allKeys := make(map[rune]bool)

	for _, item := range input {
		if _, ok := allKeys[item]; !ok {
			allKeys[item] = true
		}
	}

	return len(allKeys) != SEQUENCE_LENGTH
}

func solution(input string) {
	var dataStream []rune

	for i, letter := range input {
		if len(dataStream) == SEQUENCE_LENGTH && i > 0 {
			if isDuplicated(dataStream) {
				dataStream = dataStream[1:]
			} else {
				fmt.Printf("Output: %d", i)
				return
			}
		}

		dataStream = append(dataStream, letter)
	}
}
