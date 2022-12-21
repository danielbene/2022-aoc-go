package main

import (
	"myAwesomeModule/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	correct := 95437
	sol := part1()

	if sol != correct {
		t.Errorf("part1 test = %d; correct solution = %d", sol, correct)
	}
}

func TestPart2(t *testing.T) {
	correct := 24933642
	sol := part2()

	if sol != correct {
		t.Errorf("part2 test = %d; correct solution = %d", sol, correct)
	}
}

func TestMain(m *testing.M) {
	utils.InitDay()
	read(utils.GetTestInputPath())

	m.Run()
}
