package main

import (
	"os"
	"bufio"
	"strings"
	"fmt"
	"strconv"
)

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	score := 0
	index := 50
	file, err := os.Open("input.txt")
	handle(err)
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		if strings.Compare(string(line[0]), "L") == 0 {
			mark := index
			num, err := strconv.Atoi(line[1:])
			handle(err)
			index = ((index - num) % 100)
			if index < 0 {
    			index += 100
			}
			crossed := 0
			zero := 0
			standardized := num % 100
			crossed += (num - (standardized)) / 100
			if (mark - standardized) < 0 && mark != 0 {
				crossed += 1
			}
			score += crossed
			if index == 0 {
				zero++
				score++
			}
			fmt.Println(fmt.Sprintf("%d left %d = %d; cross %d times; stopped at zero %d times", mark, num, index, crossed, zero))
		} else if strings.Compare(string(line[0]), "R") == 0 {
			mark := index
			num, err := strconv.Atoi(line[1:])
			handle(err)
			index = (index + num) % 100
			crossed := 0
			zero := 0
			standardized := num % 100
			crossed += (num - (standardized)) / 100
			if (mark + standardized) > 100 && mark != 0 {
				crossed += 1
			}
			score += crossed
			if index == 0 {
				zero++
				score++
			}
			fmt.Println(fmt.Sprintf("%d right %d = %d; cross %d times; zero %d times", mark, num, index, crossed, zero))
		} else {
			fmt.Println("input.txt is in the wrong format !!")
		}
	}
	fmt.Println(score)
}