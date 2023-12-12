package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

/*
.....
..F7.
.FJ|.
.|.|.
.L-J.
.....
*/
func main() {
	data :=
		`-L|F7
 7S-7|
 L|7||
 -L-J|
 L|-JF`
	//dr := strings.NewReader(data)
	//sum := calcPart1(dr)
	//fmt.Println(sum)

	dr := strings.NewReader(data)
	sum := calcPart2a(dr)
	fmt.Println(sum)

	data2 :=
		`7-F7-
.FJ|7
SJLL7
|F--J
LJ.LJ`
	//dr = strings.NewReader(data2)
	//sum = calcPart1(dr)
	//fmt.Println(sum)

	dr = strings.NewReader(data2)
	sum = calcPart2a(dr)
	fmt.Println(sum)

	data3 :=
		`FF7FSF7F7F7F7F7F---7
L|LJ||||||||||||F--J
FL-7LJLJ||||||LJL-77
F--JF--7||LJLJ7F7FJ-
L---JF-JLJ.||-FJLJJ7
|F|F-JF---7F7-L7L|7|
|FFJF7L7F-JF7|JL---7
7-L-JL7||F7|L7F-7F7|
L.L7LFJ|||||FJL7||LJ
L7JLJL-JLJLJL--JLJ.L`
	dr = strings.NewReader(data3)
	sum = calcPart2a(dr)
	fmt.Println(sum)

	data4 :=
		`...........
.S-------7.
.|F-----7|.
.||.....||.
.||.....||.
.|L-7.F-J|.
.|..|.|..|.
.L--J.L--J.
...........`
	dr = strings.NewReader(data4)
	sum = calcPart2a(dr)
	fmt.Println(sum)

	data5 :=
		`..........
.S------7.
.|F----7|.
.||OOOO||.
.||OOOO||.
.|L-7F-J|.
.|II||II|.
.L--JL--J.
..........`
	dr = strings.NewReader(data5)
	sum = calcPart2a(dr)
	fmt.Println(sum)

	//f, _ := os.Open("day10/input.txt")
	//sum = calcPart1(f)
	//fmt.Println(sum)
	//f.Close()

	//f, _ := os.Open("day10/input.txt")
	//sum = calcPart2a(f)
	//fmt.Println(sum)
	//f.Close()

}

/*
The pipes are arranged in a two-dimensional grid of tiles:

| is a vertical pipe connecting north and south.
- is a horizontal pipe connecting east and west.
L is a 90-degree bend connecting north and east.
J is a 90-degree bend connecting north and west.
7 is a 90-degree bend connecting south and west.
F is a 90-degree bend connecting south and east.
. is ground; there is no pipe in this tile.
S is the starting position of the animal; there is a pipe on this tile, but your sketch doesn't show what shape the pipe has.
Based on the acoustics of the animal's scurrying, you're confident the pipe that contains the animal is one large, continuous loop.

You can count the distance each tile in the loop is from the starting point like this:

.....
.012.
.1.3.
.234.
.....
In this example, the farthest point from the start is 4 steps away.

Here's the more complex loop again:

..F7.
.FJ|.
SJ.L7
|F--J
LJ...
Here are the distances for each tile on that loop:

..45.
.236.
01.78
14567
23...
Find the single giant loop starting at S. How many steps along the loop does it take to get from the starting position to the point farthest from the starting position?
*/

type point struct {
	x, y int
}

func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	grid, sPoint := findGridAndStart(s)
	if sPoint.x == -1 {
		fmt.Println("broken")
		return 0
	}
	//look for the complete loop around S
	startPoints := buildStartPoints(sPoint, grid)

	fmt.Println(startPoints)
	for _, sp := range startPoints {
		pathPoints := findPathFrom(sPoint, sp, grid)
		if len(pathPoints) > 0 {
			fmt.Println(pathPoints)
			return len(pathPoints) / 2
		}
	}
	return -1
}

func calcPart2a(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	grid, sPoint := findGridAndStart(s)
	if sPoint.x == -1 {
		fmt.Println("broken")
		return 0
	}
	//look for the complete loop around S
	startPoints := buildStartPoints(sPoint, grid)

	fmt.Println(startPoints)
	var pathPoints []point
	for _, sp := range startPoints {
		pathPoints = findPathFrom(sPoint, sp, grid)
		if len(pathPoints) > 0 {
			break
		}
	}
	if len(pathPoints) == 0 {
		return -1
	}
	pointSet := map[point]bool{}
	for _, v := range pathPoints {
		pointSet[v] = true
	}

	fixS(sPoint, pathPoints, grid)

	//scan to see if a space is in the interior or exterior
	count := 0
	for y, row := range grid {
		in := false
		for x, curPipe := range row {
			curPoint := point{x: x, y: y}
			if pointSet[curPoint] {
				// are you a |, J, F, L, 7?
				// if so, toggle
				if strings.ContainsRune("|JFL7", rune(curPipe)) {
					// if you're a pipe, toggle
					// if the previous character is a compliment, don't toggle
					if x == 0 {
						in = true
					} else {
						//if curPipe == '|' {
						//	in = !in
						//}
						if curPipe == '7' || curPipe == 'J' {
							lastSpace := row[x-1]
							if lastSpace == 'F' || lastSpace == 'L' {
								// do nothing
							} else {
								in = !in
							}
						} else {
							in = !in
						}
					}
				}
			} else {
				if in {
					count++
					grid[y][x] = 'I'
				} else {
					grid[y][x] = ' '
				}
			}
		}
	}
	printGrid(grid)
	return count
}

