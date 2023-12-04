package day3

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func Puzzle1() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	var sum int64
	var allInput []string
	for reader.Scan() {
		allInput = append(allInput, reader.Text())
	}
	re := regexp.MustCompile(`\d+`)
	for inputIndex, input := range allInput {
		indexes := re.FindAllStringSubmatchIndex(input, -1)
		for idx, index := range indexes {
			startIndex := index[0]
			endIndex := index[1]
			exist := false
			if inputIndex != 0 {
				prevInput := allInput[inputIndex-1]
				exist = hasStarInLineByIndex(prevInput, startIndex, endIndex)
			}
			if !exist && inputIndex != len(allInput)-1 {
				nextInput := allInput[inputIndex+1]
				exist = hasStarInLineByIndex(nextInput, startIndex, endIndex)
			}
			if !exist {
				exist = hasStarInLineByIndex(input, startIndex, endIndex)
			}

			if exist {
				numStr := re.FindAllStringSubmatch(input, -1)[idx][0]
				num, _ := strconv.ParseInt(numStr, 10, 64)
				sum += num
				continue
			}
		}
	}

	fmt.Println("Day3. Puzzle1. Сумма подходящих номеров двигателя =", sum)
}

type Point struct {
	x int
	y int
}

func Puzzle2() {
	file, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var sum int64
	suitableStars := make(map[Point][]int64)
	var allInput []string
	for reader.Scan() {
		allInput = append(allInput, reader.Text())
	}
	re := regexp.MustCompile(`\d+`)
	for inputIndex, input := range allInput {
		indexes := re.FindAllStringSubmatchIndex(input, -1)
		for idx, index := range indexes {
			startIndex := index[0]
			endIndex := index[1]
			if inputIndex != 0 {
				prevInput := allInput[inputIndex-1]
				if point, exist := hasSymbolInLineByIndex(prevInput, inputIndex-1, startIndex, endIndex); exist {
					numStr := re.FindAllStringSubmatch(input, -1)[idx][0]
					num, _ := strconv.ParseInt(numStr, 10, 64)
					suitableStars[point] = append(suitableStars[point], num)
				}
			}
			if inputIndex != len(allInput)-1 {
				nextInput := allInput[inputIndex+1]
				if point, exist := hasSymbolInLineByIndex(nextInput, inputIndex+1, startIndex, endIndex); exist {
					numStr := re.FindAllStringSubmatch(input, -1)[idx][0]
					num, _ := strconv.ParseInt(numStr, 10, 64)
					suitableStars[point] = append(suitableStars[point], num)
				}
			}
			if point, exist := hasSymbolInLineByIndex(input, inputIndex, startIndex, endIndex); exist {
				numStr := re.FindAllStringSubmatch(input, -1)[idx][0]
				num, _ := strconv.ParseInt(numStr, 10, 64)
				suitableStars[point] = append(suitableStars[point], num)
			}
		}
	}

	for _, nums := range suitableStars {
		if len(nums) > 1 {
			var power int64
			init := true
			for _, num := range nums {
				if init {
					power = num
					init = false
					continue
				}
				power *= num
			}
			sum += power
		}
	}

	fmt.Println("Day3. Puzzle2. Сумма передаточных мощностей =", sum)
}

func hasStarInLineByIndex(s string, startIndex, endIndex int) bool {
	for i := startIndex - 1; i <= endIndex; i++ {
		if i < 0 || (i == len(s)) {
			continue
		}
		elem := string(s[i])

		if !isNum(elem) && !isDot(elem) {
			return true
		}
	}
	return false
}

func isNum(numStr string) bool {
	_, err := strconv.ParseInt(numStr, 10, 64)
	if err == nil {
		return true
	}
	return false
}

func isDot(s string) bool {
	return s == "."
}

func hasSymbolInLineByIndex(s string, indexRow, startIndex, endIndex int) (Point, bool) {
	for i := startIndex - 1; i <= endIndex; i++ {
		if i < 0 || (i == len(s)) {
			continue
		}
		elem := string(s[i])

		if elem == "*" {
			return Point{x: indexRow, y: i}, true
		}
	}
	return Point{}, false
}
