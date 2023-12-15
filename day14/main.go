package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
	"strings"
)

func main() {
	data :=
		`O....#....
O.OO#....#
.....##...
OO.#O....O
.O.....O#.
O.#..O.#.#
..O..#O..O
.......O..
#....###..
#OO..#....`
	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	f, _ := os.Open("day14/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day14/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()
}

func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	var rockMap [][]byte
	for s.Scan() {
		rockMap = append(rockMap, []byte(s.Text()))
	}
	maxWeight := len(rockMap)
	sum := 0
	for i, curRow := range rockMap {
		//see how far to the top a O can go
		for x, v := range curRow {
			if v == '.' || v == '#' {
				continue
			}
			maxPos := i
			numRocks := 0
			for j := i - 1; j >= 0; j-- {
				if rockMap[j][x] == '#' {
					break
				}
				if rockMap[j][x] == 'O' {
					numRocks++
				}
				if rockMap[j][x] == '.' {
					maxPos--
				}
			}
			val := maxWeight - maxPos
			//fmt.Printf("rock at (%d,%d) will roll to %d, %d rocks in front of it: %d\n", x, i, maxPos, numRocks, val)
			sum += val
		}
	}
	return sum
}

func calcPart2(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	var rockMap [][]byte
	for s.Scan() {
		rockMap = append(rockMap, []byte(s.Text()))
	}
	var scores [400]int
	for i := 0; i < 400; i++ {
		//if (i+1)%100_000 == 0 {
		//	fmt.Println(i, scores)
		//}
		//fmt.Println("north")
		from, to := rollNorth(rockMap)
		//fmt.Println("moving", len(from), "rocks")
		rockMap = moveRocks(rockMap, from, to)
		//printGrid(rockMap)
		//fmt.Println("west")
		from, to = rollWest(rockMap)
		//fmt.Println("moving", len(from), "rocks")
		rockMap = moveRocks(rockMap, from, to)
		//printGrid(rockMap)
		//fmt.Println("south")
		from, to = rollSouth(rockMap)
		//fmt.Println("moving", len(from), "rocks")
		rockMap = moveRocks(rockMap, from, to)
		//printGrid(rockMap)
		//fmt.Println("east")
		from, to = rollEast(rockMap)
		//fmt.Println("moving", len(from), "rocks")
		rockMap = moveRocks(rockMap, from, to)
		//printGrid(rockMap)
		scores[i%len(scores)] = calcScore(rockMap)
	}
	offset, cycle := findCycleAndOffset(scores)

	return cycle[(1_000_000_000-offset-1)%len(cycle)]
}

func findCycleAndOffset(scores [400]int) (int, []int) {
	// look for 4 numbers in a row, should be good enough
	for i := 0; i < len(scores)-4; i += 1 {
		nums := scores[i : i+4]
		for j := i + 4; j < len(scores)-4; j++ {
			if slices.Equal(nums, scores[j:j+4]) {
				return i, scores[i:j]
			}
		}
	}
	panic("failed")
}

func rollNorth(rockMap [][]byte) ([]point, []point) {
	var from []point
	var to []point
	for i, curRow := range rockMap {
		//see how far to the top a O can go
		for x, v := range curRow {
			if v == '.' || v == '#' {
				continue
			}
			maxPos := i
			numRocks := 0
			for j := i - 1; j >= 0; j-- {
				if rockMap[j][x] == '#' {
					break
				}
				if rockMap[j][x] == 'O' {
					numRocks++
				}
				if rockMap[j][x] == '.' {
					maxPos--
				}
			}
			//fmt.Printf("rock at (%d,%d) will roll to %d, %d rocks in front of it\n", x, i, maxPos, numRocks)
			from = append(from, point{x, i})
			to = append(to, point{x, maxPos})
		}
	}
	return from, to
}

func rollSouth(rockMap [][]byte) ([]point, []point) {
	var from []point
	var to []point
	for i := len(rockMap) - 1; i >= 0; i-- {
		curRow := rockMap[i]
		//see how far to the top a O can go
		for x, v := range curRow {
			if v == '.' || v == '#' {
				continue
			}
			maxPos := i
			numRocks := 0
			for j := i + 1; j < len(rockMap); j++ {
				if rockMap[j][x] == '#' {
					break
				}
				if rockMap[j][x] == 'O' {
					numRocks++
				}
				if rockMap[j][x] == '.' {
					maxPos++
				}
			}
			//fmt.Printf("rock at (%d,%d) will roll to %d, %d rocks in front of it\n", x, i, maxPos, numRocks)
			from = append(from, point{x, i})
			to = append(to, point{x, maxPos})
		}
	}
	return from, to
}

func rollWest(rockMap [][]byte) ([]point, []point) {
	var from []point
	var to []point
	for x := 0; x < len(rockMap[0]); x++ {
		//see how far to the left a O can go
		for y, v := range rockMap {
			if v[x] == '.' || v[x] == '#' {
				continue
			}
			maxPos := x
			numRocks := 0
			for j := x - 1; j >= 0; j-- {
				if rockMap[y][j] == '#' {
					break
				}
				if rockMap[y][j] == 'O' {
					numRocks++
				}
				if rockMap[y][j] == '.' {
					maxPos--
				}
			}
			//fmt.Printf("rock at (%d,%d) will roll to %d, %d rocks in front of it\n", x, y, maxPos, numRocks)
			from = append(from, point{x, y})
			to = append(to, point{maxPos, y})
		}
	}
	return from, to
}

func rollEast(rockMap [][]byte) ([]point, []point) {
	var from []point
	var to []point
	for x := len(rockMap[0]) - 1; x >= 0; x-- {
		//see how far to the top a O can go
		for y, v := range rockMap {
			if v[x] == '.' || v[x] == '#' {
				continue
			}
			maxPos := x
			numRocks := 0
			for j := x + 1; j < len(rockMap[0]); j++ {
				if rockMap[y][j] == '#' {
					break
				}
				if rockMap[y][j] == 'O' {
					numRocks++
				}
				if rockMap[y][j] == '.' {
					maxPos++
				}
			}
			//fmt.Printf("rock at (%d,%d) will roll to %d, %d rocks in front of it\n", x, y, maxPos, numRocks)
			from = append(from, point{x, y})
			to = append(to, point{maxPos, y})
		}
	}
	return from, to
}

type point struct {
	x, y int
}

func moveRocks(grid [][]byte, from []point, to []point) [][]byte {
	moved := map[point]bool{}
	for i, v := range from {
		if !moved[v] {
			grid[v.y][v.x] = '.'
			moved[v] = true
		}
		grid[to[i].y][to[i].x] = 'O'
	}
	return grid
}

func calcScore(grid [][]byte) int {
	sum := 0
	maxWeight := len(grid)
	for i, curRow := range grid {
		for _, v := range curRow {
			if v == 'O' {
				sum += maxWeight - i
			}
		}
	}
	return sum
}

func printGrid(grid [][]byte) {
	var out string
	for _, v := range grid {
		out = out + string(v) + "\n"
	}
	fmt.Println(out)
}
