package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func isUnique(bs []byte) bool {
	occurences1 := make(map[byte]bool)

	for _, b := range bs {
		if _, ok := occurences1[b]; ok {
			return false
		} else {
			occurences1[b] = true
		}
	}
	return true
}

func main() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	// Q1
	for i, _ := range bs {
		if i < 4 {
			continue
		}

		seq := bs[i-4 : i]

		if isUnique(seq) {
			fmt.Println("found start sequence at ", i)
			break
		}

	}

	// Q1
	for i, _ := range bs {
		if i < 14 {
			continue
		}

		seq := bs[i-14 : i]

		if isUnique(seq) {
			fmt.Println("found start sequence at ", i)
			break
		}

	}
}
