package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	assignments := strings.Split(string(bs), "\n")

	// Q1
	num_duplicate_assignments := 0

	// Q2
	num_overlapping_assignments := 0

	for _, ass := range assignments {
		ass = strings.TrimSuffix(ass, "\r")
		elf := strings.Split(ass, ",")
		//fmt.Printf("%v, %v\n", elf[0], elf[1])
		elf0 := strings.Split(elf[0], "-")
		elf1 := strings.Split(elf[1], "-")

		elf0_begin, _ := strconv.Atoi(elf0[0])
		elf0_end, _ := strconv.Atoi(elf0[1])
		elf1_begin, _ := strconv.Atoi(elf1[0])
		elf1_end, _ := strconv.Atoi(elf1[1])

		// Q1
		if ((elf0_begin >= elf1_begin) && (elf0_end <= elf1_end)) || ((elf1_begin >= elf0_begin) && (elf1_end <= elf0_end)) {
			num_duplicate_assignments += 1
			//fmt.Println("is a duplicate assignment")
		}

		// Q2
		if !((elf0_begin < elf1_begin && elf0_end < elf1_begin) || (elf1_begin < elf0_begin && elf1_end < elf0_begin)) {
			num_overlapping_assignments += 1
		}
	}

	fmt.Println(num_duplicate_assignments)
	fmt.Println(num_overlapping_assignments)

}
