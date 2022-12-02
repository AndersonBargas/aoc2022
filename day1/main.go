package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	elvesCalories := make([]int, 0)

	elfCalories := make([]int, 0)
	for fileScanner.Scan() {
		itemCalories := fileScanner.Text()
		if itemCalories == "" {
			totalElfCalories := sumUpTheElfCalories(elfCalories)
			elvesCalories = append(elvesCalories, totalElfCalories)
			elfCalories = make([]int, 0)
			continue
		}

		convertedCalories, err := strconv.Atoi(itemCalories)
		if err != nil {
			panic(err)
		}
		elfCalories = append(elfCalories, convertedCalories)
	}

	readFile.Close()
	sort.Sort(sort.Reverse(sort.IntSlice(elvesCalories)))

	fmt.Printf("The elf with most calories have a total of %d\n", elvesCalories[0])
	fmt.Printf("The 3 top elves with the most calories, together have a total of %d\n", sumUpTheElfCalories(elvesCalories[:3]))
}

func sumUpTheElfCalories(elfCalories []int) int {
	sumUpResult := 0
	for _, calories := range elfCalories {
		sumUpResult += calories
	}
	return sumUpResult
}
