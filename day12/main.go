package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type outstandingAction struct {
	pot   int
	plant bool
}

func nextGen(pots map[int]bool, notes map[string]bool) map[int]bool {
	var tempPots strings.Builder
	var actions []outstandingAction
	for pod := range pots {
		for i := pod - 2; i <= pod+2; i++ {
			for j := -2; j <= 2; j++ {
				val, exists := pots[i+j]
				if exists && val {
					tempPots.WriteString("#")
				} else {
					tempPots.WriteString(".")
				}
			}
			action := false
			bing := tempPots.String()
			action = notes[bing]
			actions = append(actions, outstandingAction{i, action})
			tempPots.Reset()
		}
	}
	for _, action := range actions {
		if action.plant {
			pots[action.pot] = true
		} else {
			delete(pots, action.pot)
		}
	}
	return pots
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	pots := make(map[int]bool, 2000)
	notes := make(map[string]bool, 50)

	for scanner.Scan() {
		line := scanner.Text()

		if strings.Index(line, "state") > -1 {
			for k, v := range strings.Split(line[15:], "") {
				if v == "#" {
					pots[k] = true
				}
			}
		}

		if strings.Index(line, "=>") > -1 {
			if line[9:] == "#" {
				notes[line[:5]] = true
			} else {
				notes[line[:5]] = false
			}
		}
	}

	part2Sum := 0
	for i := 1; i <= 200; i++ {
		pots = nextGen(pots, notes)

		if i == 20 {
			sum := 0
			for k, v := range pots {
				if v {
					sum += k
				}
			}
			fmt.Printf("part 1 = %d\n", sum)
		}
		// after some generation the increase in linear
		if i == 199 {
			for k, v := range pots {
				if v {
					part2Sum += k
				}
			}
		}
		if i == 200 {
			newSum := 0
			for k, v := range pots {
				if v {
					newSum += k
				}
			}
			increase := newSum - part2Sum
			part2Sum = newSum + (increase * (50e9 - 200))
			fmt.Printf("part 2 = %d\n", part2Sum)
		}
	}
}
