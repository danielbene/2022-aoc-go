package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
	"strings"
)

var input []string

func read(filePath string) {
	file, err := os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	utils.CheckError(scanner.Err())
}

func getPriority(item rune) int {
	asciiVal := int(item)
	if asciiVal > 90 {
		return asciiVal - 96
	} else {
		return asciiVal - 64 + 26
	}
}

func getIntersection(strs ...string) rune {
	for _, r := range strs[0] {
		if strings.ContainsRune(strs[1], r) {
			if len(strs) == 3 {
				if !strings.ContainsRune(strs[2], r) {
					continue
				}

				return r
			}

			return r
		}
	}

	panic("No intersecting item.")
}

func part1() int {
	utils.StartTimer()

	sum := 0
	for _, line := range input {
		halfPoint := len(line) / 2
		firstHalf := line[:halfPoint]
		secondHalf := line[halfPoint:]

		sameItem := getIntersection(firstHalf, secondHalf)
		sum += getPriority(sameItem)
	}

	return sum
}

func part2() int {
	utils.StartTimer()

	sum := 0
	for i := 0; i < len(input); i += 3 {
		sameItem := getIntersection(input[i], input[i+1], input[i+2])
		sum += getPriority(sameItem)
	}

	return sum
}

func main() {
	utils.InitDay()

	read(utils.GetInputPath())

	utils.WriteSolutionInt(part1())
	utils.WriteSolutionInt(part2())

	fmt.Println("done.")
}
