package days

import (
	"fmt"
	"myAwesomeModule/utils"
	"time"
)

func Part1() {
	utils.StartTimer()
	time.Sleep(1000)
	defer utils.WriteSolution("asd")
}

func Part2() {
	utils.StartTimer()
	time.Sleep(5000)
	defer utils.WriteSolution("fgh")
}

func Day1() {
	Part1()
	Part2()

	name := "day1"
	fmt.Println("Hello", name)

	utils.Hey()
}
