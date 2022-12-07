package utils

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
)

var colorCyan = "\033[36m"

func BuildAndRunPreviousTasks() {
	fmt.Println("=======PREVIOUS RESULTS=======")
	matches, err := filepath.Glob("../d*/q*/*.go")
	if err != nil {
		panic(err.Error())
	}

	r, _ := regexp.Compile(`[a-zA-Z]{1}\d`)

	for _, match := range matches {
		s := r.FindAllString(match, 2)
		date, question := s[0], s[1]

		err := exec.Command("go", "build", match).Run()
		Check(err)

		output, err := exec.Command("./main").Output()
		Check(err)
		fmt.Print(string(colorCyan))
		fmt.Printf("Date %s Question %s %s\n", strings.ToUpper(date), strings.ToUpper(question), string(output))

		exec.Command("rm", "main").Run()
	}
}
