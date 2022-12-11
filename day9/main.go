package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type position struct {
	xH int
	yH int
	xT int
	yT int
}

var positionMap = make(map[string]int)

func (p position) move(direction string, amount int) {
	for i := 0; i < amount; i++ {
		if direction == "U" {
			p.yH--
		} else if direction == "D" {
			p.yH++
		} else if direction == "L" {
			p.xH--
		} else if direction == "R" {
			p.xH++
		}

		if p.yH-p.yT > 1 {
			p.yT++
		} else if p.yT-p.yH > 1 {
			p.yT--
		} else if p.xH-p.xT > 1 {
			p.xT++
		} else if p.xT-p.xH > 1 {
			p.xT--
		}

		tailPosStr := strings.Join([]string{strconv.Itoa(p.xT), strconv.Itoa(p.yT)}, "|")
		positionMap[tailPosStr] = 1
	}

}

func main() {
	bs, err := ioutil.ReadFile("testinput.txt")
	if err != nil {
		log.Fatal(err)
	}

	commandList := strings.Split(string(bs), "\n")

	var ropePos = position{0, 0, 0, 0}
	positionMap["0|0"] = 1

	for _, cmd := range commandList {
		cmdSplit := strings.Split(cmd, " ")
		direction := cmdSplit[0]
		amount, _ := strconv.Atoi(cmdSplit[1])
		ropePos.move(direction, amount)
	}

	fmt.Println(positionMap)
	fmt.Println(len(positionMap))
}