func calcPart2(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	grid, sPoint := findGridAndStart(s)
	if sPoint.x == -1 {
		fmt.Println("broken")
		return 0
	}
	//look for the complete loop around S
	startPoints := buildStartPoints(sPoint, grid)

	fmt.Println(startPoints)
	var pathPoints []point
	for _, sp := range startPoints {
		pathPoints = findPathFrom(sPoint, sp, grid)
		if len(pathPoints) > 0 {
			break
		}
	}
	if len(pathPoints) == 0 {
		return -1
	}

	// figure out if something is in or out based on the last angle:
	/*
		L - anything north or east is in
		J - anything north of west is in
		F - anything south or east is in
		7 - anything south or west is in

		first figure out the shape of S and replace
		then find the first element in the path that is not - or |
		then use that to label everything directly next to it as in or out
		then when done, take everything between the ins and mark them as in
		and then count the ins
	*/
	fixS(sPoint, pathPoints, grid)
	// find the first corner in point list
	var firstCorner int
	for i, v := range pathPoints {
		if bytes.IndexByte([]byte("LF7J"), grid[v.y][v.x]) != -1 {
			firstCorner = i
			break
		}
	}
	doNorth := false
	doEast := false
	changeDir := func(b byte) {
		switch b {
		//north and east
		case 'L':
			doNorth = true
			doEast = true
		// south and east
		case 'F':
			doNorth = false
			doEast = true
		// south and west
		case '7':
			doNorth = false
			doEast = false
		// north and west
		case 'J':
			doNorth = true
			doEast = false
		}
	}
	//inMap := map[byte][2]delta{
	//	//north and east
	//	'L': {{y: -1}, {x: 1}},
	//	// south and east
	//	'F': {{y: 1}, {x: 1}},
	//	// south and west
	//	'7': {{y: 1}, {x: -1}},
	//	// north and west
	//	'J': {{y: -1}, {x: -1}},
	//}
	//var curInMap [2]delta
	pointSet := map[point]bool{}
	for _, v := range pathPoints {
		pointSet[v] = true
	}
	for i := 0; i < len(pathPoints); i++ {
		curPoint := pathPoints[(i+firstCorner)%len(pathPoints)]
		curPipe := grid[curPoint.y][curPoint.x]
		changeDir(curPipe)
		switch curPipe {
		case '|':
			// mark east or west if not in the pipe
			if doEast {
				updateIfOk(grid, pointSet, point{x: curPoint.x + 1, y: curPoint.y})
			} else {
				updateIfOk(grid, pointSet, point{x: curPoint.x - 1, y: curPoint.y})
			}
		case '-':
			// mark north or south if not in the pipe
			if doNorth {
				updateIfOk(grid, pointSet, point{x: curPoint.x, y: curPoint.y - 1})
			} else {
				updateIfOk(grid, pointSet, point{x: curPoint.x, y: curPoint.y + 1})
			}
		case 'L':
			// mark northeast if not in the pipe
			updateIfOk(grid, pointSet, point{x: curPoint.x + 1, y: curPoint.y - 1})
		case 'J':
			// mark northwest if not in the pipe
			updateIfOk(grid, pointSet, point{x: curPoint.x - 1, y: curPoint.y - 1})
		case 'F':
			// mark southeast if not in the pipe
			updateIfOk(grid, pointSet, point{x: curPoint.x + 1, y: curPoint.y + 1})
		case '7':
			// mark southwest if not in the pipe
			updateIfOk(grid, pointSet, point{x: curPoint.x - 1, y: curPoint.y - 1})
		}
	}
	printGrid(grid)
	return 0
}

func fixS(sPoint point, pathPoints []point, grid [][]byte) {
	// figure out S shape (last element in pathpoints)
	east := point{x: sPoint.x + 1, y: sPoint.y}
	west := point{x: sPoint.x - 1, y: sPoint.y}
	south := point{x: sPoint.x, y: sPoint.y + 1}
	north := point{x: sPoint.x, y: sPoint.y - 1}

	prevPoint := pathPoints[len(pathPoints)-2]
	var possible string
	switch prevPoint {
	case west:
		possible = "-J7"
	case east:
		possible = "-FL"
	case north:
		possible = "|LJ"
	case south:
		possible = "|F7"
	}
	nextPoint := pathPoints[0]
	var intersect string
	switch nextPoint {
	case west:
		intersect = "-J7"
	case east:
		intersect = "-FL"
	case north:
		intersect = "|LJ"
	case south:
		intersect = "|F7"
	}
	var actual byte
	for _, v := range intersect {
		if strings.ContainsRune(possible, v) {
			actual = byte(v)
		}
	}
	fmt.Println(string(actual))
	grid[sPoint.y][sPoint.x] = actual
}

