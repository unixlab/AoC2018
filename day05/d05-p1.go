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

		reacted := false

		for {
			reacted = false
			for i := 0; i < len(line)-1; i++ {
				if isUpper(line[i:i+1]) {
					if strings.ToLower(line[i:i+1]) == line[i+1:i+2] {
						line = line[:i] + line[i+2:]
						reacted = true
					}
					continue
				}
				if isLower(line[i:i+1]) {
					if strings.ToUpper(line[i:i+1]) == line[i+1:i+2] {
						line = line[:i] + line[i+2:]
						reacted = true
					}
					continue
				}
			}
			if !reacted {
				break
			}
		}

		fmt.Println(len(line))
}
