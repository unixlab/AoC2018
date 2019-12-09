package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

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

	for len(steps) > 0 {
		var possible []string
		for step, deps := range steps {
			if len(deps) == 0 {
				possible = append(possible, step)
			}
		}

		sort.Strings(possible)
		step := possible[0]

		delete(steps, step)
		for char, depends := range steps {
			for k, v := range depends {
				if v == step {
					steps[char] = append(steps[char][:k], steps[char][k+1:]...)
				}
			}
		}
		fmt.Printf("%s", step)
	}
	fmt.Println()
}
