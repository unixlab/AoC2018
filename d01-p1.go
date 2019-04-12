package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("d01-input.txt")
	if err != nil {
		panic(err)
	}

	count := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		count += number
	}

	fmt.Println(count)
}
