package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func IsInvalidPartOne(id int) bool {
	strNum := strconv.Itoa(id)
	strlen := len(strNum)
	if strlen%2 != 0 {
		return false
	}
	half := strlen / 2
	first_half := strNum[0:half]
	second_half := strNum[half:strlen]
	return first_half == second_half
}

func IsRepetition(s string) bool {
	n := len(s)
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sub := s[:i]
			repeated := ""
			for j := 0; j < n/i; j++ {
				repeated += sub
			}
			if repeated == s {
				return true
			}
		}
	}
	return false
}

func parseRange(strRange string) (int, int) {
	intr := strings.Split(strRange, "-")
	start, err := strconv.Atoi(intr[0])
	if err != nil {
		log.Fatal("error converting str to int")
	}
	end, err := strconv.Atoi(intr[1])
	if err != nil {
		log.Fatal("error converting str to int")
	}
	return start, end
}

func processRange(start, end int, processFunc func(int) bool) int {
	sum := 0
	for i := start; i <= end; i++ {
		if processFunc(i) {
			sum += i
		}
	}
	return sum
}

func IsInValidPartTwo(fileStr string) int {
	invalidCount := 0
	ranges := strings.Split(fileStr, ",")
	for _, localRange := range ranges {
		small, large := parseRange(localRange)
		invalidCount += processRange(small, large, func(i int) bool {
			return IsRepetition(strconv.Itoa(i))
		})
	}
	return invalidCount
}

func main() {
	file, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln("error reading file")
	}

	inputStr := string(file)

	partOneStart := time.Now()
	sum := 0
	ranges := strings.Split(inputStr, ",")
	for _, el := range ranges {
		start, end := parseRange(el)
		sum += processRange(start, end, IsInvalidPartOne)
	}
	partOneElapsed := time.Since(partOneStart)

	partTwoStart := time.Now()
	partTwoResult := IsInValidPartTwo(inputStr)
	partTwoElapsed := time.Since(partTwoStart)

	fmt.Println("Final Result: ", sum)
	fmt.Printf("Part One took %v to finish execution\n", partOneElapsed)
	fmt.Println("Final Result Part 2: ", partTwoResult)
	fmt.Printf("Part Two took %v to finish execution\n", partTwoElapsed)
}
