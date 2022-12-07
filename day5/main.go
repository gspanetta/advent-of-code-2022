package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type tStack []string

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	parts := strings.Split(string(bs), "\n\n")

	//starting_stacks := parts[0]
	procedure := strings.Split(parts[1], "\n")
	//procedure := []string{"move 1 from 2 to 1", "move 3 from 1 to 3", "move 2 from 2 to 1", "move 1 from 1 to 2"}

	stacks := make([][]string, 9)
	stacks[0] = []string{"C", "Z", "N", "B", "M", "W", "Q", "V"}
	stacks[1] = []string{"H", "Z", "R", "W", "C", "B"}
	stacks[2] = []string{"F", "Q", "R", "J"}
	stacks[3] = []string{"Z", "S", "W", "H", "F", "N", "M", "T"}
	stacks[4] = []string{"G", "W", "F", "L", "N", "Q", "P"}
	stacks[5] = []string{"L", "P", "W"}
	stacks[6] = []string{"V", "B", "D", "R", "G", "C", "Q", "J"}
	stacks[7] = []string{"Z", "Q", "N", "B", "W"}
	stacks[8] = []string{"H", "L", "F", "C", "G", "T", "J"}

	//stacks := make([]tStack, 3)
	//stacks[0] = tStack{"Z", "N"}
	//stacks[1] = tStack{"M", "C", "D"}
	//stacks[2] = tStack{"P"}

	for _, step := range procedure {
		if len(step) < 8 {
			break
		}
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		command := re.FindAllString(step, -1)

		// translate
		amount, _ := strconv.Atoi(command[0])
		from, _ := strconv.Atoi(command[1])
		from = from - 1
		to, _ := strconv.Atoi(command[2])
		to = to - 1

		for i := 0; i < amount; i++ {
			e := stacks[from][len(stacks[from])-1]
			stacks[from][len(stacks[from])-1] = ""
			stacks[from] = stacks[from][:len(stacks[from])-1]
			stacks[to] = append(stacks[to], e)
		}
	}

	fmt.Println(stacks)

	// Q2
	stacks[0] = []string{"C", "Z", "N", "B", "M", "W", "Q", "V"}
	stacks[1] = []string{"H", "Z", "R", "W", "C", "B"}
	stacks[2] = []string{"F", "Q", "R", "J"}
	stacks[3] = []string{"Z", "S", "W", "H", "F", "N", "M", "T"}
	stacks[4] = []string{"G", "W", "F", "L", "N", "Q", "P"}
	stacks[5] = []string{"L", "P", "W"}
	stacks[6] = []string{"V", "B", "D", "R", "G", "C", "Q", "J"}
	stacks[7] = []string{"Z", "Q", "N", "B", "W"}
	stacks[8] = []string{"H", "L", "F", "C", "G", "T", "J"}

	cheat_stack := []string{}

	for _, step := range procedure {
		if len(step) < 8 {
			break
		}
		re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
		command := re.FindAllString(step, -1)

		// translate
		amount, _ := strconv.Atoi(command[0])
		from, _ := strconv.Atoi(command[1])
		from = from - 1
		to, _ := strconv.Atoi(command[2])
		to = to - 1

		for i := 0; i < amount; i++ {
			e := stacks[from][len(stacks[from])-1]
			stacks[from][len(stacks[from])-1] = ""
			stacks[from] = stacks[from][:len(stacks[from])-1]
			cheat_stack = append(cheat_stack, e)
		}

		for i := 0; i < amount; i++ {
			e := cheat_stack[len(cheat_stack)-1]
			cheat_stack[len(cheat_stack)-1] = ""
			cheat_stack = cheat_stack[:len(cheat_stack)-1]
			stacks[to] = append(stacks[to], e)
		}
	}

	fmt.Println(stacks)
}
