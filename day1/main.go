package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func getNumber(turn string) int {
	sign := turn[0]
	num := turn[1:]

	intNumber, err := strconv.Atoi(num)

	if err != nil {
		fmt.Println("error converting the stirng")
		os.Exit(1)
	}
	if sign == 'R' {
		return intNumber
	} else {
		return -1 * intNumber
	}

}

func countZeroCross(current, turn int) int {
	count := 0
	if turn > 0 {
		count = (current + turn) / 100
	} else if turn < 0 {
		val := -(current + turn)
		term1 := 0
		if val >= 0 {
			term1 = val / 100
		} else {
			term1 = (val - 99) / 100
		}

		term2 := 0
		if current > 0 {
			term2 = -1
		}

		count = term1 - term2
	}
	return count
}

func partOne(inputs []int) int {

	current := 50
	zerocount := 0

	for _, turn := range inputs {
		current = (current + turn) % 100
		if current < 0 {
			current += 100
		}

		if current == 0 {
			zerocount++
		}
	}

	return zerocount
}

func partTwo(inputs []int) int {
	current := 50
	zerocount := 0

	for _, turn := range inputs {
		zerocount += countZeroCross(current, turn)
		current = (current + turn) % 100
		if current < 0 {
			current += 100
		}
	}

	fmt.Println("Zero crossed:", zerocount)
	return zerocount
}

func main() {
	file, err := os.Open("input.txt")

	var inputs []int

	if err != nil {
		fmt.Println("Error Reading file")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		turn := getNumber(scanner.Text())
		inputs = append(inputs, turn)
	}

	fmt.Println("Anser for Part1: ", partOne(inputs))
	fmt.Println("Anser for Part2: ", partTwo(inputs))

}
