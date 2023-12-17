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
		`rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7`
	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	//f, _ := os.Open("day15/input.txt")
	//sum = calcPart1(f)
	//fmt.Println(sum)
	//f.Close()

	f, _ := os.Open("day15/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()

}

func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	sum := 0
	for s.Scan() {
		curLine := s.Text()
		parts := strings.Split(curLine, ",")
		for _, part := range parts {
			curHash := hash(part)
			fmt.Println(curHash)
			sum += curHash
		}
	}
	return sum
}

func calcPart2(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	type box struct {
		label string
		value int
	}
	var boxes [256][]box
	for s.Scan() {
		curLine := s.Text()
		parts := strings.Split(curLine, ",")
		for _, part := range parts {
			if strings.ContainsRune(part, '-') {
				label, _, _ := strings.Cut(part, "-")
				pos := hash(label)
				loc := -1
				for i, v := range boxes[pos] {
					if v.label == label {
						loc = i
						break
					}
				}
				if loc != -1 {
					boxes[pos] = slices.Delete(boxes[pos], loc, loc+1)
				}
			}
			if strings.ContainsRune(part, '=') {
				label, valStr, _ := strings.Cut(part, "=")
				val, _ := strconv.Atoi(valStr)
				b := box{
					label: label,
					value: val,
				}
				pos := hash(label)
				loc := -1
				for i, v := range boxes[pos] {
					if v.label == label {
						loc = i
						break
					}
				}
				if loc == -1 {
					boxes[pos] = append(boxes[pos], b)
				} else {
					boxes[pos][loc] = b
				}
			}
		}
	}
	/*
		One plus the box number of the lens in question.
		The slot number of the lens within the box: 1 for the first lens, 2 for the second lens, and so on.
		The focal length of the lens.
		At the end of the above example, the focusing power of each lens is as follows:

		rn: 1 (box 0) * 1 (first slot) * 1 (focal length) = 1
		cm: 1 (box 0) * 2 (second slot) * 2 (focal length) = 4
		ot: 4 (box 3) * 1 (first slot) * 7 (focal length) = 28
		ab: 4 (box 3) * 2 (second slot) * 5 (focal length) = 40
		pc: 4 (box 3) * 3 (third slot) * 6 (focal length) = 72
		So, the above example ends up with a total focusing power of 145.
	*/
	for i, boxRow := range boxes {
		if len(boxRow) == 0 {
			continue
		}
		fmt.Print(i, ": ")
		for _, v := range boxRow {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
	sum := 0
	for i, boxRow := range boxes {
		for i2, v := range boxRow {
			curBox := (i + 1) * (i2 + 1) * v.value
			fmt.Println(i, v.label, curBox)
			sum += curBox
		}
	}
	return sum
}

func hash(in string) int {
	curHash := 0
	for _, v := range in {
		curHash += int(v)
		curHash *= 17
		curHash %= 256
	}
	return curHash
}
