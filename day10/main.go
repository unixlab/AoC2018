package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type star struct {
	posX int
	posY int
	velX int
	velY int
}

func getMinMaxXY(stars []star) (int, int, int, int) {
	var minX, maxX, minY, maxY int

	for _, v := range stars {
		if v.posX < minX {
			minX = v.posX
		}
		if v.posX > maxX {
			maxX = v.posX
		}
		if v.posY < minY {
			minY = v.posY
		}
		if v.posY > maxY {
			maxY = v.posY
		}
	}

	return minX, maxX, minY, maxY
}

func printStars(stars []star) {
	var line strings.Builder
	minX, maxX, minY, maxY := getMinMaxXY(stars)
	for i := minY; i <= maxY; i++ {
		line.Reset()
		for j := minX; j <= maxX; j++ {
			if checkExistingStar(j,i,stars){
				line.WriteRune('#')
			} else {
				line.WriteRune(' ')
			}
		}
		if strings.Index(line.String(), "#") > -1 {
			fmt.Println(line.String())
		}
	}
}

func checkExistingStar(x int, y int, stars []star) bool {
	for _ , v := range stars {
		if v.posY == y && v.posX == x {
			return true
		}
	}
	return false
}

func moveStars(stars []star) []star {
	for k , v := range stars {
		stars[k].posX = v.posX + v.velX
		stars[k].posY = v.posY + v.velY
	}
	return stars
}

func intAbs(n int) int {
	return int(math.Abs(float64(n)))
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	regex, _ := regexp.Compile("^position=<([0-9 -]+), ([0-9 -]+)> velocity=<([0-9 -]+), ([0-9 -]+)>$")

	var stars, initalStars []star

	for scanner.Scan() {
		line := scanner.Text()
		regexRes := regex.FindAllStringSubmatch(line, -1)
		gX, _ := strconv.Atoi(strings.ReplaceAll(regexRes[0][1], " ", ""))
		gY, _ := strconv.Atoi(strings.ReplaceAll(regexRes[0][2], " ", ""))
		vX, _ := strconv.Atoi(strings.ReplaceAll(regexRes[0][3], " ", ""))
		vY, _ := strconv.Atoi(strings.ReplaceAll(regexRes[0][4], " ", ""))

		stars = append(stars, star{gX, gY, vX, vY})
		initalStars = append(initalStars, star{gX, gY, vX, vY})
	}

	smallestField := -1
	smallestFieldSize := 0

	for i := 0; i < 100000; i++ {
		stars = moveStars(stars)
		minX, maxX, minY, maxY := getMinMaxXY(stars)
		minX = intAbs(minX)
		minY = intAbs(minY)

		fieldSize := (minX+maxX) * (minY+maxY)

		if fieldSize < smallestFieldSize || smallestField == -1 {
			smallestField = i
			smallestFieldSize = fieldSize
		}
	}

	for i := 0; i <= smallestField; i++ {
		initalStars = moveStars(initalStars)
	}

	fmt.Printf("this is the message after %d seconds\n", smallestField+1)
	printStars(initalStars)
}
