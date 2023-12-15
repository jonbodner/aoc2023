package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	data :=
		`#.##..##.
..#.##.#.
##......#
##......#
..#.##.#.
..##..##.
#.#.##.#.

#...##..#
#....#..#
..##..###
#####.##.
#####.##.
..##..###
#....#..#`
	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	f, _ := os.Open("day13/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day13/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()

}

func calcPart2(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	sum := 0
	for s.Scan() {
		nextGrid := buildGrid(s)
		printGrid(nextGrid)
		axis, pos := findReflect(nextGrid, checkNums2)
		fmt.Println(axis, pos)
		if axis == horizontal {
			sum += 100 * pos
		} else {
			sum += pos
		}
	}
	return sum
}

type orient int

const (
	_ orient = iota
	horizontal
	vertical
)

func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	sum := 0
	for s.Scan() {
		nextGrid := buildGrid(s)
		printGrid(nextGrid)
		axis, pos := findReflect(nextGrid, checkNums)
		fmt.Println(axis, pos)
		if axis == horizontal {
			sum += 100 * pos
		} else {
			sum += pos
		}
	}
	return sum
}

func findReflect(grid [][]byte, cn func([]int) int) (orient, int) {
	// check vertical because it's easier
	allNums := make([]int, len(grid))
	for i, v := range grid {
		allNums[i] = toNum(v)
	}
	fmt.Println(allNums)
	pos := cn(allNums)
	if pos != -1 {
		return horizontal, pos
	}

	allVNums := make([]int, len(grid[0]))
	for i := 0; i < len(grid[0]); i++ {
		curCol := make([]byte, len(grid))
		for j := 0; j < len(grid); j++ {
			curCol[j] = grid[j][i]
		}
		allVNums[i] = toNum(curCol)
	}
	fmt.Println(allVNums)
	pos = cn(allVNums)
	if pos == -1 {
		panic("nope")
	}
	return vertical, pos
}

func checkNums(allNums []int) int {
	// see if there are identical numbers
	for i := 0; i < len(allNums)-1; i++ {
		if allNums[i] == allNums[i+1] {
			fmt.Println("potential:", i, len(allNums))
			// see if it extends to an edge
			// left edge
			if i < len(allNums)/2 {
				good := true
				for j := 1; j <= i; j++ {
					if allNums[i-j] != allNums[i+1+j] {
						good = false
						break
					}
				}
				if good {
					fmt.Println("match!", i+1)
					return i + 1
				}
			} else {
				// right edge
				good := true
				for j := 1; i+j < len(allNums)-1; j++ {
					if allNums[i-j] != allNums[i+1+j] {
						good = false
						break
					}
				}
				if good {
					fmt.Println("match!", i+1)
					return i + 1
				}
			}
		}
	}
	return -1
}

func NumOfSetBits(n int) int {
	count := 0
	for n != 0 {
		count += n & 1
		n >>= 1
	}
	return count
}

func checkNums2(allNums []int) int {
	// see if there are identical numbers
	for i := 0; i < len(allNums)-1; i++ {
		comp := allNums[i] ^ allNums[i+1]
		numSmudge := NumOfSetBits(comp)
		if numSmudge < 2 {
			fmt.Println("potential:", i, len(allNums))
			// see if it extends to an edge
			// left edge
			if i < len(allNums)/2 {
				good := true
				for j := 1; j <= i; j++ {
					result := allNums[i-j] ^ allNums[i+1+j]
					fmt.Println(allNums[i-j], allNums[i+1+j], result, NumOfSetBits(result), NumOfSetBits(result) == 1)
					numSmudge += NumOfSetBits(result)
					if numSmudge > 1 {
						// check to see if only one bit set
						fmt.Println("fail")
						good = false
						break
					}
				}
				if good {
					fmt.Println("match!", i+1)
					return i + 1
				}
			} else {
				// right edge
				good := true
				for j := 1; i+j < len(allNums)-1; j++ {
					result := allNums[i-j] ^ allNums[i+1+j]
					fmt.Println(allNums[i-j], allNums[i+1+j], result, NumOfSetBits(result), NumOfSetBits(result) == 1)
					numSmudge += NumOfSetBits(result)
					if numSmudge > 1 {
						fmt.Println("fail")
						// check to see if only one bit set
						good = false
						break
					}
				}
				if good {
					fmt.Println("match!", i+1)
					return i + 1
				}
			}
		}
	}
	return -1
}

func toNum(b []byte) int {
	out := 0
	for _, v := range b {
		out = out << 1
		if v == '#' {
			out++
		}
	}
	return out
}

func buildGrid(s *bufio.Scanner) [][]byte {
	var out [][]byte
	out = append(out, []byte(s.Text()))
	for s.Scan() {
		curRow := s.Text()
		if curRow == "" {
			return out
		}
		out = append(out, []byte(curRow))
	}
	return out
}

func printGrid(grid [][]byte) {
	var out string
	for _, v := range grid {
		out = out + string(v) + "\n"
	}
	fmt.Println(out)
}
