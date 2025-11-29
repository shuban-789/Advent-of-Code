package main

import (
    "fmt"
    "bufio"
    "os"
	"strings"
)

func check(grid [][]string, x int, y int) bool {
	count := 0
	maxthresh := len(grid) - 2
	minthresh := len(grid) - maxthresh - 1

	if x <= maxthresh && y <= maxthresh {
		if strings.Compare(grid[x+1][y+1], "M") == 0 {
			if x >= minthresh && y >= minthresh {
				if strings.Compare(grid[x-1][y-1], "S") == 0 {
					count += 1
				}
			}
		}
	}

	if x >= minthresh && y >= minthresh {
		if strings.Compare(grid[x-1][y-1], "M") == 0 {
			if x <= maxthresh && y <= maxthresh {
				if strings.Compare(grid[x+1][y+1], "S") == 0 {
					count += 1
				}
			}
		}
	}

	if x >= minthresh && y <= maxthresh {
		if strings.Compare(grid[x-1][y+1], "M") == 0 {
			if x <= maxthresh && y >= minthresh {
				if strings.Compare(grid[x+1][y-1], "S") == 0 {
					count += 1
				}
			}
		}
	}


	if x <= maxthresh && y >= minthresh {
		if strings.Compare(grid[x+1][y-1], "M") == 0 {
			if x >= minthresh && y <= maxthresh {
				if strings.Compare(grid[x-1][y+1], "S") == 0 {
					count += 1
				}
			}
		}
	}

	if count == 2 {
		return true
	}
	return false
}

func main() {
    var grid [][]string
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
		arr := strings.Split(line, "")
		grid = append(grid, arr)
	}
	fmt.Println(grid)
	total := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if strings.Compare(grid[i][j], "A") == 0 {
				if check(grid, i, j) {
					total += 1
				}
			} else {
				continue
			}
		}
	}
	fmt.Println(total)
}