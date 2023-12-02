package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var sum int64
	for reader.Scan() {
		var strNum []string
		for _, elem := range reader.Text() {
			s := string(elem)
			_, err = strconv.ParseInt(s, 10, 64)
			if err == nil {
				strNum = append(strNum, s)
			}
		}
		if len(strNum) == 1 {
			num, _ := strconv.ParseInt(strNum[0]+strNum[0], 10, 64)
			sum += num
		} else if len(strNum) > 1 {
			newNum := strNum[0] + strNum[len(strNum)-1]
			num, _ := strconv.ParseInt(newNum, 10, 64)
			sum += num
		}
	}

	fmt.Println("Day1. Puzzle1. Сумма калибровочных значений =", sum)
}

var variants = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func Puzzle2() {
	file, err := os.Open("./day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	var sum int64

	for reader.Scan() {
		text := reader.Text()

		digitsByFirstIndex := findDigits(text, strings.Index)
		digitsByLastIndex := findDigits(text, strings.LastIndex)

		if len(digitsByFirstIndex) == 1 {
			num, _ := strconv.ParseInt(digitsByFirstIndex[0]+digitsByLastIndex[0], 10, 64)
			sum += num
		} else if len(digitsByFirstIndex) > 1 {
			newNum := digitsByFirstIndex[0] + digitsByLastIndex[len(digitsByLastIndex)-1]
			num, _ := strconv.ParseInt(newNum, 10, 64)
			sum += num
		}
	}

	fmt.Println("Day1. Puzzle2. Сумма калибровочных значений (с учетом чисел ввиде слов) =", sum)
}

func convertToStrNum(s string) (string, bool) {
	wordNum := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	strNum, ok := wordNum[s]
	return strNum, ok
}

func findDigits(s string, indexFunc func(s string, substr string) int) []string {
	indexVariant := make(map[int]string, len(variants))
	var indexArray []int
	for _, variant := range variants {
		i := indexFunc(s, variant)
		if i < 0 {
			continue
		}
		if convertedStr, ok := convertToStrNum(variant); ok {
			indexVariant[i] = convertedStr
		} else {
			indexVariant[i] = variant
		}
		indexArray = append(indexArray, i)
	}
	slices.Sort(indexArray)

	result := make([]string, 0, len(indexArray))
	for _, i := range indexArray {
		result = append(result, indexVariant[i])
	}
	return result
}
