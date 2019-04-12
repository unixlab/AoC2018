package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("d02-input.txt")
	if err != nil {
		panic(err)
	}

	var ids []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		ids = append(ids, scanner.Text())
	}

	for _, id := range ids {
		for _, matcher := range ids {
			for i := 0; i < len(matcher); i++ {
				if id != matcher && id[:i] + id[i+1:] == matcher[:i] + matcher[i+1:] {
					fmt.Println(id)
					fmt.Println(matcher)
					fmt.Printf("%*s\n", i+1, id[i:i+1])
					fmt.Println(id[:i] + id[i+1:])
					fmt.Println()
				}
			}
		}
	}
}
