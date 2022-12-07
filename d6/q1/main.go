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

func isDuplicated(input []string) bool {
	allKeys := make(map[string]bool)

	for _, item := range input {
		if _, ok := allKeys[item]; !ok {
			allKeys[item] = true
		}
	}

	return len(allKeys) != 4
}

func solution(input string) {
	const SEQUENCE_LENGTH = 4
	var dataStream []string

	for i, letter := range input {

		if len(dataStream) == SEQUENCE_LENGTH && i > 0 {
			if isDuplicated(dataStream) {
				dataStream = dataStream[1:]
			} else {
				fmt.Printf("Output: %d", i)
				return
			}
		}

		dataStream = append(dataStream, string(letter))
	}

}
