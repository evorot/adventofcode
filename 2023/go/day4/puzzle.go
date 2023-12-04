package day4

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
	file, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)
	re := regexp.MustCompile(`Card\s+(\d+):\s+([\d\s]+)\s+\|\s+([\d\s]+)`)
	var sumPoints int64
	for reader.Scan() {
		text := reader.Text()
		matches := re.FindAllStringSubmatch(text, -1)
		winCards := parseNumbers(matches[0][2])
		winCardsMap := make(map[int64]struct{}, len(winCards))
		for _, card := range winCards {
			winCardsMap[card] = struct{}{}
		}
		myCards := parseNumbers(matches[0][3])

		var points int64
		isFirstPoint := true
		for _, card := range myCards {
			if _, exist := winCardsMap[card]; exist {
				if isFirstPoint {
					points += 1
					isFirstPoint = false
					continue
				}
				points *= 2
			}
		}
		sumPoints += points
	}

	fmt.Println("Day4. Puzzle1. Сумма очков в лотерее =", sumPoints)
}

func parseNumbers(str string) []int64 {
	numbers := strings.Fields(str)
	result := make([]int64, 0, len(numbers))
	for _, numberStr := range numbers {
		number, err := strconv.ParseInt(numberStr, 10, 64)
		if err == nil {
			result = append(result, number)
		}
	}
	return result
}

func Puzzle2() {
	file, err := os.Open("./day4/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewScanner(file)

	re := regexp.MustCompile(`Card\s+(\d+):\s+([\d\s]+)\s+\|\s+([\d\s]+)`)

	cardsCounter := make(map[int]int)
	var allInput []string
	for reader.Scan() {
		allInput = append(allInput, reader.Text())
	}
	lenAllInput := len(allInput)
	for _, input := range allInput {
		matches := re.FindAllStringSubmatch(input, -1)
		cardNumber, _ := strconv.Atoi(matches[0][1])
		oldQty := cardsCounter[cardNumber]
		cardsCounter[cardNumber] = oldQty + 1

		winCards := parseNumbers(matches[0][2])
		winCardsMap := make(map[int64]struct{}, len(winCards))
		for _, card := range winCards {
			winCardsMap[card] = struct{}{}
		}
		myCards := parseNumbers(matches[0][3])

		var qty int
		for _, card := range myCards {
			if _, exist := winCardsMap[card]; exist {
				qty += 1
			}
		}

		for i := 0; i < cardsCounter[cardNumber]; i++ {
			for i := 1; i <= qty; i++ {
				if i > lenAllInput {
					break
				}
				oldQty := cardsCounter[cardNumber+i]
				cardsCounter[cardNumber+i] = oldQty + 1
			}
		}

	}
	var sumPoints int
	for _, qty := range cardsCounter {
		sumPoints += qty
	}

	fmt.Println("Day4. Puzzle1. Сумма scratchcards =", sumPoints)
}
