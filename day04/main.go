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
		`Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11`
	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	f, _ := os.Open("day04/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day04/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()
}

/*
The Elf leads you over to the pile of colorful cards. There, you discover dozens of scratchcards, all with their opaque covering already scratched off. Picking one up, it looks like each card has two lists of numbers separated by a vertical bar (|): a list of winning numbers and then a list of numbers you have. You organize the information into a table (your puzzle input).

As far as the Elf has been able to figure out, you have to figure out which of the numbers you have appear in the list of winning numbers. The first match makes the card worth one point and each match after the first doubles the point value of that card.

For example:

Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11
In the above example, card 1 has five winning numbers (41, 48, 83, 86, and 17) and eight numbers you have (83, 86, 6, 31, 17, 9, 48, and 53). Of the numbers you have, four of them (48, 83, 17, and 86) are winning numbers! That means card 1 is worth 8 points (1 for the first match, then doubled three times for each of the three matches after the first).

Card 2 has two winning numbers (32 and 61), so it is worth 2 points.
Card 3 has two winning numbers (1 and 21), so it is worth 2 points.
Card 4 has one winning number (84), so it is worth 1 point.
Card 5 has no winning numbers, so it is worth no points.
Card 6 has no winning numbers, so it is worth no points.
So, in this example, the Elf's pile of scratchcards is worth 13 points.
*/
func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	sum := 0
	for s.Scan() {
		guesses, answers := getGuessesAndAnswers(s)
		common := guesses.intersectCount(answers)
		if common > 0 {
			sum += int(math.Exp2(float64(common - 1)))
		}
	}
	return sum
}

func calcPart2(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	cardCount := map[int]int{}
	card := 1
	for s.Scan() {
		cardCount[card]++
		guesses, answers := getGuessesAndAnswers(s)
		common := guesses.intersectCount(answers)
		for i := card + 1; i <= card+common; i++ {
			cardCount[i] += cardCount[card]
		}
		card++
	}

	sum := 0
	for _, v := range cardCount {
		sum += v
	}
	return sum
}

func getGuessesAndAnswers(s *bufio.Scanner) (set[int], set[int]) {
	curLine := s.Text()
	// cut off the :
	_, rest, _ := strings.Cut(curLine, ":")
	// split into guesses and answers
	guessesStr, answersStr, _ := strings.Cut(rest, "|")
	guesses := set[int]{}
	for _, v := range strings.Fields(guessesStr) {
		val, _ := strconv.Atoi(v)
		guesses[val] = true
	}
	answers := set[int]{}
	for _, v := range strings.Fields(answersStr) {
		val, _ := strconv.Atoi(v)
		answers[val] = true
	}
	return guesses, answers
}

type set[T comparable] map[T]bool

func (s set[T]) intersectCount(o set[T]) int {
	count := 0
	for v := range o {
		if s[v] {
			count++
		}
	}
	return count
}
