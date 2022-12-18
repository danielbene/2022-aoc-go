package main

import (
	"myAwesomeModule/utils"
	"testing"
)

func TestPart1(t *testing.T) {
	correct := "CMZ"
	sol := part1()

	if sol != correct {
		t.Errorf("part1 test = %s; correct solution = %s", sol, correct)
	}
}

func TestPart2(t *testing.T) {
	correct := "MCD"
	sol := part2()

	if sol != correct {
		t.Errorf("part2 test = %s; correct solution = %s", sol, correct)
	}
}

func TestMain(m *testing.M) {
	utils.InitDay()
	read(utils.GetTestInputPath())

	m.Run()
}
