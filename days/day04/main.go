package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
	"strconv"
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

func part1() int {
	utils.StartTimer()

	cnt := 0
	for _, line := range input {
		parts := strings.Split(line, ",")
		elf1 := strings.Split(parts[0], "-")
		elf2 := strings.Split(parts[1], "-")

		e11, _ := strconv.Atoi(elf1[0])
		e12, _ := strconv.Atoi(elf1[1])
		e21, _ := strconv.Atoi(elf2[0])
		e22, _ := strconv.Atoi(elf2[1])

		if ((e11 <= e21 && e21 <= e12) && (e11 <= e22 && e22 <= e12)) ||
			((e21 <= e11 && e11 <= e22) && (e21 <= e12 && e12 <= e22)) {
			cnt += 1
		}
	}

	return cnt
}

func part2() int {
	utils.StartTimer()

	cnt := 0
	for _, line := range input {
		parts := strings.Split(line, ",")
		elf1 := strings.Split(parts[0], "-")
		elf2 := strings.Split(parts[1], "-")

		e11, _ := strconv.Atoi(elf1[0])
		e12, _ := strconv.Atoi(elf1[1])
		e21, _ := strconv.Atoi(elf2[0])
		e22, _ := strconv.Atoi(elf2[1])

		if ((e11 <= e21 && e21 <= e12) || (e11 <= e22 && e22 <= e12)) ||
			((e21 <= e11 && e11 <= e22) || (e21 <= e12 && e12 <= e22)) {
			cnt += 1
		}
	}

	return cnt
}

func main() {
	utils.InitDay()

	read(utils.GetInputPath())

	utils.WriteSolutionInt(part1())
	utils.WriteSolutionInt(part2())

	fmt.Println("done.")
}
