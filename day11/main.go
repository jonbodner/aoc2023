package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	data :=
		`...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#.....`
	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	f, _ := os.Open("day11/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day11/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()
}

func calcPart1(dr io.Reader) int {
	return doWork(dr, 2)
}

func doWork(dr io.Reader, distance int) int {
	s := bufio.NewScanner(dr)
	grid := buildGrid(s)
	emptyRows, emptyCols := findEmptyRowsAndCols(grid)
	fmt.Println(emptyRows, emptyCols)
	points := findPoints(grid, emptyRows, emptyCols, distance)
	distances := calcDistances(points)
	sum := 0
	for _, v := range distances {
		sum += v
	}
	return sum
}

func calcPart2(dr io.Reader) int {
	return doWork(dr, 1_000_000)
}

func findPoints(grid [][]byte, emptyRows []int, emptyCols []int, distance int) []point {
	factor := distance - 1
	var out []point
	numEmptyRows := 0
	for y, row := range grid {
		if numEmptyRows < len(emptyRows) && emptyRows[numEmptyRows] == y {
			numEmptyRows++
			continue
		}
		numEmptyCols := 0
		for x, c := range row {
			if numEmptyCols < len(emptyCols) && emptyCols[numEmptyCols] == x {
				numEmptyCols++
				continue
			}
			if c == '#' {
				out = append(out, point{x: x + numEmptyCols*factor, y: y + numEmptyRows*factor})
			}
		}
	}
	return out
}

func findEmptyRowsAndCols(grid [][]byte) ([]int, []int) {
	var emptyRows []int
	for i, v := range grid {
		if bytes.IndexByte(v, '#') == -1 {
			emptyRows = append(emptyRows, i)
		}
	}
	// find all empty columns and write them twice
	var emptyCols []int
	for x := 0; x < len(grid[0]); x++ {
		found := false
		for y := 0; y < len(grid); y++ {
			if grid[y][x] == '#' {
				found = true
				break
			}
		}
		if !found {
			emptyCols = append(emptyCols, x)
		}
	}
	return emptyRows, emptyCols
}

func calcDistances(points []point) []int {
	var out []int
	for i := 0; i < len(points)-1; i++ {
		firstPoint := points[i]
		for j := i + 1; j < len(points); j++ {
			secondPoint := points[j]
			dist := int(math.Abs(float64(firstPoint.x-secondPoint.x)) + math.Abs(float64(firstPoint.y-secondPoint.y)))
			//fmt.Println(i+1, j+1, dist)
			out = append(out, dist)
		}
	}
	return out
}

type point struct {
	x, y int
}

func gridString(grid [][]byte) string {
	var out string
	for _, v := range grid {
		out += string(v) + "\n"
	}
	return out
}

func buildGrid(s *bufio.Scanner) [][]byte {
	var out [][]byte
	for s.Scan() {
		curRow := []byte(s.Text())
		out = append(out, curRow)
	}
	return out
}
