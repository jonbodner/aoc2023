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
		`Time:      7  15   30
Distance:  9  40  200`

	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	data2 :=
		`Time:      71530
Distance:  940200`

	dr = strings.NewReader(data2)
	sum = calcPart1(dr)
	fmt.Println(sum)

	f, _ := os.Open("day06/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day06/input2.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()
}

func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	s.Scan()
	timeLine := s.Text()
	_, timeFields, _ := strings.Cut(timeLine, ":")
	times := stringToIntSlice(timeFields)
	s.Scan()
	distanceLine := s.Text()
	_, distanceFields, _ := strings.Cut(distanceLine, ":")
	distances := stringToIntSlice(distanceFields)

	sum := 1
	for i := 0; i < len(times); i++ {
		total := 0
		curTime := times[i]
		curDistance := distances[i]
		for j := curTime / 2; j > 0; j-- {
			distance := (curTime - j) * j
			if distance > curDistance {
				total += 2
			} else {
				break
			}
		}
		if curTime%2 == 0 {
			total--
		}
		sum *= total
	}
	return sum
}

func stringToIntSlice(s string) []int {
	vals := strings.Fields(s)
	valNums := make([]int, len(vals))
	for i := 0; i < len(vals); i++ {
		valNums[i], _ = strconv.Atoi(vals[i])
	}
	return valNums
}
