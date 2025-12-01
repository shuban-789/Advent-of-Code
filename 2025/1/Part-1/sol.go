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
			if index == 0 {
				score++
			}
			fmt.Println(fmt.Sprintf("%d left %d = %d", mark, num, index))
		} else if strings.Compare(string(line[0]), "R") == 0 {
			mark := index
			num, err := strconv.Atoi(line[1:])
			handle(err)
			index = (index + num) % 100
			if index == 0 {
				score++
			}
			fmt.Println(fmt.Sprintf("%d right %d = %d", mark, num, index))
		} else {
			fmt.Println("input.txt is in the wrong format !!")
		}
	}
	fmt.Println(score)
}