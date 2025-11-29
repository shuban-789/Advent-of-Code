package main

import (
	"fmt"
	"bufio"
	"regexp"
	"os"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
			panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var construct string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		construct += line
	}
	scanerr := scanner.Err()
	handleError(scanerr)
	re1 := regexp.MustCompile(`mul\(\d+,\d+\)`)
	re2 := regexp.MustCompile(`[-+]?\b\d+\b`)
	arr := re1.FindAllString(construct, -1)
	sum := 0
	for i := 0; i < len(arr); i++ {
		a, ea := strconv.Atoi(re2.FindAllString(arr[i], -1)[0])
		handleError(ea)
		b, eb := strconv.Atoi(re2.FindAllString(arr[i], -1)[1])
		handleError(eb)
		sum +=  a * b
	}
	fmt.Println(sum)
}