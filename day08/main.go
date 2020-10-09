package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMetadataSum(data []int) (int, int) {
	var sum int
	var offset int
	offset += 2
	for i := 0; i < data[0]; i++ {
		childSum, length := getMetadataSum(data[offset:])
		sum += childSum
		offset += length
	}
	for i := 0; i < data[1]; i++ {
		sum += data[offset+i]
	}
	offset += data[1]
	return sum, offset
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var line string

	// we know we only get one line
	for scanner.Scan() {
		line = scanner.Text()
	}

	var data []int

	for _, v := range strings.Split(line, " ") {
		number, _ := strconv.Atoi(v)
		data = append(data, number)
	}

	metadataSum, _ := getMetadataSum(data)
	fmt.Printf("sum => %d\n", metadataSum)
}
