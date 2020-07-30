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

	// loop over all coordinates - lets call this a grid
	for y := minY - 1; y <= maxY+1; y++ {
		for x := minX - 1; x <= maxX+1; x++ {
			// set minDistance higher than possible
			minDistance := maxY * maxX
			// define variable for distance id ( line number )
			var minDistanceId int
			// loop over all given coordinates
			for currentId, coordinate := range coordinates {
				// get distance from current grid position to current coordinate
				distance := getDistance(x, y, coordinate.X, coordinate.Y)
				// if we hit the smallest distance twice, clear the id
				if distance == minDistance {
					minDistanceId = 0
				}
				// if we have a smaller distance than before, set new minimum distance and id
				if distance < minDistance {
					minDistance = distance
					minDistanceId = currentId + 1
				}
			}
			if minDistance == 0 {
				// if distance is 0, we are exactly on a coordinate
				coordinates[minDistanceId-1].Counter++
			} else if minDistanceId == 0 {
				// if the id is 0, it was cleared and we have multiple shortest paths
				// than mean we do not need to count this
			} else {
				// increase counter for the nearest
				coordinates[minDistanceId-1].Counter++
				// if we are at the boarder, we are infinite
				// than we want to invalidate the current id
				if x == minX || x == maxX || y == minY || y == maxY {
					coordinates[minDistanceId-1].Infinite = true
				}
			}

		}
	}

	maxArea := 0
	for _, value := range coordinates {
		if !value.Infinite {
			if value.Counter > maxArea {
				maxArea = value.Counter
			}
		}
	}
	fmt.Println(maxArea)
}
