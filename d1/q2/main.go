package main

import (
	_ "embed"
	"fmt"
	"sort"
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
	var calories []int

	caloriesPerGroup := 0
	for _, s := range input {

		if s == "" {
			calories = append(calories, caloriesPerGroup)
			caloriesPerGroup = 0
			continue
		}

		intVar, _ := strconv.Atoi(s)
		caloriesPerGroup += intVar
	}

	sort.Sort(sort.Reverse(sort.IntSlice(calories)))

	output := 0
	for i := 0; i < 3; i++ {
		output += calories[i]
	}

	fmt.Printf("Output: %d", output)
}
