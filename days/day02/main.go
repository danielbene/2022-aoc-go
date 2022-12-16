package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
)

var col1 []int
var col2 []int

// part 1 sol: 13052

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
	// Y -draw, X -loose, Z -win

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
