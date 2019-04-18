package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sum(input [60]int) int {
	sum := 0
	for i := range input {
		sum += input[i]
	}
	return sum
}

func main() {
	file, err := os.Open("d04-input.txt")
	if err != nil {
		panic(err)
	}

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	sort.Strings(lines)

	guards := make(map[int][60]int)
	currentGuard := -1
	sleepSince := -1

	for _, line := range lines {
		hashIndex := strings.Index(line, "#")
		if hashIndex > 0 {
			endIndex := strings.Index(line[hashIndex:], " ")
			currentGuard, _ = strconv.Atoi(line[hashIndex+1:hashIndex+endIndex])
			continue
		}

		if strings.HasPrefix(line[19:], "falls asleep") && sleepSince == -1 {
			sleepSince, _ = strconv.Atoi(line[15:17])
			continue
		}

		if strings.HasPrefix(line[19:], "wakes up") && sleepSince >= 0 {
			now, _ := strconv.Atoi(line[15:17])
			currentGuardMinutes := guards[currentGuard]
			for i := sleepSince; i <= now; i++ {
				currentGuardMinutes[i]++
			}
			guards[currentGuard] = currentGuardMinutes
			sleepSince = -1
			continue
		}
		panic(line)
	}

	maxGuard := 0
	maxGuardMinutes := 0

	for guard, minutes := range guards {
		if sum(minutes) > maxGuardMinutes {
			maxGuard = guard
			maxGuardMinutes = sum(minutes)
		}
	}

	maxGuardMinute := -1
	maxGuardMinuteCounter := 0

	for minute, counter := range guards[maxGuard] {
		if counter > maxGuardMinuteCounter {
			maxGuardMinute = minute
			maxGuardMinuteCounter = counter
		}
	}

	fmt.Printf("%d x %d = %d\n", maxGuard, maxGuardMinute, maxGuard*maxGuardMinute)
}