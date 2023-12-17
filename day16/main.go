package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
)

func main() {
	data :=
		`.|...\....
|.-.\.....
.....|-...
........|.
..........
.........\
..../.\\..
.-.-/..|..
.|....-|.\
..//.|....`
	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	f, _ := os.Open("day16/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	f, _ = os.Open("day16/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()
}

type Direction int

const (
	_ Direction = iota
	Up
	Right
	Down
	Left
)

type vector struct {
	x, y      int
	direction Direction
}

func (v vector) String() string {
	d := ""
	switch v.direction {
	case Up:
		d = "Up"
	case Down:
		d = "Down"
	case Left:
		d = "Left"
	case Right:
		d = "Right"
	}
	return fmt.Sprintf("(%d,%d,%s)", v.x, v.y, d)
}

func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	var grid [][]byte
	for s.Scan() {
		grid = append(grid, []byte(s.Text()))
	}
	curPos := vector{0, 0, Right}
	sum := calcEnergized(grid, curPos)
	return sum
}

func calcPart2(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	var grid [][]byte
	for s.Scan() {
		grid = append(grid, []byte(s.Text()))
	}
	maxSum := 0
	// top row
	for i := 0; i < len(grid[0]); i++ {
		for d := 1; d <= 4; d++ {
			curPos := vector{i, 0, Direction(d)}
			sum := calcEnergized(grid, curPos)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	// left col
	for i := 0; i < len(grid); i++ {
		for d := 1; d <= 4; d++ {
			curPos := vector{0, i, Direction(d)}
			sum := calcEnergized(grid, curPos)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	// right col
	for i := 0; i < len(grid); i++ {
		for d := 1; d <= 4; d++ {
			curPos := vector{len(grid[0]) - 1, i, Direction(d)}
			sum := calcEnergized(grid, curPos)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	// bottom row
	for i := 0; i < len(grid[0]); i++ {
		for d := 1; d <= 4; d++ {
			curPos := vector{i, len(grid) - 1, Direction(d)}
			sum := calcEnergized(grid, curPos)
			if sum > maxSum {
				maxSum = sum
			}
		}
	}
	return maxSum
}

func calcEnergized(grid [][]byte, curPos vector) int {
	var energized [][]bool
	energized = make([][]bool, len(grid))
	for i := range energized {
		energized[i] = make([]bool, len(grid[i]))
	}
	var wg sync.WaitGroup
	wg.Add(1)
	findEnergized(grid, energized, curPos, map[vector]bool{}, &wg)
	wg.Wait()
	sum := 0
	for _, v := range energized {
		for _, v2 := range v {
			if v2 {
				sum++
			}
		}
	}
	return sum
}

func findEnergized(grid [][]byte, energized [][]bool, pos vector, loops map[vector]bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for validPos(pos, len(grid), len(grid[0])) {
		if loops[pos] {
			return
		}
		loops[pos] = true
		energized[pos.y][pos.x] = true
		//fmt.Println(pos)
		//fmt.Println(string(grid[pos.y][pos.x]))
		//printEnergized(energized)
		switch grid[pos.y][pos.x] {
		case '|':
			switch pos.direction {
			case Up:
				// doesn't split, just keeps going up
				pos.y--
			case Down:
				// doesn't split, just keeps going up
				pos.y++
			case Left, Right:
				//splits up and down
				// recurse for down
				wg.Add(1)
				findEnergized(grid, energized, vector{
					x:         pos.x,
					y:         pos.y + 1,
					direction: Down,
				}, loops, wg)
				pos.y--
				pos.direction = Up
			}
		case '-':
			switch pos.direction {
			case Left:
				// doesn't split, just keeps going Left
				pos.x--
			case Right:
				// doesn't split, just keeps going Right
				pos.x++
			case Up, Down:
				//splits left and right
				// recurse for right
				wg.Add(1)
				findEnergized(grid, energized, vector{
					x:         pos.x + 1,
					y:         pos.y,
					direction: Right,
				}, loops, wg)
				pos.x--
				pos.direction = Left
			}
		case '.':
			// pass through
			switch pos.direction {
			case Left:
				// doesn't split, just keeps going Left
				pos.x--
			case Right:
				// doesn't split, just keeps going Right
				pos.x++
			case Up:
				pos.y--
			case Down:
				pos.y++
			}
		case '\\':
			switch pos.direction {
			case Left:
				// go up from here
				pos.y--
				pos.direction = Up
			case Right:
				// go down from here
				pos.y++
				pos.direction = Down
			case Up:
				// go left from here
				pos.x--
				pos.direction = Left
			case Down:
				// go right from here
				pos.x++
				pos.direction = Right
			}
		case '/':
			switch pos.direction {
			case Left:
				// go down from here
				pos.y++
				pos.direction = Down
			case Right:
				// go up from here
				pos.y--
				pos.direction = Up
			case Up:
				// go right from here
				pos.x++
				pos.direction = Right
			case Down:
				// go left from here
				pos.x--
				pos.direction = Left
			}
		}
	}
}

func validPos(pos vector, maxY int, maxX int) bool {
	return pos.x >= 0 && pos.x < maxX && pos.y >= 0 && pos.y < maxY
}

func printEnergized(energized [][]bool) {
	for _, v := range energized {
		for _, v2 := range v {
			if v2 {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}
