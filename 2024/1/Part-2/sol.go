package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func occurences(a []int, n int) int {
	ret := 0
	for i := 0; i < len(a); i++ {
		if a[i] == n {
			ret++
		}
	}
	return ret
}

func main() {
	var x []int
	var y []int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println(">>>")
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		arr := strings.Fields(line)
		a, e := strconv.Atoi(arr[0])
		handleError(e)
		b, e := strconv.Atoi(arr[1])
		handleError(e)
		x = append(x, a)
		y = append(y, b)
	}

	ret := 0
	for i := 0; i < len(x); i++ {
		ret += x[i] * occurances(y, x[i])
	}
	fmt.Println(ret)
}
