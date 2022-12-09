package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
	"strconv"
)

var input []int

func read(filePath string) {
	file, err := os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	elfCal := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			input = append(input, elfCal)
			elfCal = 0
		} else {
			cal, _ := strconv.Atoi(line)
			elfCal += cal
		}
	}

	input = append(input, elfCal)
	utils.CheckError(scanner.Err())
}

func part1() int {
	utils.StartTimer()

	// part1 solution

	return 0
}

func part2() int {
	utils.StartTimer()

	// part2 solution

	return 0
}

func main() {
	utils.InitDay()

	read(utils.GetInputPath())

	utils.WriteSolutionInt(part1())
	utils.WriteSolutionInt(part2())

	fmt.Println("done.")
}
