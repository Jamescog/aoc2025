package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func LargestBattery(btrline string) (string, int) {
	largest := 0
	index := 0

	for i := 0; i < len(btrline); i++ {
		intLine := int(btrline[i]) - int('0')

		if intLine > largest {
			largest = intLine
			index = i
		}
	}
	largestStr := strconv.Itoa(largest)
	return largestStr, index
}

func Part1(input []string) int {
	sum := 0
	for _, line := range input {
		final := ""
		validPart := line[:len(line)-1]
		largest1, index := LargestBattery(validPart)
		final += largest1
		validPart2 := line[index+1:]
		largest2, _ := LargestBattery(validPart2)
		final += largest2
		finalInt, _ := strconv.Atoi(final)

		sum += finalInt

	}

	return sum
}

func Get12Battries(line string) int {
	k := 12
	remove := len(line) - k

	stack := make([]byte, 0, len(line))

	for i := 0; i < len(line); i++ {
		digit := line[i]

		for remove > 0 && len(stack) > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			remove--
		}

		stack = append(stack, digit)
	}

	stack = stack[:k]

	val, _ := strconv.Atoi(string(stack))
	fmt.Println("Output:", val)
	return val
}

func Part2(input []string) int {
	sum := 0
	for _, line := range input {
		finalInt := Get12Battries(line)
		sum += finalInt

	}
	return sum
}

func main() {
	file, err := os.Open("input.txt")

	if err != nil {
		log.Fatal("Error Opening file")
	}
	defer file.Close()

	input := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	part1 := Part1(input)
	part2 := Part2(input)

	fmt.Println("Answer for part1 : ", part1)
	fmt.Println("Answer for part2 : ", part2)

}
