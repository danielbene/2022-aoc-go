package main

import (
	"fmt"
	"myAwesomeModule/utils"
	"os"
	"time"
)

func part1() {
	utils.StartTimer()
	time.Sleep(5000)
	defer utils.WriteSolution("template - part1")
}

func part2() {
	utils.StartTimer()
	time.Sleep(1000)
	defer utils.WriteSolution("template - part2")
}

func main() {
	utils.InitDay()

	part1()
	part2()

	path, _ := os.Getwd()
	fmt.Println(path + " done.")
}
