package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

/*
In Camel Cards, you get a list of hands, and your goal is to order them based on the strength of each hand. A hand consists of five cards labeled one of A, K, Q, J, T, 9, 8, 7, 6, 5, 4, 3, or 2. The relative strength of each card follows this order, where A is the highest and 2 is the lowest.

Every hand is exactly one type. From strongest to weakest, they are:

Five of a kind, where all five cards have the same label: AAAAA
Four of a kind, where four cards have the same label and one card has a different label: AA8AA
Full house, where three cards have the same label, and the remaining two cards share a different label: 23332
Three of a kind, where three cards have the same label, and the remaining two cards are each different from any other card in the hand: TTT98
Two pair, where two cards share one label, two other cards share a second label, and the remaining card has a third label: 23432
One pair, where two cards share one label, and the other three cards have a different label from the pair and each other: A23A4
High card, where all cards' labels are distinct: 23456
Hands are primarily ordered based on type; for example, every full house is stronger than any three of a kind.

If two hands have the same type, a second ordering rule takes effect. Start by comparing the first card in each hand. If these cards are different, the hand with the stronger first card is considered stronger. If the first card in each hand have the same label, however, then move on to considering the second card in each hand. If they differ, the hand with the higher second card wins; otherwise, continue with the third card in each hand, then the fourth, then the fifth.

So, 33332 and 2AAAA are both four of a kind hands, but 33332 is stronger because its first card is stronger. Similarly, 77888 and 77788 are both a full house, but 77888 is stronger because its third card is stronger (and both hands have the same first and second card).

To play Camel Cards, you are given a list of hands and their corresponding bid (your puzzle input). For example:

32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
This example shows five hands; each hand is followed by its bid amount. Each hand wins an amount equal to its bid multiplied by its rank, where the weakest hand gets rank 1, the second-weakest hand gets rank 2, and so on up to the strongest hand. Because there are five hands in this example, the strongest hand will have rank 5 and its bid will be multiplied by 5.

So, the first step is to put the hands in order of strength:

32T3K is the only one pair and the other hands are all a stronger type, so it gets rank 1.
KK677 and KTJJT are both two pair. Their first cards both have the same label, but the second card of KK677 is stronger (K vs T), so KTJJT gets rank 2 and KK677 gets rank 3.
T55J5 and QQQJA are both three of a kind. QQQJA has a stronger first card, so it gets rank 5 and T55J5 gets rank 4.
Now, you can determine the total winnings of this set of hands by adding up the result of multiplying each hand's bid with its rank (765 * 1 + 220 * 2 + 28 * 3 + 684 * 4 + 483 * 5). So the total winnings in this example are 6440.
*/
func main() {
	data :=
		`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`
	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	f, _ := os.Open("day07/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day07/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()
}

type Rank int

const (
	High Rank = iota
	OneP
	TwoP
	Three
	FH
	Four
	Five
)

func findRank(hand string) Rank {
	m := map[rune]int{}
	for _, v := range hand {
		m[v]++
	}
	countMap := map[int]int{}
	for _, v := range m {
		countMap[v]++
	}
	if _, ok := countMap[5]; ok {
		return Five
	}
	if _, ok := countMap[4]; ok {
		return Four
	}
	if _, ok := countMap[3]; ok {
		if _, ok := countMap[2]; ok {
			return FH
		}
		return Three
	}
	if num, ok := countMap[2]; ok {
		if num == 1 {
			return OneP
		}
		return TwoP
	}
	return High
}

type handBid struct {
	hand string
	bid  int
	rank Rank
}

func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	var hands []handBid
	for s.Scan() {
		curLine := s.Text()
		parts := strings.Fields(curLine)
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, handBid{parts[0], bid, findRank(parts[0])})
	}
	//fmt.Println(hands)
	sort.Slice(hands, func(i, j int) bool {
		firstHand := hands[i]
		secondHand := hands[j]
		if firstHand.rank < secondHand.rank {
			return true
		}
		if firstHand.rank > secondHand.rank {
			return false
		}
		for handI := 0; handI < 5; handI++ {
			if order[firstHand.hand[handI]] < order[secondHand.hand[handI]] {
				return true
			}
			if order[firstHand.hand[handI]] > order[secondHand.hand[handI]] {
				return false
			}
		}
		// should be impossible
		return true
	})

	//fmt.Println(hands)
	sum := 0
	for i, v := range hands {
		sum += v.bid * (i + 1)
	}

	return sum
}

var order = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

func calcPart2(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	var hands []handBid
	for s.Scan() {
		curLine := s.Text()
		parts := strings.Fields(curLine)
		bid, _ := strconv.Atoi(parts[1])
		hands = append(hands, handBid{parts[0], bid, findRank2(parts[0])})
	}
	fmt.Println(hands)
	sort.Slice(hands, func(i, j int) bool {
		firstHand := hands[i]
		secondHand := hands[j]
		if firstHand.rank < secondHand.rank {
			return true
		}
		if firstHand.rank > secondHand.rank {
			return false
		}
		for handI := 0; handI < 5; handI++ {
			if order2[firstHand.hand[handI]] < order2[secondHand.hand[handI]] {
				return true
			}
			if order2[firstHand.hand[handI]] > order2[secondHand.hand[handI]] {
				return false
			}
		}
		// should be impossible
		return true
	})

	fmt.Println(hands)
	sum := 0
	for i, v := range hands {
		sum += v.bid * (i + 1)
	}

	return sum
}

func findRank2(hand string) Rank {
	m := map[rune]int{}
	numJ := 0
	for _, v := range hand {
		if v == 'J' {
			numJ++
		} else {
			m[v]++
		}
	}
	countMap := map[int]int{}
	for _, v := range m {
		countMap[v]++
	}
	if _, ok := countMap[5]; ok {
		return Five
	}
	if _, ok := countMap[4]; ok {
		if numJ == 1 {
			return Five
		}
		return Four
	}
	if _, ok := countMap[3]; ok {
		if numJ == 2 {
			return Five
		}
		if numJ == 1 {
			return Four
		}
		if _, ok := countMap[2]; ok {
			return FH
		}
		return Three
	}
	if num, ok := countMap[2]; ok {
		if num == 1 {
			if numJ == 3 {
				return Five
			}
			if numJ == 2 {
				return Four
			}
			if numJ == 1 {
				return Three
			}
			return OneP
		}
		if numJ == 1 {
			return FH
		}
		return TwoP
	}
	if numJ == 5 {
		return Five
	}
	if numJ == 4 {
		return Five
	}
	if numJ == 3 {
		return Four
	}
	if numJ == 2 {
		return Three
	}
	if numJ == 1 {
		return OneP
	}
	return High
}

var order2 = map[byte]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
	'J': 0,
}
