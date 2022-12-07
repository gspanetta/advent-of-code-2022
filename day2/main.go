package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// ROCK: 	A, X	1
// PAPER: 	B, Y	2
// SCIS: 	C, Z	3

// Win:  6
// Draw: 3
// Loss: 0
func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	rounds := strings.Split(string(bs), "\n")

	// Q1
	totalPoints := 0
	for _, r := range rounds {
		if len(r) < 3 {
			continue
		}
		strat := strings.Split(r, " ")

		//fmt.Printf("%v-%v\n", strat[0], strat[1])

		if strat[1] == "X" { // Rock
			totalPoints += 1
			if strat[0] == "A" {
				totalPoints += 3
			} else if strat[0] == "B" {
				totalPoints += 0
			} else if strat[0] == "C" {
				totalPoints += 6
			}
		} else if strat[1] == "Y" { // Paper
			totalPoints += 2
			if strat[0] == "A" {
				totalPoints += 6
			} else if strat[0] == "B" {
				totalPoints += 3
			} else if strat[0] == "C" {
				totalPoints += 0
			}
		} else if strat[1] == "Z" { // Scis
			totalPoints += 3
			if strat[0] == "A" {
				totalPoints += 0
			} else if strat[0] == "B" {
				totalPoints += 6
			} else if strat[0] == "C" {
				totalPoints += 3
			}
		}
	}

	fmt.Println(totalPoints)

	// Q2
	// X - need to lose (0)
	// Y - need to draw (3)
	// Z - need to win (6)
	totalPoints = 0
	for _, r := range rounds {
		if len(r) < 3 {
			continue
		}
		strat := strings.Split(r, " ")

		if strat[1] == "X" { // lose
			totalPoints += 0
			if strat[0] == "A" { // rock - need to choose scis to lose
				totalPoints += 3
			} else if strat[0] == "B" { // paper - need to choose rock to lose
				totalPoints += 1
			} else if strat[0] == "C" { // scis - need to chose paper to lose
				totalPoints += 2
			}
		} else if strat[1] == "Y" { // draw
			totalPoints += 3
			if strat[0] == "A" { // rock - need to chose rock to draw
				totalPoints += 1
			} else if strat[0] == "B" { // paper - need to chose paper to draw
				totalPoints += 2
			} else if strat[0] == "C" { // scis - need to chose scis to draw
				totalPoints += 3
			}
		} else if strat[1] == "Z" { // win
			totalPoints += 6
			if strat[0] == "A" { // rock - need to chose paper to win
				totalPoints += 2
			} else if strat[0] == "B" { // paper - need to chose scis to win
				totalPoints += 3
			} else if strat[0] == "C" { // scis - need to chose rock to win
				totalPoints += 1
			}
		}
	}

	fmt.Println(totalPoints)
}
