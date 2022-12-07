package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	//go:embed resources/input.txt
	resourcesData []byte
)

type Stack []string

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	} else {
		index := len(*s) - 1   // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index]      // Remove it from the stack by slicing it off.
		return element, true
	}
}

func (s *Stack) Peek() string {
	return (*s)[len(*s)-1]
}

func main() {
	start := time.Now()
	solution(strings.Split(string(resourcesData), "\n"))
	elapsed := time.Since(start)
	fmt.Printf("%s Execution took %s", " --", elapsed)
	//utils.BuildAndRunPreviousTasks()
}

func (s *Stack) buildStacks(input string) {
	s.Push(input)
}

func solution(input []string) {
	blankRow := 0
	for {
		if input[blankRow] == "" {
			break
		}
		blankRow++
	}

	var stacks []Stack = make([]Stack, blankRow)

	for i := blankRow - 2; i >= 0; i-- {
		stackIndex := 0
		dataIndex := 1
		row := input[i]

		for i, container := range row {
			if i == dataIndex {
				if string(container) != " " {
					stacks[stackIndex].buildStacks(string(container))
				}

				stackIndex++
				dataIndex += 4
			}
		}
	}

	r, _ := regexp.Compile(`[0-9]+`)

	for i := blankRow + 1; i < len(input); i++ {
		row := input[i]
		s := r.FindAllString(row, 3)

		times, _ := strconv.Atoi(s[0])
		from, _ := strconv.Atoi(s[1])
		to, _ := strconv.Atoi(s[2])

		tempStack := Stack{}

		for i := 0; i < times; i++ {
			data, _ := stacks[from-1].Pop()
			tempStack.Push(data)
		}

		for i := 0; i < times; i++ {
			data, _ := tempStack.Pop()
			stacks[to-1].Push(data)
		}
	}

	for _, stack := range stacks {
		fmt.Print(stack.Peek())
	}
}
