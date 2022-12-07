package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
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

func solution(input []string) {
	min := 0
	max := math.MinInt

	for _, s := range input {

		if s == "" {
			if min > max {
				max = min
			}
			min = 0
			continue
		}

		intVar, _ := strconv.Atoi(s)
		min += intVar
	}

	fmt.Printf("Output: %d", max)
}
