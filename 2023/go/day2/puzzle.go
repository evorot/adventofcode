package day2

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Puzzle1() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	re := regexp.MustCompile(`Game (\d+): (.+)`)
	var sum int64
	for reader.Scan() {
		text := reader.Text()
		matches := re.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			gameNumber := match[1]
			sets := match[2]
			setsArray := strings.Split(sets, ";")
			okGame := true
			for _, set := range setsArray {
				colorQtyRe := regexp.MustCompile(`(\d+) (\w+)`)
				colorQtyMatches := colorQtyRe.FindAllStringSubmatch(set, -1)
				for _, colorQtyMatch := range colorQtyMatches {
					qtyStr := colorQtyMatch[1]
					qty, _ := strconv.ParseInt(qtyStr, 10, 64)
					color := colorQtyMatch[2]
					if !isOkSet(color, qty) {
						okGame = false
					}
				}
			}
			if okGame {
				num, _ := strconv.ParseInt(gameNumber, 10, 64)
				sum += num
			}
		}
	}

	fmt.Println("Day2. Puzzle1. Сумма индентификаторов валидных игр =", sum)
}

func Puzzle2() {
	file, err := os.Open("./day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	re := regexp.MustCompile(`Game (\d+): (.+)`)
	var sum int64
	for reader.Scan() {
		text := reader.Text()
		matches := re.FindAllStringSubmatch(text, -1)
		colorQty := make(map[string]int64)
		for _, match := range matches {
			sets := match[2]
			setsArray := strings.Split(sets, ";")
			for _, set := range setsArray {
				colorQtyRe := regexp.MustCompile(`(\d+) (\w+)`)
				colorQtyMatches := colorQtyRe.FindAllStringSubmatch(set, -1)
				for _, colorQtyMatch := range colorQtyMatches {
					qtyStr := colorQtyMatch[1]
					qty, _ := strconv.ParseInt(qtyStr, 10, 64)
					color := colorQtyMatch[2]
					currentQty, exist := colorQty[color]
					if !exist || qty > currentQty {
						colorQty[color] = qty
					}
				}
			}

			var multiQty int64
			initMulti := true
			for _, qty := range colorQty {
				if initMulti {
					multiQty = qty
					initMulti = false
				} else {
					multiQty *= qty
				}
			}
			sum += multiQty
		}
	}

	fmt.Println("Day2. Puzzle2. Мощность кубиков =", sum)
}

func isOkSet(color string, qty int64) bool {
	colorQty := map[string]int64{
		"red":   12,
		"green": 13,
		"blue":  14,
	}
	return colorQty[color] >= qty
}
