package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	Id int
	Left int
	Top int
	Height int
	Width int
}

func parseClaim(input string) Claim {
	var currentClaim Claim

	startIndex := 1
	endIndex := strings.Index(input, " ")
	currentClaim.Id, _ = strconv.Atoi(input[startIndex:endIndex])

	startIndex = endIndex + 3
	endIndex = strings.Index(input, ",")
	currentClaim.Left, _ = strconv.Atoi(input[startIndex:endIndex])

	startIndex = endIndex + 1
	endIndex = strings.Index(input, ":")
	currentClaim.Top, _ = strconv.Atoi(input[startIndex:endIndex])

	startIndex = endIndex + 2
	endIndex = strings.Index(input, "x")
	currentClaim.Width, _ = strconv.Atoi(input[startIndex:endIndex])

	startIndex = endIndex + 1
	currentClaim.Height, _ = strconv.Atoi(input[startIndex:])

	return currentClaim
}

func main() {
	file, err := os.Open("d03-input.txt")
	if err != nil {
		panic(err)
	}

	fabric := [1000][1000]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		claim := parseClaim(scanner.Text())
		for i := 0; i < claim.Width; i++ {
			for j := 0; j < claim.Height; j++ {
				fabric[claim.Left+i][claim.Top+j]++
			}
		}
	}

	counter := 0

	for i := 0; i < len(fabric); i++ {
		for j := 0; j < len(fabric[i]); j++ {
			if fabric[i][j] > 1 {
				counter++
			}
		}
	}

	fmt.Println(counter)
}