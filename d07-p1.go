package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Step struct {
	Before string
	Step   string
}

func appendIfUniq(inputArray []string, inputValue string) []string {
	for _, value := range inputArray {
		if value == inputValue {
			return inputArray
		}
	}
	return append(inputArray, inputValue)
}

func indexOf(element string, data []string) int {
	for key, value := range data {
		if value == element {
			return key
		}
	}
	return -1
}

func main() {
	// open file
	file, err := os.Open("d07-input-sub.txt")
	if err != nil {
		panic(err)
	}

	// array for all steps
	var steps []Step
	var characters []string

	// get steps and save in array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var step Step
		step.Before = scanner.Text()[5:6]
		step.Step = scanner.Text()[36:37]
		steps = append(steps, step)
		characters = appendIfUniq(characters, step.Before)
		characters = appendIfUniq(characters, step.Step)
	}

	done := false

	for !done {
		done = true
		for _, step := range steps {
			position := indexOf(step.Step, characters)
			for checkPosition := position; checkPosition > 0; checkPosition-- {
				if characters[checkPosition] == step.Before {
					fmt.Printf("move %s before %s\n", characters[checkPosition], characters[position])
					var newCharacters []string
					newCharacters = append(newCharacters, characters[:checkPosition]...)
					newCharacters = append(newCharacters, characters[checkPosition+1:position]...)
					newCharacters = append(newCharacters, characters[checkPosition])
					newCharacters = append(newCharacters, characters[position:]...)
					fmt.Println(characters)
					fmt.Println(newCharacters)
					fmt.Println()
					characters = newCharacters
				}
			}
		}
	}
	fmt.Println(strings.Join(characters, ""))
}
