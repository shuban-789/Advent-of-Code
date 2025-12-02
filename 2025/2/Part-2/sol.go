package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
	"github.com/glenn-brown/golang-pkg-pcre/src/pkg/pcre" // please speed i need this ahh import
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
				regex := pcre.MustCompile(`^(.+)\1+$`, 0)
				matcher := regex.MatcherString(strnum, 0)
				if matcher.Matches() {
					score += i
					fmt.Println(fmt.Sprintf("(%d) for item: %s", i, item))
				}
			}
		}
	}
	fmt.Println(score)
}