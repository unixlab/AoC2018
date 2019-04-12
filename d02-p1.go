package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("d02-input.txt")
	if err != nil {
		panic(err)
	}

	letters := []string{"a","b","c","d","e","f","g","h","i","j","k","l","m","n","o","p","q","r","s","t","u","v","w","x","y","z"}
	two := 0
	three := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		foundTwo := false
		foundThree := false
		for _, letter := range letters {
			letterCount := strings.Count(scanner.Text(), letter)
			if letterCount == 2 && !foundTwo {
				foundTwo = true
				two++
			}
			if letterCount == 3 && !foundThree {
				foundThree = true
				three++
			}
		}
	}
	fmt.Printf("%d x %d = %d\n", two, three, two*three)
}
