package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func horizontal(grid [][]string, x int, y int) int {
	count := 0
	maxthresh := (len(grid)) - 4
	minthresh := (len(grid)) - maxthresh - 1
	if x <= maxthresh {
		if strings.Compare(grid[x+1][y], "M") == 0 {
			if strings.Compare(grid[x+2][y], "A") == 0 {
				if strings.Compare(grid[x+3][y], "S") == 0 {
					count += 1
				}
			}
		}
	}

	if x >= minthresh {
		if strings.Compare(grid[x-1][y], "M") == 0 {
			if strings.Compare(grid[x-2][y], "A") == 0 {
				if strings.Compare(grid[x-3][y], "S") == 0 {
					count += 1
				}
			}
		}
	}

	return count
}

func vertical(grid [][]string, x int, y int) int {
	count := 0
	maxthresh := (len(grid)) - 4
	minthresh := (len(grid)) - maxthresh - 1
	if y <= maxthresh {
		if strings.Compare(grid[x][y+1], "M") == 0 {
			if strings.Compare(grid[x][y+2], "A") == 0 {
				if strings.Compare(grid[x][y+3], "S") == 0 {
					count += 1
				}
			}
		}
	}

	if y >= minthresh {
		if strings.Compare(grid[x][y-1], "M") == 0 {
			if strings.Compare(grid[x][y-2], "A") == 0 {
				if strings.Compare(grid[x][y-3], "S") == 0 {
					count += 1
				}
			}
		}
	}

	return count
}

func diag1(grid [][]string, x int, y int) int {
	count := 0
	maxthresh := (len(grid)) - 4
	minthresh := (len(grid)) - maxthresh - 1
	if x <= maxthresh && y <= maxthresh {
		if strings.Compare(grid[x+1][y+1], "M") == 0 {
			if strings.Compare(grid[x+2][y+2], "A") == 0 {
				if strings.Compare(grid[x+3][y+3], "S") == 0 {
					count += 1
				}
			}
		}
	}

	if x >= minthresh && y >= minthresh {
		if strings.Compare(grid[x-1][y-1], "M") == 0 {
			if strings.Compare(grid[x-2][y-2], "A") == 0 {
				if strings.Compare(grid[x-3][y-3], "S") == 0 {
					count += 1
				}
			}
		}
	}

	return count
}

func diag2(grid [][]string, x int, y int) int {
	count := 0
	maxthresh := (len(grid)) - 4
	minthresh := (len(grid)) - maxthresh - 1
	if x <= maxthresh && y >= minthresh {
		if strings.Compare(grid[x+1][y-1], "M") == 0 {
			if strings.Compare(grid[x+2][y-2], "A") == 0 {
				if strings.Compare(grid[x+3][y-3], "S") == 0 {
					count += 1
				}
			}
		}
	}

	if x >= minthresh && y <= maxthresh {
		if strings.Compare(grid[x-1][y+1], "M") == 0 {
			if strings.Compare(grid[x-2][y+2], "A") == 0 {
				if strings.Compare(grid[x-3][y+3], "S") == 0 {
					count += 1
				}
			}
		}
	}

	return count
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
			if strings.Compare(grid[i][j], "X") == 0 {
				total += horizontal(grid, i, j)
				total += vertical(grid, i, j)
				total += diag1(grid, i, j)
				total += diag2(grid, i, j)
			} else {
				continue
			}
		}
	}
	fmt.Println(total)
}
