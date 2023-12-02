package main

import (
	"aoc2023/day1"
	"fmt"
)

func main() {
	wrap(day1.Puzzle1, day1.Puzzle2)
}

func wrap(puzzle1, puzzle2 func()) {
	fmt.Println("********************")
	puzzle1()
	fmt.Println("--------------------")
	puzzle2()
	fmt.Println("********************")
}