func updateIfOk(grid [][]byte, pointSet map[point]bool, p point) {
	// don't update if the point is on the pipe
	if pointSet[p] {
		return
	}
	// don't update if it's off the grid
	if p.x < 0 || p.y < 0 || p.x > len(grid[0]) || p.y > len(grid) {
		return
	}
	grid[p.y][p.x] = 'I'
}

func findGridAndStart(s *bufio.Scanner) ([][]byte, point) {
	var grid [][]byte
	// find S while loading
	sPoint := point{-1, -1}
	curY := 0
	for s.Scan() {
		curRow := []byte(s.Text())
		grid = append(grid, curRow)
		if x := bytes.IndexByte(curRow, 'S'); x != -1 {
			sPoint.x = x
			sPoint.y = curY
		}
		curY++
	}
	return grid, sPoint
}

type delta struct {
	x, y int
}

/*
north: -1 y
south: +1 y
east: +1 x
west: -1 x
*/
var dirMap = map[byte][2]delta{
	// north/south
	'|': {{y: -1}, {y: 1}},
	// east/west
	'-': {{x: 1}, {x: -1}},
	// north/east
	'L': {{y: -1}, {x: 1}},
	// north/west
	'J': {{y: -1}, {x: -1}},
	// south/west
	'7': {{y: 1}, {x: -1}},
	// south/east
	'F': {{y: 1}, {x: 1}},
}

func findPathFrom(startPoint, firstPoint point, grid [][]byte) []point {
	lastPoint := startPoint
	curPoint := firstPoint
	var pathPoints []point
	fmt.Println("starting from ", curPoint)
	for {
		pathPoints = append(pathPoints, curPoint)
		curState := grid[curPoint.y][curPoint.x]
		switch curState {
		case 'S':
			fmt.Println("reached S again!", curPoint)
			return pathPoints
		case '.':
			fmt.Println("reached end, not the loop")
			return nil
		default:
			moves := dirMap[curState]
			hasNext := false
			for _, v := range moves {
				newPoint := point{x: curPoint.x + v.x, y: curPoint.y + v.y}
				//fmt.Println("potential new point", newPoint)

				if newPoint.x >= 0 && newPoint.y >= 0 && newPoint.x < len(grid[0]) && newPoint.y < len(grid) && newPoint != lastPoint {
					//fmt.Println("good point!")
					hasNext = true
					lastPoint = curPoint
					curPoint = newPoint
					break
				}
			}
			if !hasNext {
				fmt.Println("reached end, not the loop")
				return nil
			}
		}
	}

}

func buildStartPoints(sPoint point, grid [][]byte) []point {
	var startPoints []point
	east := point{x: sPoint.x + 1, y: sPoint.y}
	west := point{x: sPoint.x - 1, y: sPoint.y}
	south := point{x: sPoint.x, y: sPoint.y + 1}
	north := point{x: sPoint.x, y: sPoint.y - 1}
	potentialStartPoints := []point{west, east, south, north}
	for _, psp := range potentialStartPoints {
		if psp.x < 0 || psp.y < 0 || psp.x >= len(grid[0]) || psp.y >= len(grid) {
			fmt.Println("skipping start point", psp)
			continue
		}
		// can't go y+1 (north) or y-1 (south) to a -
		// can't go x+1 (east) or x-1 (west) to a |
		// can't go x-1 (west) or y+1 (south) to a 7
		// can't go x+1 (east) or y-1 (north) to an L
		// can't go x-1 (west) or y-1 (north)to a J
		// can't go x+1 (east) or y+1 (south) to an F
		pspVal := grid[psp.y][psp.x]
		switch pspVal {
		case '.':
			fmt.Println("not a pipe")
		case '-':
			if psp != north && psp != south {
				startPoints = append(startPoints, psp)
			} else {
				fmt.Println("can't go north or south to a -")
			}
		case '|':
			if psp != east && psp != west {
				startPoints = append(startPoints, psp)
			} else {
				fmt.Println("can't go east or west to a |")
			}
		case '7':
			if psp != west && psp != south {
				startPoints = append(startPoints, psp)
			} else {
				fmt.Println("can't go west or south to a 7")
			}
		case 'L':
			if psp != east && psp != north {
				startPoints = append(startPoints, psp)
			} else {
				fmt.Println("can't go east or north to a L")
			}
		case 'J':
			if psp != west && psp != north {
				startPoints = append(startPoints, psp)
			} else {
				fmt.Println("can't go west or north to a J")
			}
		case 'F':
			if psp != east && psp != south {
				startPoints = append(startPoints, psp)
			} else {
				fmt.Println("can't go east or south to a F")
			}
		}
	}
	return startPoints
}

func printRow(g []byte) {
	var out string
	for _, v := range g {
		out = out + string(v)
	}
	fmt.Println(out)
}

func printGrid(g [][]byte) {
	var out string
	for _, v := range g {
		for _, v2 := range v {
			out = out + string(v2)
		}
		out = out + "\n"
	}
	fmt.Println(out)
}
