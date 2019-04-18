package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func isUpper(input string) bool {
	if input == strings.ToUpper(input) {
		return true
	} else {
		return false
	}
}

func isLower(input string) bool {
	if input == strings.ToLower(input) {
		return true
	} else {
		return false
	}
}

func doReact(input string) string {
	reacted := false
	for {
		reacted = false
		for i := 0; i < len(input)-1; i++ {
			if isUpper(input[i:i+1]) {
				if strings.ToLower(input[i:i+1]) == input[i+1:i+2] {
					input = input[:i] + input[i+2:]
					reacted = true
				}
				continue
			}
			if isLower(input[i:i+1]) {
				if strings.ToUpper(input[i:i+1]) == input[i+1:i+2] {
					input = input[:i] + input[i+2:]
					reacted = true
				}
				continue
			}
		}
		if !reacted {
			break
		}
	}
	return input
}

func main() {
	file, err := os.Open("d05-input.txt")
	if err != nil {
		panic(err)
	}

	var line string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line = scanner.Text()
	}

	letters := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	savedPolymer := doReact(line)
	minLength := len(savedPolymer)

	for _, letter := range letters {
		withoutLower := strings.ReplaceAll(savedPolymer, letter, "")
		newPolymer := strings.ReplaceAll(withoutLower, strings.ToUpper(letter), "")
		newPolymerLength := len(doReact(newPolymer))
		if newPolymerLength < minLength {
			minLength = newPolymerLength
		}
	}
	fmt.Println(minLength)
}
