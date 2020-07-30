package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isUniq(element int, list []int) bool {
	for _, value := range list {
		if value == element {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("d01-input.txt")
	if err != nil {
		panic(err)
	}

	var numbers []int
	count := 0

	ReadFile:
		for {
			file.Seek(0,0)
			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				number, err := strconv.Atoi(scanner.Text())
				if err != nil {
					panic(err)
				}

				count += number

				if isUniq(count, numbers) {
					numbers = append(numbers, count)
				} else {
					fmt.Println(count)
					break ReadFile
				}
			}
		}
}
