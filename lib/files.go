package lib

import (
	"os"
)

func ReadFile(filename string) string {
	path := "/workspaces/Advent-of-Code-2024"
	data, err := os.ReadFile(path + filename)

	if err != nil {
		panic(err)
	}

	return string(data)
}