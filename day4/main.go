package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getCountForLine(linestr string) (int, string) {
	lines := strings.Split(linestr, ",")
	sum := 0
	line1 := lines[0]
	line2 := lines[1]
	line3 := lines[2]

	finalStr := []rune(line2)

	for i, symbol := range line2 {
		localSum := 0
		totalLen := len(line2)
		if symbol == '@' {

			// left
			if i != 0 {
				if line2[i-1] == '@' {
					localSum++
				}
			}

			// upleft
			if i > 0 {
				if line1[i-1] == '@' {
					localSum++
				}
			}

			//right
			if i != totalLen-1 {
				if line2[i+1] == '@' {
					localSum++
				}
			}

			//upright
			if i != totalLen-1 {
				if line1[i+1] == '@' {
					localSum++
				}
			}

			//top
			if line1[i] == '@' {
				localSum++
			}

			//bottom

			if line3[i] == '@' {
				localSum++
			}

			// bottom-left
			if i != 0 {
				if line3[i-1] == '@' {
					localSum++
				}
			}

			//bottom-right

			if i != totalLen-1 {
				if line3[i+1] == '@' {
					localSum++
				}
			}

			//finally
			if localSum < 4 {
				sum++
				finalStr[i] = 'x'
			}

		}
	}
	return sum, string(finalStr)

}

func getRoundResult(lines []string) (int, []string) {
	roundSum := 0
	roundTransformed := []string{}

	for i := 0; i < len(lines); i++ {
		line1 := strings.Repeat(".", len(lines))
		line2 := lines[i]
		line3 := strings.Repeat(".", len(lines))

		if i != 0 {
			line1 = lines[i-1]
		}
		if i != len(lines)-1 {
			line3 = lines[i+1]
		}
		lineStr := strings.Join([]string{line1, line2, line3}, ",")
		gc, tranformed := getCountForLine(lineStr)
		roundSum += gc
		roundTransformed = append(roundTransformed, tranformed)
	}
	return roundSum, roundTransformed
}

func getInput(fileName string) []string {
	lines := []string{}
	file, _ := os.Open(fileName)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func main() {
	globalSum := 0

	lines := getInput("input.txt")
	part1, _ := getRoundResult(lines)
	fmt.Printf("Part1 Answer: %v\n", part1)

outerLoop:
	for {
		roundint, roundstr := getRoundResult(lines)
		globalSum += roundint
		if roundint > 0 {
			lines = roundstr
		} else {
			break outerLoop
		}
	}

	fmt.Printf("Part2 Answer: %v\n", globalSum)

}
