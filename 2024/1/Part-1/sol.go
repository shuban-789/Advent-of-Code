package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
    "strconv"
    "sort"
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
    sort.Ints(x)
    sort.Ints(y)
    for i := 0; i < len(x); i++ {
        ret += abs(x[i] - y[i])
    }
    fmt.Println(ret)
}
