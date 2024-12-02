package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/salmonllama/Advent-of-Code-2024/lib"
)

func main() {
	raw := lib.ReadFile("/day02/input.txt")

	lines := splitLines(raw)

	var reports [][]int

	for _, item := range lines {
		if item != "" {
			reports = append(reports, parseIntSlice(item))
		}
	}

	example := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	var safeReports int
	for _, report := range example {
		if isSafe(report) { safeReports++ }
	}

	fmt.Println(safeReports)
}

func splitLines(input string) []string {
	return strings.Split(input, "\n")
}

func parseIntSlice(input string) []int {
	var i []int
	s := strings.Split(input, " ")

	for _, item := range s {
		val, err := strconv.Atoi(item)
		if err != nil { panic(err) }

		i = append(i, val)
	}

	return i
}

func isSafe(target []int) bool {
	if (isAsc(target) || isDesc(target)) {
		if isAcceptableDiff(target) {
			fmt.Printf("%d: Safe\n", target)
			return true
		}
	}

	return false
}

func isAsc(target []int) bool {
	asc, ascIndex := isAscending(target)
	if asc {
		// fmt.Printf("%d: Safe - Ascending\n", target)
		return true
	} else {
		ascRetry := remove(target, ascIndex)

		secondAsc, _ := isAscending(ascRetry)

		if secondAsc { 
			// fmt.Printf("%d: Safe (modified) - Ascending\n", ascRetry)
			return true
		}
	}
	fmt.Printf("%d: Unsafe - Not ascending, can't be modified\n", target)
	return false
}

func isDesc(target []int) bool {
	desc, descIndex := isDescending(target)
	if desc {
		// fmt.Printf("%d: Safe - Descending\n", target)
		return true
	} else {
		descRetry := remove(target, descIndex)

		secondDesc, _ := isDescending(descRetry)

		if secondDesc { 
			// fmt.Printf("%d: Safe (modified) - Descending\n", descRetry)
			return true
		}
	}
	fmt.Printf("%d: Unsafe - Not descending, can't be modified\n", target)
	return false
}

func remove(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func isAscending(target []int) (bool, int) {
	lastValue := 0
	currentValue := 0
	ascending := true

	for ascending {
		for i, item := range target {
			currentValue = item

			// if i == 0 {
			// 	lastValue = currentValue
			// 	continue
			// }

			if currentValue < lastValue {
				return false, i
			}

			lastValue = currentValue
		}
		return true, 0
	}
	return true, 0
}

func isDescending(target []int) (bool, int) {
	lastValue := math.MaxInt
	currentValue := 0
	descending := true

	for descending {
		for i, item := range target {
			currentValue = item

			// if i == 0 {
			// 	lastValue = currentValue
			// 	continue
			// }
			
			if currentValue > lastValue {
				return false, i
			}

			lastValue = currentValue
		}
		return true, 0
	}
	return true, 0
}

func isAcceptableDiff(target []int) bool {
	current := 0
	last := 0

	for i, item := range target {
		current = item

		if i == 0 { 
			last = current
			continue 
		}
		
		if current == last {
			fmt.Printf("%d: Unsafe - Two equal values\n", target)
			return false
		} else if (current - last) > 3 {
			fmt.Printf("%d: Unsafe - Too high difference\n", target)
			return false
		} else if (last - current) > 3 {
			fmt.Printf("%d: Unsafe - Too high difference\n", target)
			return false
		}

		last = current
	}
	return true
}