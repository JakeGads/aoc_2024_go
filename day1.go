package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read the file and print its contents
	data, _ := os.ReadFile("inputs/day1")

	list1, list2 := gather(data)
	log.Println("Part 1: ", part1(list1, list2))
	log.Println("Part 2: ", part2(list1, list2))
}

func part1(list1 []int, list2 []int) int {
	// Sort the lists
	sort.Ints(list1)
	sort.Ints(list2)
	sum := 0
	for i := 0; i < len(list1); i++ {
		sum += int(math.Abs(float64(list1[i] - list2[i])))
	}
	return sum
}

func part2(list1 []int, list2 []int) int {
	simularity_score := 0
	for i := 0; i < len(list1); i++ {
		simularity_score += list1[i] * count_occurances(list2, list1[i])
	}
	return simularity_score
}

func count_occurances(list []int, num int) int {
	count := 0
	for _, n := range list {
		if n == num {
			count++
		}
	}
	return count
}

func gather(data []byte) ([]int, []int) {
	list1 := []int{}
	list2 := []int{}

	// Convert the data to a list of integers
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		nums := strings.Split(line, "   ")

		num1, _ := strconv.Atoi(nums[0])
		num2, _ := strconv.Atoi(nums[1])

		list1 = append(list1, num1)
		list2 = append(list2, num2)
	}

	return list1, list2
}
