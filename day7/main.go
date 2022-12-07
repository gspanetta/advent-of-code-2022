package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

var directorySize = make(map[string]int)
var rootDir = "D:\\02_FORTBILDUNG\\adventofcode2022\\day7\\rootdir"

func main() {
	createFileStructure()

	getSizeOfDir(rootDir)
	fmt.Println(directorySize)

	// Q1
	Q1_solution := getSumOfDirSizesU100k()
	fmt.Println("Q1: ", Q1_solution)

	// Q2
	totalDiskSpace := 70000000
	neededDiskSpace := 30000000
	usedDiskSpace := directorySize[rootDir]
	freeDiskSpace := totalDiskSpace - usedDiskSpace
	neededDiskSpace = neededDiskSpace - freeDiskSpace

	fmt.Println("Needed disk space: ", neededDiskSpace)

	smallestPossible := usedDiskSpace
	for _, v := range directorySize {
		if v >= neededDiskSpace && v < smallestPossible {
			smallestPossible = v
		}
	}
	fmt.Println("Smallest single directory size to free up total of 30M: ", smallestPossible)
}

func getSumOfDirSizesU100k() int {
	sum := 0
	for _, v := range directorySize {
		if v <= 100000 {
			sum += v
		}
	}
	return sum
}

func getSizeOfDir(path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	sum := 0
	for _, file := range files {
		if file.IsDir() {
			getSizeOfDir(path + "\\" + file.Name())
			//fmt.Println(file.Name())
			sum += directorySize[path+"\\"+file.Name()]
		} else {
			data, err := ioutil.ReadFile(path + "\\" + file.Name())
			if err != nil {
				log.Fatal(err)
			}
			size, err := strconv.Atoi(string(data))
			if err != nil {
				log.Fatal(err)
			}
			sum += size
		}
	}
	directorySize[path] = sum
}

func createFileStructure() {
	bs, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(bs), "\n")

	currentDir := rootDir

	for _, line := range lines {
		word := strings.Split(line, " ")
		if len(word) == 3 && word[1] == "cd" && word[2] == "/" {
			currentDir = rootDir
		} else if len(word) == 3 && word[1] == "cd" {
			newpath := currentDir + "\\" + word[2]
			err := os.MkdirAll(newpath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			currentDir = newpath
		} else if len(word) == 2 && word[0] == "dir" {
			newpath := currentDir + "\\" + word[1]
			err := os.MkdirAll(newpath, os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
		} else if len(word) == 2 && word[1] == "ls" {
			continue
		} else if len(word) == 2 && isNumeric(word[0]) {
			newFilepath := currentDir + "\\" + word[1]
			_, err = os.Create(newFilepath)
			if err != nil {
				log.Fatal(err)
			}
			err := ioutil.WriteFile(newFilepath, []byte(word[0]), 0666)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			fmt.Println("Can not recognize this command ", word)
		}
	}
}

func isNumeric(s string) bool {
	b := true
	for _, c := range s {
		if c < '0' || c > '9' {
			b = false
			break
		}
	}
	return b
}
