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
	re := regexp.MustCompile(`[-+]?\b\d+\b`)
	combined := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	arr := combined.FindAllString(construct, -1)
	sum := 0
	enableChecking := 1
	for i := 0; i < len(arr); i++ {
		if arr[i] == "do()" {
			enableChecking = 1
			continue
		} else if arr[i] == "don't()" {
			enableChecking = 0
			continue
		}

		if enableChecking == 0 {
			continue
		} else if enableChecking == 1 {
			if arr[i] != "do()" && arr[i] != "don't()" {
				a, ea := strconv.Atoi(re.FindAllString(arr[i], -1)[0])
				handleError(ea)
				b, eb := strconv.Atoi(re.FindAllString(arr[i], -1)[1])
				handleError(eb)
				sum +=  a * b
			} else {
				continue
			}
		}
	}
	fmt.Println(sum)
}