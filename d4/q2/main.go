package main

import (
	"berkin/advent-of-code/utils"
	_ "embed"
	"fmt"
	"strconv"
	"strings"
	"time"
)

var (
	//go:embed resources/input.txt
	resourcesData []byte
)

type pair struct {
	startPoint int
	endPoint   int
}

func main() {
	start := time.Now()
	solution(strings.Split(string(resourcesData), "\n"))
	elapsed := time.Since(start)
	fmt.Printf("%s Execution took %s", " --", elapsed)
	//utils.BuildAndRunPreviousTasks()
}

func (p pair) isOverlaps(otherPair pair) bool {
	if p.startPoint > otherPair.endPoint || p.endPoint < otherPair.startPoint {
		return false
	}
	return true
}

func convertStringToPair(input string) pair {
	s := strings.Split(input, "-")
	startPoint, err := strconv.Atoi(s[0])
	utils.Check(err)

	finishPoint, err := strconv.Atoi(s[1])
	utils.Check(err)

	return pair{startPoint, finishPoint}
}

func solution(input []string) {
	result := 0

	for i := 0; i < len(input); i++ {
		pairs := strings.Split(input[i], ",")
		firstPair, secondPair := convertStringToPair(pairs[0]), convertStringToPair(pairs[1])

		if firstPair.isOverlaps(secondPair) {
			result++
		}
	}

	fmt.Printf("Output: %d", result)
}
