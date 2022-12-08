package main

import (
	"fmt"
	"myAwesomeModule/utils"
	"os"
)

func part1() {
	utils.StartTimer()

	// part1 solution here

	defer utils.WriteSolution("template - part1")
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
