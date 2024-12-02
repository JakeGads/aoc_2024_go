package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("inputs/day2")
	if err != nil {
		log.Fatal(err)
	}

	x := parseData(data)
	day2_part1(x)
	day2_part2(x)
}

func day2_part1(data [][]int) {
	count := 0
	for _, d := range data {
		if validate(d) {
			count++
		}
	}
	log.Println("Part 1: ", count)
}

func day2_part2(data [][]int) {
	count := 0
	for _, d := range data {
		for i := 0; i < len(d); i++ {
			data_copy := make([]int, len(d))
			copy(data_copy, d)
			data_copy = removeIndex(data_copy, i)
			if validate(data_copy) {
				count++
				break
			}
		}
	}
	log.Println("Part 2: ", count)
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func validate(data []int) bool {
	if len(data) < 1 {
		return false
	}
	is_safe := true
	last := data[0]
	is_increaseing := data[1]-data[2] > 0
	for i := 1; i < len(data); i++ {
		if !is_safe {
			break
		}
		if is_increaseing && data[i] > last {
			is_safe = false
			break
		} else if !is_increaseing && data[i] < last {
			is_safe = false
			break
		}

		diff := int(math.Abs(float64(data[i] - last)))

		if !(diff >= 1 && diff <= 3) {
			is_safe = false
			break
		}

		last = data[i]
	}
	return is_safe
}

func parseData(data []byte) [][]int {
	lines := strings.Split(string(data), "\n")
	result := make([][]int, len(lines))

	for i, line := range lines {
		if line == "" {
			continue
		}
		nums := strings.Fields(line)
		ints := make([]int, len(nums))
		for j, num := range nums {
			n, err := strconv.Atoi(num)
			if err != nil {
				log.Fatal(err)
			}
			ints[j] = n
		}
		result[i] = ints
	}

	return result
}
