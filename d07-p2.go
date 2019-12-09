package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Worker struct {
	CurrentWork string
	Remainder   int
}

func main() {
	file, _ := os.Open("d07-input.txt")
	scanner := bufio.NewScanner(file)

	steps := make(map[string][]string, 26)
	for scanner.Scan() {
		steps[scanner.Text()[36:37]] = append(steps[scanner.Text()[36:37]], scanner.Text()[5:6])
		_, keyExists := steps[scanner.Text()[5:6]]
		if !keyExists {
			steps[scanner.Text()[5:6]] = []string{}
		}
	}

	var worker [5]Worker
	neededTime := 0

	for {

		for workerNumber, workerData := range worker {
			if workerData.Remainder > 0 {
				worker[workerNumber].Remainder = workerData.Remainder - 1
			}
			if workerData.Remainder == 0 {
				for char, depends := range steps {
					for k, v := range depends {
						if v == workerData.CurrentWork {
							steps[char] = append(steps[char][:k], steps[char][k+1:]...)
						}
					}
					worker[workerNumber].CurrentWork = ""
				}
			}
		}

		var possible []string
		for step, deps := range steps {
			if len(deps) == 0 {
				possible = append(possible, step)
			}
		}
		sort.Strings(possible)

		for workerNumber, workerData := range worker {
			if workerData.CurrentWork == "" && len(possible) > 0 {
				worker[workerNumber].CurrentWork = possible[0]
				worker[workerNumber].Remainder = 60 + int(possible[0][0]-65)
				delete(steps, possible[0])
				possible = possible[1:]
			}
		}

		neededTime++

		allIdle := true
		for _, workerData := range worker {
			if workerData.Remainder > 0 {
				allIdle = false
			}
		}

		if len(steps) == 0 && allIdle {
			break
		}
	}
	fmt.Println(neededTime)
}
