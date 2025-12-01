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

func main() {
	file, err := os.Open("test.txt")
	zerocount := 0

	current := 50

	if err != nil {
		fmt.Println("Error Reading file")
		os.Exit(1)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		turn := getNumber(scanner.Text())
		current = (current + turn) % 100
		if current < 0 {
			current += 100
		}
		if current == 0 {
			zerocount++
		}
	}

	fmt.Println("The password is: ", zerocount)

}
