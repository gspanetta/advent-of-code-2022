package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	totalCalLists := strings.Split(string(bs), "\n\n")

	var elfCalories []int

	for _, list := range totalCalLists {
		sCalList := strings.Split(list, "\n")

		sumCalories := 0
		for _, sCal := range sCalList {
			cal, err := strconv.Atoi(sCal)
			if err == nil {
				sumCalories += cal
			}
		}

		elfCalories = append(elfCalories, sumCalories)
	}

	//fmt.Println(elfCalories)

	// Q1 - which elf carries the most calories?
	biggestCal := 0
	elfWithBiggestCal := 0
	for i, sumCal := range elfCalories {
		if sumCal > biggestCal {
			biggestCal = sumCal
			elfWithBiggestCal = i + 1
		}
	}

	fmt.Println("Elf ", elfWithBiggestCal, " is carrying ", biggestCal, " calories!")

	// Q2 - top 3 elves
	sort.Ints(elfCalories)
	top3cal := elfCalories[len(elfCalories)-1] + elfCalories[len(elfCalories)-2] + elfCalories[len(elfCalories)-3]

	fmt.Println("Top 3 in total is ", top3cal)
}
