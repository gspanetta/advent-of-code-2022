package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

// input = 99x99
// testinput = 5x5
const arrSize = 99
const dataFile = "input.txt"

var treeArr [arrSize][arrSize]int
var treeArrT [arrSize][arrSize]int
var visibleTreeArr [arrSize][arrSize]int

func fillTreeArraysWithData(bs []byte) {
	x, y := 0, 0
	for _, b := range bs {
		if b != '\n' {
			treeArr[y][x], _ = strconv.Atoi(string(b))
			treeArrT[x][y] = treeArr[y][x]
			x++
		} else {
			y++
			x = 0
		}
	}
}

func checkVisibleTrees() {
	// check rows
	for y, line := range treeArr {
		max := 0
		for x := 0; x < arrSize; x++ {
			if x == 0 || y == 0 {
				visibleTreeArr[x][y] = 1
				max = line[x]
			} else if line[x] > max {
				max = line[x]
				visibleTreeArr[x][y] = 1
			}
		}
		max = 0
		for x := arrSize - 1; x > 0; x-- {
			if x == arrSize || y == arrSize {
				visibleTreeArr[x][y] = 1
				max = line[x]
			} else if line[x] > max {
				max = line[x]
				visibleTreeArr[x][y] = 1
			}
		}
	}

	// check columns (transformated rows)
	for y, line := range treeArrT {
		max := 0
		for x := 0; x < arrSize; x++ {
			if x == 0 || y == 0 {
				visibleTreeArr[y][x] = 1
				max = line[x]
			} else if line[x] > max {
				max = line[x]
				visibleTreeArr[y][x] = 1
			}
		}
		max = 0
		for x := arrSize - 1; x > 0; x-- {
			if x == arrSize-1 || y == arrSize-1 {
				visibleTreeArr[y][x] = 1
				max = line[x]
			} else if line[x] > max {
				max = line[x]
				visibleTreeArr[y][x] = 1
			}
		}
	}
}

func countVisibleTrees(visibleTreeArr [arrSize][arrSize]int) int {
	res := 0
	for _, line := range visibleTreeArr {
		for _, col := range line {
			res += col
		}
	}
	return res
}

func highestScenicScore() int {
	max := 0
	for x, line := range treeArr {
		for y, tree := range line {
			score1 := getVisibleTreesLeft(x, y)
			score2 := getVisibleTreesRight(x, y)
			score3 := getVisibleTreesAbove(x, y)
			score4 := getVisibleTreesBelow(x, y)
			score_total := score1 * score2 * score3 * score4

			fmt.Println("score of ", x, "/", y, "(", tree, ") is ", score_total)

			if score_total > max {
				max = score_total
			}
		}
	}
	return max
}

func getVisibleTreesLeft(x int, y int) int {
	tree_counter := 0
	tree_height := treeArr[x][y]

	if y == 0 {
		tree_counter = 0
	} else {
		for _y := y - 1; _y >= 0; _y-- {
			tree_counter++
			if treeArr[x][_y] >= tree_height {
				break
			}
		}
	}

	return tree_counter
}

func getVisibleTreesRight(x int, y int) int {
	tree_counter := 0
	tree_height := treeArr[x][y]

	if y == arrSize-1 {
		tree_counter = 0
	} else {
		for _y := y + 1; _y < arrSize; _y++ {
			tree_counter++
			if treeArr[x][_y] >= tree_height {
				break
			}
		}
	}

	return tree_counter
}

func getVisibleTreesAbove(x int, y int) int {
	tree_counter := 0
	tree_height := treeArr[x][y]

	if x == 0 {
		tree_counter = 0
	} else {
		for _x := x - 1; _x >= 0; _x-- {
			tree_counter++
			if treeArr[_x][y] >= tree_height {
				break
			}
		}
	}

	return tree_counter
}

func getVisibleTreesBelow(x int, y int) int {
	tree_counter := 0
	tree_height := treeArr[x][y]

	if x == arrSize-1 {
		tree_counter = 0
	} else {
		for _x := x + 1; _x < arrSize; _x++ {
			tree_counter++
			if treeArr[_x][y] >= tree_height {
				break
			}
		}
	}

	return tree_counter
}

func main() {
	bs, err := ioutil.ReadFile(dataFile)
	if err != nil {
		log.Fatal(err)
	}

	fillTreeArraysWithData(bs)
	//printArr(treeArr)

	checkVisibleTrees()
	printArr(visibleTreeArr)

	q1_num_visible_trees := countVisibleTrees(visibleTreeArr)
	fmt.Println("Q1: ", q1_num_visible_trees)

	q2_highest_scenic_score := highestScenicScore()
	fmt.Println("Q2: ", q2_highest_scenic_score)

}

func printArr(a [arrSize][arrSize]int) {
	for _, line := range a {
		fmt.Println(line)
	}
}
