package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type game struct {
	cur *marble
	len int
}

type marble struct {
	points int
	next   *marble
	prev   *marble
}

func newMarble(value int) *marble {
	return &marble{value, nil, nil}
}

func (g game) addMarble(value int) game {
	m := newMarble(value)
	m.prev = g.cur
	m.next = g.cur.next
	g.cur.next.prev = m
	g.cur.next = m
	g.cur = m
	g.len++
	return g
}

func (g game) removeMarble() game {
	g.cur.next.prev = g.cur.prev
	g.cur.prev.next = g.cur.next
	g.cur = g.cur.next
	return g
}

func (g game) next() game {
	g.cur = g.cur.next
	return g
}

func (g game) prev() game {
	g.cur = g.cur.prev
	return g
}

func (g game) prevN(n int) game {
	for i := 0; i < n; i++ {
		g.cur = g.cur.prev
	}
	return g
}

func getMaxScore(numberOfPlayers int, marbles int) (int, int) {
	players := make([]int, numberOfPlayers)

	var thisGame game
	m := newMarble(0)
	m.next = m
	m.prev = m
	thisGame.cur = m
	thisGame.len = 1

	for i := 1; i <= marbles; i++ {
		if i%23 == 0 {
			players[i%numberOfPlayers] += i
			thisGame = thisGame.prevN(8)
			players[i%numberOfPlayers] += thisGame.cur.points
			thisGame = thisGame.removeMarble()
			thisGame = thisGame.next()
			continue
		}
		thisGame = thisGame.addMarble(i)
		thisGame = thisGame.next()
	}

	maxPlayer := 0
	maxPlayerValue := 0
	for k, v := range players {
		if v > maxPlayerValue {
			maxPlayer = k
			maxPlayerValue = v
		}
	}

	return maxPlayer, maxPlayerValue
}

func main() {
	regex, _ := regexp.Compile("^([0-9]+) players; last marble is worth ([0-9]+) points$")

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		regexMatches := regex.FindAllStringSubmatch(scanner.Text(), -1)
		numberOfPlayers, _ := strconv.Atoi(regexMatches[0][1])
		marbles, _ := strconv.Atoi(regexMatches[0][2])

		// Part 1
		player, points := getMaxScore(numberOfPlayers, marbles)
		fmt.Printf("for part 1 player %d wins with %d points\n", player, points)

		// Part 2
		player, points = getMaxScore(numberOfPlayers, marbles*100)
		fmt.Printf("for part 2 player %d wins with %d points\n", player, points)
	}
}
