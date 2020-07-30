package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Coordinate struct {
	X        int
	Y        int
	Infinite bool
	Counter  int
}

func getDistance(x1 int, y1 int, x2 int, y2 int) int {
	return int(math.Abs(float64(x1)-float64(x2)) + math.Abs(float64(y1)-float64(y2)))
}

func main() {
	// open file
	file, err := os.Open("d06-input.txt")
	if err != nil {
		panic(err)
	}

	// get array for coordinates
	var coordinates []Coordinate

	// read file line by line and add coordinates to coordinate array
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ReplaceAll(scanner.Text(), " ", "")
		comma := strings.Index(line, ",")
		x, _ := strconv.Atoi(line[:comma])
		y, _ := strconv.Atoi(line[comma+1:])
		coordinates = append(coordinates, Coordinate{x, y, false, 0})
	}

	// define variables for min/max of x/y
	minX := -1
	maxX := -1
	minY := -1
	maxY := -1

	// get min/max of x/y
	for _, coordinate := range coordinates {
		// check min x
		if coordinate.X < minX || minX < 0 {
			minX = coordinate.X
		}
		// check max x
		if coordinate.X > maxX {
			maxX = coordinate.X
		}
		// check min y
		if coordinate.Y < minY || minY < 0 {
			minY = coordinate.Y
		}
		// check max y
		if coordinate.Y > maxY {
			maxY = coordinate.Y
		}
	}

	// set area size to zero
	areaSize := 0

	// loop over all coordinates - lets call this a grid
	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			cumulatedDistance := 0
			// loop over all given coordinates
			for _, coordinate := range coordinates {
				// add distance from current grid position to current coordinates cumulated distance
				cumulatedDistance += getDistance(x, y, coordinate.X, coordinate.Y)
			}
			// if the cumulated distance is smaller than 10k add the coordiate to the area
			if cumulatedDistance < 10000 {
				areaSize++
			}
		}
	}
	fmt.Println(areaSize)
}
