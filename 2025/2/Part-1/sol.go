package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	score := 0
	file, err := os.Open("input.txt")
	handle(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, ",")
		for _, item := range items {
			nums := strings.Split(item, "-")
			lb, err1 := strconv.Atoi(nums[0])
			ub, err2 := strconv.Atoi(nums[1])
			handle(err1)
			handle(err2)
			for i := lb; i <= ub; i++ {
				strnum := strconv.Itoa(i)
				if len(strnum) % 2 == 0 {
					num1, err1 := strconv.Atoi(strnum[:len(strnum)/2])
					handle(err1)
					num2, err2 := strconv.Atoi(strnum[len(strnum)/2:])
					handle(err2)
					num, err3 := strconv.Atoi(strnum)
					handle(err3)
					if num1 == num2 {
						score += num
						fmt.Println(fmt.Sprintf("(%d) for item: %s", num, item))
					}
				}
			}
		}
	}
	fmt.Println(score)
}