package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/salmonllama/Advent-of-Code-2024/lib"
)

func main() {
	raw := lib.ReadFile("/day01/input.txt")

	lines := splitLines(raw)

	left, right := populate(lines)

	slices.Sort(left)
	slices.Sort(right)

	solution := process(left, right)

	fmt.Println(solution)

	/** Part 2 logic **/

	var similarities []int
	for _, item := range left {
		n := countOcurrences(item, right)

		similarities = append(similarities, (item * n))
	}

	sum := 0
	for _, item := range similarities {
		sum += item
	}
	fmt.Println(sum)
}

func splitLines(input string) []string {
	return strings.Split(input, "\n")
}

func populate(input []string) (left []int, right []int) {
	for _, item := range input {
		if item != "" {
			row := strings.Split(item, "   ")

			leftInt, err := strconv.Atoi(row[0])
			if err != nil { panic(err) }
			
			rightInt, err := strconv.Atoi(row[1])
			if err != nil { panic(err) }

			left = append(left, leftInt)
			right = append(right, rightInt)
		}
	}

	return left, right
}

func process(oLeft []int, oRight []int) int {
	sum := 0
	var diffs []int

	

	for i := 0; i < len(oLeft); i++ {
		var s int
		if (oLeft[i] > oRight[i]) {
			s = oLeft[i] - oRight[i]
		} else {
			s = oRight[i] - oLeft[i]
		}

		diffs = append(diffs, s)
	}

	for _, item := range diffs {
		sum += item
	}

	// for i, item := range oLeft {

	// 	sum += (oRight[i] - item)
	// }

	return sum
}

func countOcurrences(target int, list []int) int {
	occurences := 0

	for _, item := range list {
		if item == target {
			occurences++
		}
	}
	
	return occurences
}