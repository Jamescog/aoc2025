package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func IntRange(r string) (int, int) {
	rangeSplit := strings.Split(r, "-")
	lower := rangeSplit[0]
	upper := rangeSplit[1]
	lowerInt, _ := strconv.Atoi(lower)
	upperInt, _ := strconv.Atoi(upper)

	return lowerInt, upperInt
}

func IsInRange(r, candidate string) bool {

	lowerInt, upperInt := IntRange(r)
	candidateInt, _ := strconv.Atoi(candidate)

	return candidateInt >= lowerInt && candidateInt <= upperInt
}

func GetOverlap(bounds []string) (bool, string) {
	old, new := bounds[0], bounds[1]
	oldStart, _ := strconv.Atoi(strings.Split(old, "-")[0])
	oldEnd, _ := strconv.Atoi(strings.Split(old, "-")[1])
	newStart, _ := strconv.Atoi(strings.Split(new, "-")[0])
	newEnd, _ := strconv.Atoi(strings.Split(new, "-")[1])
	if newStart <= oldEnd {
		start := oldStart
		end := oldEnd
		if newEnd > oldEnd {
			end = newEnd
		}
		return true, fmt.Sprintf("%d-%d", start, end)
	}
	return false, new
}

func SortRanges(ranges []string) []string {
	sort.Slice(ranges, func(i, j int) bool {
		aParts := strings.Split(ranges[i], "-")
		bParts := strings.Split(ranges[j], "-")
		aNum, _ := strconv.Atoi(aParts[0])
		bNum, _ := strconv.Atoi(bParts[0])
		return aNum < bNum
	})
	return ranges
}

func SweepRanges(rangeInclusive []string) []string {
	finalRanges := []string{}
	sorted := SortRanges(rangeInclusive)
	for _, r := range sorted {
		if len(finalRanges) == 0 {
			finalRanges = append(finalRanges, r)
		} else {
			ok, merged := GetOverlap([]string{finalRanges[len(finalRanges)-1], r})
			if ok {
				finalRanges[len(finalRanges)-1] = merged
			} else {
				finalRanges = append(finalRanges, r)
			}
		}
	}
	return finalRanges
}

func Part2(rangeInclusive []string) int {
	fmt.Printf("Staring with total of %v\n", len(rangeInclusive))
	finalRanges := SweepRanges(rangeInclusive)
	fmt.Printf("Sweeped to  total of %v\n", len(finalRanges))

	finalSum := 0

	for _, r := range finalRanges {
		l, u := IntRange(r)

		diff := u - l + 1

		finalSum += diff
	}

	return finalSum

}

func Part1(rangeInclusive, availableIngredients []string) int {
	freshCount := 0

	for _, ingredeint := range availableIngredients {
	innerLoop:
		for _, ranges := range rangeInclusive {
			if IsInRange(ranges, ingredeint) {
				freshCount++
				break innerLoop
			}
		}
	}
	return freshCount
}

func main() {
	file, _ := os.Open("input.txt")
	rangeInclusive := []string{}
	availableIngredients := []string{}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		}
		if strings.Contains(text, "-") {
			rangeInclusive = append(rangeInclusive, text)
		} else {
			availableIngredients = append(availableIngredients, text)
		}
	}

	fmt.Printf("Part1: %v\n", Part1(rangeInclusive, availableIngredients))
	fmt.Printf("Part2: %v\n", Part2(rangeInclusive))

}
