package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
	"strings"
)

var input string

func read(filePath string) {
	file, err := os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = scanner.Text()
	}

	utils.CheckError(scanner.Err())
}

func scanMessage(markerPos int) int {
outer:
	for i := markerPos; i <= len(input); i++ {
		subStr := input[i-markerPos : i]
		for _, r := range subStr {
			if strings.Count(subStr, string(r)) > 1 {
				continue outer
			}
		}

		return i
	}

	panic("No valid substring found.")
}

func part1() int {
	utils.StartTimer()

	return scanMessage(4)
}

func part2() int {
	utils.StartTimer()

	return scanMessage(14)
}

func main() {
	utils.InitDay()

	read(utils.GetInputPath())

	utils.WriteSolutionInt(part1())
	utils.WriteSolutionInt(part2())

	fmt.Println("done.")
}
