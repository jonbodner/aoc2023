package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	sum := 0
	data :=
		`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..`
	dr := strings.NewReader(data)
	sum = calcPart1(dr)
	fmt.Println(sum)

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	f, _ := os.Open("day03/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day03/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()
}

/*
The engineer explains that an engine part seems to be missing from the engine, but nobody can figure out which one. If you can add up all the part numbers in the engine schematic, it should be easy to work out which part is missing.

The engine schematic (your puzzle input) consists of a visual representation of the engine. There are lots of numbers and symbols you don't really understand, but apparently any number adjacent to a symbol, even diagonally, is a "part number" and should be included in your sum. (Periods (.) do not count as a symbol.)

Here is an example engine schematic:

467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
In this schematic, two numbers are not part numbers because they are not adjacent to a symbol: 114 (top right) and 58 (middle right). Every other number is adjacent to a symbol and so is a part number; their sum is 4361.

Of course, the actual engine schematic is much larger. What is the sum of all of the part numbers in the engine schematic?
*/
type pos struct {
	x, y int
}

type numInfo struct {
	num        int
	start, end pos
}

func calcPart1(dr io.Reader) int {
	var nums []numInfo
	symbols := map[pos]rune{}

	// find all numbers and symbols
	nums, symbols = buildData(dr)
	//fmt.Println(nums)
	//fmt.Println(symbols)
	// scan to see what numbers are near symbols
	sum := 0
	for _, v := range nums {
		found := false
	outer:
		for x := v.start.x - 1; x <= v.end.x+1; x++ {
			for y := v.start.y - 1; y <= v.end.y+1; y++ {
				p := pos{x, y}
				if val, ok := symbols[p]; ok {
					fmt.Println("adding:", v, p, string(val))
					sum += v.num
					found = true
					break outer
				}
			}
		}
		if !found {
			fmt.Println("skipping:", v.num)
		}
	}
	return sum
}

func calcPart2(dr io.Reader) int {
	var nums []numInfo
	symbols := map[pos]rune{}

	// find all numbers and symbols
	nums, symbols = buildData(dr)
	fmt.Println(nums)
	fmt.Println(symbols)
	// scan to see what numbers are near symbols
	gears := map[pos][]int{}

	for _, v := range nums {
		found := false
	outer:
		for x := v.start.x - 1; x <= v.end.x+1; x++ {
			for y := v.start.y - 1; y <= v.end.y+1; y++ {
				p := pos{x, y}
				if val, ok := symbols[p]; ok && val == '*' {
					gears[p] = append(gears[p], v.num)
					fmt.Println("gear:", p, "number", v.num)
					found = true
					break outer
				}
			}
		}
		if !found {
			fmt.Println("skipping:", v.num)
		}
	}
	sum := 0
	for _, v := range gears {
		if len(v) == 2 {
			sum += v[0] * v[1]
		}
	}
	return sum
}

func buildData(dr io.Reader) ([]numInfo, map[pos]rune) {
	var nums []numInfo
	symbols := map[pos]rune{}

	s := bufio.NewScanner(dr)
	row := 0
	for s.Scan() {
		curLine := s.Text()
		inNum := false
		curVal := 0
		startCol := -1
		for col, v := range curLine {
			switch v {
			case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
				if !inNum {
					inNum = true
					startCol = col
				}
				curVal = curVal*10 + int(v-'0')
			case '.':
				if inNum {
					nums = append(nums, numInfo{
						num:   curVal,
						start: pos{row, startCol},
						end:   pos{row, col - 1},
					})
					inNum = false
					curVal = 0
					startCol = -1
				}
			default:
				if inNum {
					nums = append(nums, numInfo{
						num:   curVal,
						start: pos{row, startCol},
						end:   pos{row, col - 1},
					})
					inNum = false
					curVal = 0
					startCol = -1
				}
				symbols[pos{row, col}] = v
			}
		}
		if inNum {
			nums = append(nums, numInfo{
				num:   curVal,
				start: pos{row, startCol},
				end:   pos{row, len(curLine) - 1},
			})
			inNum = false
			curVal = 0
			startCol = -1
		}
		row++
	}
	return nums, symbols
}
