package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

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

	maxGuardMinute := 0
	maxGuardMinuteCounter := 0
	maxGuard := 0

	for guard, minutes := range guards {
		for minute, minuteCounter := range minutes {
			if minuteCounter > maxGuardMinuteCounter {
				maxGuardMinute = minute
				maxGuardMinuteCounter = minuteCounter
				maxGuard = guard
			}
		}
	}

	fmt.Printf("%d x %d = %d\n", maxGuard, maxGuardMinute, maxGuard*maxGuardMinute)
}