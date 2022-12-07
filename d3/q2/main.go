package main

import (
	"berkin/advent-of-code/utils"
	_ "embed"
	"errors"
	"fmt"
	"strings"
	"time"
	"unicode"
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
	//utils.BuildAndRunPreviousTasks()
}

func findCommonLetter(groups []string) (rune, error) {
	for _, l := range groups[0] {
		var isExistsInAllGroup bool = true

		for _, group := range groups[1:] {
			if !strings.ContainsRune(group, l) {
				isExistsInAllGroup = false
				break
			}
		}

		if isExistsInAllGroup {
			return l, nil
		}
	}

	return rune(0), errors.New("no common letter")
}

func calculateLetterScore(letter rune) int {
	var asciiOffset int

	if unicode.IsLower(letter) {
		asciiOffset = 96
	} else {
		asciiOffset = 38
	}

	return (int(letter) - asciiOffset)
}

func solution(input []string) {
	result := 0

	const GROUP_SIZE = 3

	for i := 0; i < len(input); i = i + GROUP_SIZE {
		group := input[i : i+GROUP_SIZE]

		commonChar, err := findCommonLetter(group)
		utils.Check(err)

		result += calculateLetterScore(commonChar)
	}

	fmt.Printf("Output: %d", result)
}
