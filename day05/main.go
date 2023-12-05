package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data :=
		`seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4`

	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	f, _ := os.Open("day05/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day05/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()
}

func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	s.Scan()
	_, seeds, _ := strings.Cut(s.Text(), ":")
	seedNums := stringToIntSlice(seeds)
	states := calcStates(s)
	//fmt.Println(states)
	minLocation := math.MaxInt
	for _, v := range seedNums {
		curVal := v
		for _, curMapper := range states {
			curVal = curMapper.findMap(curVal)
		}
		if curVal < minLocation {
			minLocation = curVal
		}
	}
	return minLocation
}

func calcPart2(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	s.Scan()
	_, seeds, _ := strings.Cut(s.Text(), ":")
	seedNums := stringToIntSlice(seeds)
	states := calcStates(s)
	ch := make(chan int, len(seedNums)/2)
	for i := 0; i < len(seedNums); i += 2 {
		i := i
		go func() {
			minLocation := math.MaxInt
			for j := seedNums[i]; j < seedNums[i]+seedNums[i+1]; j++ {
				curVal := j
				for _, curMapper := range states {
					curVal = curMapper.findMap(curVal)
				}
				if curVal < minLocation {
					minLocation = curVal
				}
			}
			ch <- minLocation
		}()
	}
	minLocation := math.MaxInt
	for i := 0; i < len(seedNums)/2; i++ {
		curV := <-ch
		if curV < minLocation {
			minLocation = curV
		}
	}
	return minLocation
}

func calcStates(s *bufio.Scanner) []mapper {
	var states []mapper
	var curMapper mapper
	for s.Scan() {
		curLine := s.Text()
		switch {
		case len(curLine) == 0:
			if curMapper != nil {
				states = append(states, curMapper)
			}
			curMapper = mapper{}
		case strings.Index(curLine, ":") != -1:
			// skip
		default:
			curMapper = append(curMapper, makeMapRecord(curLine))
		}
	}
	if curMapper != nil {
		states = append(states, curMapper)
	}
	return states
}

type mapRecord struct {
	to       int
	from     int
	valRange int
}

func makeMapRecord(s string) mapRecord {
	valNums := stringToIntSlice(s)
	return mapRecord{
		to:       valNums[0],
		from:     valNums[1],
		valRange: valNums[2],
	}
}

func stringToIntSlice(s string) []int {
	vals := strings.Fields(s)
	valNums := make([]int, len(vals))
	for i := 0; i < len(vals); i++ {
		valNums[i], _ = strconv.Atoi(vals[i])
	}
	return valNums
}

type mapper []mapRecord

func (m mapper) findMap(i int) int {
	for _, v := range m {
		if i >= v.from && i < v.from+v.valRange {
			return v.to + i - v.from
		}
	}
	return i
}
