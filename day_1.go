package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func day1Part1() {
	file, err := os.Open("data/day_1.txt")

	if err != nil {
		fmt.Println("Error in opening file", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	leftList := []int{}
	rightList := []int{}
	for _, line := range lines {
		separatedLine := strings.Split(line, "   ")
		leftVal, err := strconv.Atoi(separatedLine[0])
		if err != nil {
			fmt.Printf("error while converting %s to int\n", separatedLine[0])
		}
		leftList = append(leftList, leftVal)

		rightVal, err := strconv.Atoi(separatedLine[1])
		if err != nil {
			fmt.Printf("error while converting %s to int\n", separatedLine[1])
		}
		rightList = append(rightList, rightVal)
	}

	sort.Slice(rightList, func(i, j int) bool { return rightList[i] < rightList[j] })
	sort.Slice(leftList, func(i, j int) bool { return leftList[i] < leftList[j] })

	fmt.Println(len(rightList))
	fmt.Println(len(leftList))

	result := 0.0
	for i := 0; i < len(rightList); i++ {
		fmt.Println(float64(leftList[i] - rightList[i]))
		result += math.Abs(float64(leftList[i] - rightList[i]))
	}

	fmt.Println(int(result))
}

func day1Part2() {
	file, err := os.Open("data/day_1.txt")

	if err != nil {
		fmt.Println("Error in opening file", err)
		os.Exit(1)
	}
	defer file.Close()

	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	leftList := []int{}
	rightListMap := make(map[int]int)
	for _, line := range lines {
		separatedLine := strings.Split(line, "   ")
		leftVal, err := strconv.Atoi(separatedLine[0])
		if err != nil {
			fmt.Printf("error while converting %s to int\n", separatedLine[0])
		}
		leftList = append(leftList, leftVal)

		rightVal, err := strconv.Atoi(separatedLine[1])
		if err != nil {
			fmt.Printf("error while converting %s to int\n", separatedLine[1])
		}
		rightListMap[rightVal] += 1
	}

	sort.Slice(leftList, func(i, j int) bool { return leftList[i] < leftList[j] })

	fmt.Println(len(rightListMap))
	fmt.Println(len(leftList))

	result := 0
	for i := 0; i < len(leftList); i++ {
		fmt.Println(leftList[i] * rightListMap[leftList[i]])
		result += leftList[i] * rightListMap[leftList[i]]
	}

	fmt.Println(int(result))
}
