package main

import (
	"bufio"
	"fmt"
	"myAwesomeModule/utils"
	"os"
	"strconv"
)

// TODO: implement testing

func part1() {
	utils.StartTimer()

	// https://stackoverflow.com/questions/8757389/reading-a-file-line-by-line-in-go
	file, err := os.Open(utils.GetInputPath())
	utils.CheckError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	mostCal := 0
	elfCal := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			if mostCal < elfCal {
				mostCal = elfCal
			}

			elfCal = 0
		} else {
			cal, _ := strconv.Atoi(line)
			elfCal += cal
		}
	}

	utils.CheckError(scanner.Err())

	defer utils.WriteSolution(strconv.Itoa(mostCal))
}

func part2() {
	utils.StartTimer()

	// part2 solution here

	defer utils.WriteSolution("template - part2")
}

func main() {
	utils.InitDay()

	part1()
	part2()

	path, _ := os.Getwd()
	fmt.Println(path + " done.")
}
