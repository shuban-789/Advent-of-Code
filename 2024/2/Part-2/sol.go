package main

import (
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"sort"
	"reflect"
	"os"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func incrementCheck(arr []int) bool {
	score := 0
	for j := 0; j < len(arr); j++ {
		if j > 0 {
			if abs(arr[j] - arr[j-1]) >= 1 {
				if abs(arr[j] - arr[j-1]) <= 3 {
					score++
				}
			}
		}
	}
	if score == len(arr) - 1 {
		return true
	}
	return false
}

func secondChance(chunk []int) int {
	score := 0
	for j := 0; j < len(chunk); j++ {
		var arr []int
		if j == len(chunk)-1 {
			arr = append([]int(nil), chunk[:j]...)
		} else {
			arr = append([]int(nil), chunk[:j]...)
			arr = append(arr, chunk[j+1:]...)
		}
		increase := append([]int(nil), arr...)
		decrease := append([]int(nil), arr...)
		sort.Ints(increase)
		sort.Sort(sort.Reverse(sort.IntSlice(decrease)))
		if reflect.DeepEqual(arr, increase) || reflect.DeepEqual(arr, decrease) {
			if incrementCheck(arr) {
				score += 1
			}
		} else {
			continue
		}
	}
	return score
}

func main() {
	var matrix [][]int
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		arr := strings.Fields(line)
		var row []int
		for i := 0; i < len(arr); i++ {
			num, err := strconv.Atoi(arr[i])
			handleError(err)
			row = append(row, num,)
		}
		matrix = append(matrix, row)
	}

	safe := 0
	for i := 0; i < len(matrix); i++ {
		increase := append([]int(nil), matrix[i]...)
		decrease := append([]int(nil), matrix[i]...)
		sort.Ints(increase)
		sort.Sort(sort.Reverse(sort.IntSlice(decrease)))
		if reflect.DeepEqual(matrix[i], increase) || reflect.DeepEqual(matrix[i], decrease) {
			if incrementCheck(matrix[i]) {
				safe++
			} else {
				if secondChance(matrix[i]) > 0 {
					safe++
				}
			}
		} else {
			if secondChance(matrix[i]) > 0 {
				safe++
			}
		}
	}
	fmt.Println(safe)
}