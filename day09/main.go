package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	data :=
		`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45`
	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	f, _ := os.Open("day09/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day09/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()
}

func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	sum := 0
	for s.Scan() {
		curRow := stringToIntSlice(s.Text())
		sum += calcNext(curRow)
	}
	return sum
}

func calcPart2(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	sum := 0
	for s.Scan() {
		curRow := stringToIntSlice(s.Text())
		slices.Reverse(curRow)
		sum += calcNext(curRow)
	}
	return sum
}

func calcNext(curRow []int) int {
	var allRows [][]int
	done := func(r []int) bool {
		for _, v := range r {
			if v != 0 {
				return false
			}
		}
		return true
	}
	for !done(curRow) {
		allRows = append(allRows, curRow)
		var nextRow []int
		for i := 0; i < len(curRow)-1; i++ {
			nextRow = append(nextRow, curRow[i+1]-curRow[i])
		}
		curRow = nextRow
	}
	for i := len(allRows) - 2; i >= 0; i-- {
		allRows[i] = append(allRows[i], allRows[i+1][len(allRows[i+1])-1]+allRows[i][len(allRows[i])-1])
	}
	topRow := allRows[0]
	return topRow[len(topRow)-1]
}

func stringToIntSlice(s string) []int {
	vals := strings.Fields(s)
	valNums := make([]int, len(vals))
	for i := 0; i < len(vals); i++ {
		valNums[i], _ = strconv.Atoi(vals[i])
	}
	return valNums
}
