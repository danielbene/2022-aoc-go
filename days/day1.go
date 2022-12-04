package days

import (
	"fmt"
	"myAwesomeModule/utils"
	"time"
)

func Day1Part1() {
	utils.StartTimer()
	time.Sleep(1000)
	defer utils.WriteSolution("asd")
}

func Day1Part2() {
	utils.StartTimer()
	time.Sleep(5000)
	defer utils.WriteSolution("fgh")
}

func Day1() {
	Day1Part1()
	Day1Part2()

	fmt.Println(utils.GetDayID() + " done.")
}
