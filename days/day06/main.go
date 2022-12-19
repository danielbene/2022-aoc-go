package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
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

func part1() int {
	utils.StartTimer()

	fmt.Println(input)

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
