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
		fmt.Printf("%v\n", sum)
	}

	fmt.Println("Day3. Puzzle1. Сумма индентификаторов валидных игр =", sum)
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
