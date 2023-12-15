package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	data :=
		`???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1`
	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	f, _ := os.Open("day12/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day12/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()
}

func calcPart2(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	sum := 0
	for s.Scan() {
		curLine := parseLine(s.Text())
		fmt.Println(curLine)
		multiLine := strings.Repeat(curLine.pattern+"?", 5)
		multiLine = multiLine[:len(multiLine)-1]
		multiNums := make([]int, len(curLine.nums)*5)
		for i := 0; i < 5; i++ {
			copy(multiNums[i*len(curLine.nums):], curLine.nums)
		}
		fmt.Println(multiLine)
		fmt.Println(multiNums)
		count := findOptions(line{
			pattern: multiLine,
			nums:    multiNums,
		}, false)

		fmt.Println(count)
		sum += count

	}
	return sum
}

type line struct {
	pattern string
	nums    []int
}

func parseLine(curLine string) line {
	// split on space
	pattern, numbers, _ := strings.Cut(curLine, " ")
	return line{
		pattern: pattern,
		nums:    stringToIntSlice(numbers),
	}
}

func stringToIntSlice(s string) []int {
	vals := strings.Split(s, ",")
	valNums := make([]int, len(vals))
	for i := 0; i < len(vals); i++ {
		valNums[i], _ = strconv.Atoi(vals[i])
	}
	return valNums
}

func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	sum := 0
	for s.Scan() {
		curLine := parseLine(s.Text())
		fmt.Println(curLine)
		count := findOptions(curLine, false)
		fmt.Println(count)
		sum += count

	}
	return sum
}

func findOptions(curLine line, inNum bool) int {
	if len(curLine.nums) == 0 {
		if strings.ContainsRune(curLine.pattern, '#') {
			return 0
		}
		return 1
	}
	if len(curLine.pattern) == 0 {
		return 0
	}
	// take the first number, see if we can match it
	curNum := curLine.nums[0]
	var lastHash bool
	switch curLine.pattern[0] {
	case '.':
		if inNum {
			return 0
		}
		lastHash = false
	case '#':
		inNum = true
		curNum--
		lastHash = true
	case '?':
		// if !inNum, could go either way
		if !inNum {
			var total int
			total += findOptions(line{
				pattern: "." + curLine.pattern[1:],
				nums:    append([]int{curNum}, curLine.nums[1:]...),
			}, inNum)
			total += findOptions(line{
				pattern: "#" + curLine.pattern[1:],
				nums:    append([]int{curNum}, curLine.nums[1:]...),
			}, inNum)
			return total
		} else {
			return findOptions(line{
				pattern: "#" + curLine.pattern[1:],
				nums:    append([]int{curNum}, curLine.nums[1:]...),
			}, inNum)
		}
	}
	if curNum != 0 {
		pattern := curLine.pattern[1:]
		if lastHash && len(pattern) > 0 && pattern[0] == '?' {
			pattern = "#" + pattern[1:]
		}
		return findOptions(line{
			pattern: pattern,
			nums:    append([]int{curNum}, curLine.nums[1:]...),
		}, inNum)
	}
	pattern := curLine.pattern[1:]
	if lastHash && len(pattern) > 0 && pattern[0] == '?' {
		pattern = "." + pattern[1:]
	}

	// if the next character is a #, it's also invalid
	if len(pattern) > 0 && pattern[0] == '#' {
		return 0
	}
	return findOptions(line{
		pattern: pattern,
		nums:    curLine.nums[1:],
	}, false)
}
