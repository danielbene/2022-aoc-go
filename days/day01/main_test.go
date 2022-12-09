package main

import (
	"myAwesomeModule/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	correct := 24000
	sol := part1()

	if sol != correct {
		t.Errorf("part1 test = %d; correct solution = %d", sol, correct)
	}
}

func TestPart2(t *testing.T) {
	correct := 45000
	sol := part2()

	if sol != correct {
		t.Errorf("part2 test = %d; correct solution = %d", sol, correct)
	}
}

// https://stackoverflow.com/questions/23729790/how-can-i-do-test-setup-using-the-testing-package-in-go
func TestMain(m *testing.M) {
	utils.InitDay()
	read(utils.GetTestInputPath())

	m.Run()
}
