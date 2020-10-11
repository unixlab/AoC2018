package main

import "fmt"

const SERIAL = 6303

func getFuel(grid [300][300]int, size int) (int, int, int) {
	maxFuel := 0
	maxFuelX := 0
	maxFuelY := 0

	for x := 0; x <= 300-size; x++ {
		for y := 0; y <= 300-size; y++ {
			var curFuel int

			for i := 0; i < size; i++ {
				for j := 0; j < size; j++ {
					curFuel += grid[x+i][y+j]
				}
			}

			if curFuel > maxFuel {
				maxFuel = curFuel
				maxFuelX = x
				maxFuelY = y
			}
		}
	}

	return maxFuel, maxFuelX, maxFuelY
}

func main() {
	var grid [300][300]int

	for x := 0; x < 300; x++ {
		for y := 0; y < 300; y++ {
			rackID := x + 10
			powerLevel := rackID * y
			powerLevel += SERIAL
			powerLevel = powerLevel * rackID
			hundredsDigit := powerLevel % 1000 / 100
			powerLevel = hundredsDigit - 5
			grid[x][y] = powerLevel
		}
	}

	part1MaxFuel, part1MaxFuelX, part1MaxFuelY := getFuel(grid, 3)

	part2MaxFuel := 0
	part2MaxFuelX := 0
	part2MaxFuelY := 0
	part2MaxFuelSize := 0
	for i := 0; i < 300; i++ {
		maxFuel, maxFuelX, maxFuelY := getFuel(grid, i)
		if maxFuel > part2MaxFuel {
			part2MaxFuel = maxFuel
			part2MaxFuelX = maxFuelX
			part2MaxFuelY = maxFuelY
			part2MaxFuelSize = i
		}
	}

	fmt.Printf("max fuel part 1 is %d at %d,%d\n", part1MaxFuel, part1MaxFuelX, part1MaxFuelY)
	fmt.Printf("   answer is => %d,%d\n", part1MaxFuelX, part1MaxFuelY)
	fmt.Printf("max fuel part 2 is %d at %d,%d with size %d\n", part2MaxFuel, part2MaxFuelX, part2MaxFuelY, part2MaxFuelSize)
	fmt.Printf("   answer is => %d,%d,%d\n", part2MaxFuelX, part2MaxFuelY, part2MaxFuelSize)
}
