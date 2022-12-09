package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
)

var input []int

func read(filePath string) {
	file, err := os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		score := int(str[2]) - int(str[0])
		shapeBonus := int(str[2]) - 87 // ascii [88-90]

		input = append(input, parse(score, shapeBonus))
	}

	utils.CheckError(scanner.Err())
}

func parse(score int, bonus int) int {
	if score == 23 {
		return 3 + bonus
	}

	// ascii duckery
	switch score + bonus {
	case 22, 26, 27:
		return 6 + bonus
	case 23, 24, 28:
		return 0 + bonus
	default:
		panic("Invalid state.")
	}
}

func part1() int {
	utils.StartTimer()

	totalScore := 0
	for _, score := range input {
		totalScore += score
	}

	return totalScore
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
