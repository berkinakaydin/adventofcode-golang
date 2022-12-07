package utils

import (
	"os"
)

func ReadResources(path string) string {
	content, err := os.ReadFile(path)

	Check(err)

	return string(content)
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
