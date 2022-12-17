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

func part1() string {
	utils.StartTimer()

	// part1 solution

	return "asd"
}

func part2() string {
	utils.StartTimer()

	// part2 solution

	return "asd"
}

func main() {
	utils.InitDay()

	read(utils.GetInputPath())

	utils.WriteSolutionStr(part1())
	utils.WriteSolutionStr(part2())

	fmt.Println("done.")
}
