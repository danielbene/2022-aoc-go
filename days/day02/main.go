package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
)

var col1 []int
var col2 []int

func read(filePath string) {
	file, err := os.Open(filePath)
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()
		col1 = append(col1, int(str[0]))
		col2 = append(col2, int(str[2]))
	}

	utils.CheckError(scanner.Err())
}

func part1() int {
	utils.StartTimer()

	totalScore := 0
	for i := 0; i < len(col1); i++ {
		score := col2[i] - col1[i]
		bonus := col2[i] - 87

		if score == 23 {
			totalScore += 3 + bonus
			continue
		}

		// ascii duckery
		switch score + bonus {
		case 22, 26, 27:
			totalScore += 6 + bonus
		case 23, 24, 28:
			totalScore += 0 + bonus
		default:
			panic("Invalid state.")
		}
	}

	return totalScore
}

func part2() int {
	utils.StartTimer()

	totalScore := 0
	for i := 0; i < len(col1); i++ {

		score := (col2[i] - 88) * 3
		tmpVal := col1[i] - 64
		bonus := 0

		if score == 0 {
			tmpVal -= 1
		} else if score == 6 {
			tmpVal += 1
		}

		if tmpVal == 0 {
			bonus = 3
		} else if tmpVal == 4 {
			bonus = 1
		} else {
			bonus = tmpVal
		}

		totalScore += score + bonus
	}

	return totalScore
}

func main() {
	utils.InitDay()

	read(utils.GetInputPath())

	utils.WriteSolutionInt(part1())
	utils.WriteSolutionInt(part2())

	fmt.Println("done.")
}
