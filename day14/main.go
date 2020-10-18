package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Scoreboard struct {
	Begin           *Recipe
	ElfOne          *Recipe
	ElfTwo          *Recipe
	End             *Recipe
	NumberOfRecipes int
}

type Recipe struct {
	IntVal int
	StrVal string
	Next   *Recipe
	Prev   *Recipe
}

func newRecipe(strValue string) *Recipe {
	intValue, _ := strconv.Atoi(strValue)
	return &Recipe{intValue, strValue, nil, nil}
}

func (s Scoreboard) printScoreboard() {
	cur := s.Begin
	for {
		if cur == s.ElfOne {
			fmt.Printf("(%s)", cur.StrVal)
		} else if cur == s.ElfTwo {
			fmt.Printf("[%s]", cur.StrVal)
		} else {
			fmt.Printf(" %s ", cur.StrVal)
		}
		cur = cur.Next
		if cur == s.Begin {
			break
		}
	}
	fmt.Println()
}

func (s Scoreboard) makeNewRecipes() Scoreboard {
	digits := strconv.Itoa(s.ElfOne.IntVal + s.ElfTwo.IntVal)
	for _, digit := range strings.Split(digits, "") {
		s = s.addRecipe(digit)
	}
	return s
}

func (s Scoreboard) addRecipe(strValue string) Scoreboard {
	newRes := newRecipe(strValue)
	newRes.Next = s.Begin
	newRes.Prev = s.End

	s.End.Next = newRes
	s.Begin.Prev = newRes
	s.End = newRes

	s.NumberOfRecipes++

	return s
}

func (s Scoreboard) moveElfs() Scoreboard {
	value := s.ElfOne.IntVal
	for i := 0; i < value+1; i++ {
		s.ElfOne = s.ElfOne.Next
	}
	value = s.ElfTwo.IntVal
	for i := 0; i < value+1; i++ {
		s.ElfTwo = s.ElfTwo.Next
	}
	return s
}

func (s Scoreboard) toString() string {
	var scoreboardAsString strings.Builder
	cur := s.Begin
	scoreboardAsString.WriteString(s.Begin.StrVal)
	for cur != s.End {
		cur = cur.Next
		scoreboardAsString.WriteString(cur.StrVal)
	}
	return scoreboardAsString.String()
}

func (s Scoreboard) checkForWarmUp(warmUp int) bool {
	warmUpString := strconv.Itoa(warmUp)
	scoreboardString := s.toString()
	if strings.Index(scoreboardString, warmUpString) >= 0 {
		return true
	}
	return false
}

func (s Scoreboard) warmUpPos(warmUp int) int {
	warmUpString := strconv.Itoa(warmUp)
	scoreboardString := s.toString()
	bing := strings.Index(scoreboardString, warmUpString)
	return bing
}

func main() {
	var scoreboard Scoreboard

	warmUpTime := 290431

	startElfOne := newRecipe("3")
	startElfTwo := newRecipe("7")

	startElfOne.Next = startElfTwo
	startElfOne.Prev = startElfTwo

	startElfTwo.Next = startElfOne
	startElfTwo.Prev = startElfOne

	scoreboard.ElfOne = startElfOne
	scoreboard.ElfTwo = startElfTwo

	scoreboard.Begin = startElfOne
	scoreboard.End = startElfTwo

	scoreboard.NumberOfRecipes = 2

	// scoreboard.printScoreboard()

	var endOfWarmUpTime *Recipe
	for scoreboard.NumberOfRecipes < warmUpTime+10 {
		scoreboard = scoreboard.makeNewRecipes()
		scoreboard = scoreboard.moveElfs()

		// scoreboard.printScoreboard()

		if scoreboard.NumberOfRecipes >= warmUpTime && endOfWarmUpTime == nil {
			if scoreboard.NumberOfRecipes > warmUpTime {
				endOfWarmUpTime = scoreboard.End.Prev
			} else {
				endOfWarmUpTime = scoreboard.End
			}
		}

	}

	fmt.Print("part 1 => ")
	for i := 0; i < 10; i++ {
		endOfWarmUpTime = endOfWarmUpTime.Next
		fmt.Print(endOfWarmUpTime.StrVal)
	}
	fmt.Println()

	for !scoreboard.checkForWarmUp(warmUpTime) {
		for i := 0; i < 100000; i++ {
			scoreboard = scoreboard.makeNewRecipes()
			scoreboard = scoreboard.moveElfs()
		}
	}
	fmt.Printf("part 2 => %d\n", scoreboard.warmUpPos(warmUpTime))
}
