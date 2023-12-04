package main

import (
	"aoc2023/day1"
	"aoc2023/day2"
	"aoc2023/day3"
	"fmt"
)

func main() {
	wrap(day1.Puzzle1, day1.Puzzle2)
	wrap(day2.Puzzle1, day2.Puzzle2)
	wrap(day3.Puzzle1, day3.Puzzle2)
}

func wrap(puzzle1, puzzle2 func()) {
	fmt.Println("********************")
	puzzle1()
	fmt.Println("--------------------")
	puzzle2()
	fmt.Println("********************")
}
