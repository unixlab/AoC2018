package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getMetadataSumPart1(data []int) (int, int) {
	var sum int
	var offset int
	offset += 2
	for i := 0; i < data[0]; i++ {
		childSum, length := getMetadataSumPart1(data[offset:])
		sum += childSum
		offset += length
	}
	for i := 0; i < data[1]; i++ {
		sum += data[offset+i]
	}
	offset += data[1]
	return sum, offset
}

func getMetadataSumPart2(data []int) (int, int) {
	var sum int
	var offset int
	var childSums []int
	offset += 2
	for i := 0; i < data[0]; i++ {
		childSum, length := getMetadataSumPart2(data[offset:])
		childSums = append(childSums, childSum)
		offset += length
	}
	if data[0] == 0 {
		for i := 0; i < data[1]; i++ {
			sum += data[offset+i]
		}
	} else {
		for i := 0; i < data[1]; i++ {
			if data[offset+i] <= data[0] && data[offset+i] > 0 {
				sum += childSums[data[offset+i]-1]
			}
		}
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

	metadataSum, _ := getMetadataSumPart1(data)
	fmt.Printf("sum => %d\n", metadataSum)

	metadataSum, _ = getMetadataSumPart2(data)
	fmt.Printf("sum => %d\n", metadataSum)
}
