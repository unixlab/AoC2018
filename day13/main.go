package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type mine struct {
	cars  []car
	order []int
	track [][]string
}

type car struct {
	direction int
	turns     int
	lane      int
	pos       int
	crashed	  bool
}

func (m mine) updateCarOrder() mine {
	carValues := make(map[int]int)
	var values []int
	var curOrder []int

	for carNumber, car := range m.cars {
		value := car.lane*1000 + car.pos
		values = append(values, value)
		carValues[value] = carNumber
	}

	sort.Ints(values)

	for _, value := range values {
		curOrder = append(curOrder, carValues[value])
	}

	m.order = curOrder

	return m
}

func (m mine) printTrack() {
	for laneNumber, lane := range m.track {
		for fieldNumber, _ := range lane {
			fmt.Print(m.getPosWithCar(laneNumber, fieldNumber))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (m mine) getPosWithCar(lane int, pos int) string {
	for _, car := range m.cars {
		if car.lane == lane && car.pos == pos {
			switch car.direction {
			case 0:
				return "^"
			case 1:
				return ">"
			case 2:
				return "v"
			case 3:
				return "<"
			}
		}
	}
	return m.track[lane][pos]
}

func (c car) turnRight() car {
	c.direction++
	if c.direction == 4 {
		c.direction = 0
	}
	return c
}

func (c car) turnLeft() car {
	c.direction--
	if c.direction == -1 {
		c.direction = 3
	}
	return c
}

func (m mine) checkCrash() (bool, int) {
	allPos := make(map[int]bool)
	for _, car := range m.cars {
		if car.crashed {
			continue
		}
		value := car.lane*1000 + car.pos
		_, exists := allPos[value]
		if exists {
			return true, value
		} else {
			allPos[value] = true
		}
	}
	return false, 0
}

func (m mine) markCarByValue(value int) mine {
	var newOrder []int
	for _, carId := range m.order {
		if carId >= 0 {
			car := m.cars[carId]
			if car.lane*1000+car.pos != value {
				newOrder = append(newOrder, carId)
			} else {
				m.cars[carId].crashed = true
				newOrder = append(newOrder, -1)
			}
		} else {

			newOrder = append(newOrder, -1)
		}
	}
	m.order = newOrder
	return m
}

func (m mine) removeCarsByMark() mine {
	var newCars []car
	for _, carId := range m.order {
		if carId >= 0 {
			car := m.cars[carId]
			//newOrder = append(newOrder, carId)
			newCars = append(newCars, car)
		}
	}
	m.cars = newCars
	m = m.updateCarOrder()
	return m
}

func main() {
	var mine mine
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	curLane := 0
	for scanner.Scan() {
		curTrack := strings.Split(scanner.Text(), "")

		mine.track = append(mine.track, curTrack)

		for curPos, field := range curTrack {
			curDir := -1
			switch field {
			case "^":
				curDir = 0
			case ">":
				curDir = 1
			case "v":
				curDir = 2
			case "<":
				curDir = 3
			}
			if curDir >= 0 {
				mine.cars = append(mine.cars, car{curDir, 0, curLane, curPos, false})
				mine = mine.updateCarOrder()
				if curDir%2 == 0 {
					mine.track[curLane][curPos] = "|"
				} else {
					mine.track[curLane][curPos] = "-"
				}
			}
		}
		curLane++
	}

	noCrash := true
	for len(mine.cars) > 1 {
		for _, carId := range mine.order {
			if carId == -1 {
				continue
			}
			curCar := mine.cars[carId]

			switch curCar.direction {
			case 0:
				curCar.lane -= 1
			case 1:
				curCar.pos += 1
			case 2:
				curCar.lane += 1
			case 3:
				curCar.pos -= 1
			}

			switch mine.track[curCar.lane][curCar.pos] {
			case "\\":
				if curCar.direction == 0 {
					curCar = curCar.turnLeft()
					break
				}
				if curCar.direction == 1 {
					curCar = curCar.turnRight()
					break
				}
				if curCar.direction == 2 {
					curCar = curCar.turnLeft()
					break
				}
				if curCar.direction == 3 {
					curCar = curCar.turnRight()
					break
				}
			case "/":
				if curCar.direction == 0 {
					curCar = curCar.turnRight()
					break
				}
				if curCar.direction == 1 {
					curCar = curCar.turnLeft()
					break
				}
				if curCar.direction == 2 {
					curCar = curCar.turnRight()
					break
				}
				if curCar.direction == 3 {
					curCar = curCar.turnLeft()
					break
				}
			case "+":
				if curCar.turns%3 == 0 {
					curCar = curCar.turnLeft()
					curCar.turns++
					break
				}
				if curCar.turns%3 == 1 {
					curCar.turns++
					break
				}
				if curCar.turns%3 == 2 {
					curCar = curCar.turnRight()
					curCar.turns++
					break
				}
			}
			mine.cars[carId] = curCar

			crash, crashValue := mine.checkCrash()
			if crash {
				if noCrash {
					fmt.Printf("crash at %d,%d\n", curCar.pos, curCar.lane)
				}
				noCrash = false
				mine = mine.markCarByValue(crashValue)
			}
		}
		mine = mine.removeCarsByMark()
	}
	fmt.Printf("location of last car is %d,%d\n", mine.cars[0].pos, mine.cars[0].lane)
}
