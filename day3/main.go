package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	prios := make(map[string]int)
	items := []string{
		"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m",
		"n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z",
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M",
		"N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	for i, item := range items {
		prios[item] = i + 1
	}

	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	backpacks := strings.Split(string(bs), "\n")

	// Q1
	sum_prios := 0
	same_item := 0

	for _, bp := range backpacks {
		nItems := len(bp)
		comp1 := bp[:nItems/2]
		comp2 := bp[nItems/2 : nItems]

		same_item = 0
		for _, item1 := range comp1 {
			for _, item2 := range comp2 {
				if item1 == item2 {
					same_item = int(item1)
					break
				}
			}
			if same_item != 0 {
				break
			}
		}

		if same_item < 91 && same_item > 64 {
			sum_prios += same_item - 38
		} else if same_item < 123 && same_item > 96 {
			sum_prios += same_item - 96
		} else {
			fmt.Println("Error: wtf ist this? ", same_item)
		}
	}

	fmt.Println(sum_prios)

	// Q2
	sum_prios = 0
	for i := 0; i < 300; i += 3 {
		threesome := []string{backpacks[i], backpacks[i+1], backpacks[i+2]}
		occurence := make(map[string]int)

		for _, item := range items {
			for _, letter := range threesome[0] {
				if item == string(letter) {
					occurence[item] = 1
				}
			}

			for _, letter := range threesome[1] {
				if item == string(letter) {
					if occurence[item] == 1 {
						occurence[item] = 2
					}
				}
			}

			for _, letter := range threesome[2] {
				if item == string(letter) {
					if occurence[item] == 2 {
						occurence[item] = 3
						//fmt.Println("threesome at ", i, " badge is ", item)
					}
				}
			}
		}
		for key, element := range occurence {
			if element == 3 {
				sum_prios += prios[key]
			}
		}

	}

	fmt.Println(sum_prios)
}
