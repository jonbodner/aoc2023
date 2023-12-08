package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strings"
)

func main() {
	data :=
		`RL

AAA = (BBB, CCC)
BBB = (DDD, EEE)
CCC = (ZZZ, GGG)
DDD = (DDD, DDD)
EEE = (EEE, EEE)
GGG = (GGG, GGG)
ZZZ = (ZZZ, ZZZ)`
	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	data =
		`LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)`
	dr = strings.NewReader(data)
	sum = calcPart1(dr)
	fmt.Println(sum)

	data =
		`LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)`
	dr = strings.NewReader(data)
	sum2 := calcPart2(dr)
	fmt.Println(sum2)

	f, _ := os.Open("day08/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day08/input.txt")
	sum2 = calcPart2(f)
	fmt.Println(sum2)
	f.Close()

}

func calcPart2(dr io.Reader) uint64 {
	s := bufio.NewScanner(dr)
	s.Scan()
	rules := s.Text()
	s.Scan()
	nodes := buildNodes(s)
	sum := walk2(nodes, rules)
	return sum
}

func walk2(nodes map[string]*node, rules string) uint64 {
	var curNodes []*node
	for k, v := range nodes {
		if strings.HasSuffix(k, "A") {
			curNodes = append(curNodes, v)
		}
	}
	var counts []int
	for _, v := range curNodes {
		count := 0
		pos := 0
		curNode := v
		for !strings.HasSuffix(curNode.name, "Z") {
			switch rules[pos] {
			case 'R':
				curNode = curNode.right
			case 'L':
				curNode = curNode.left
			}
			count++
			pos = (pos + 1) % len(rules)

		}
		counts = append(counts, count)
	}
	// find the least common multiple
	var primes []int
	done := func() bool {
		for _, v := range counts {
			if v != 1 {
				return false
			}
		}
		return true
	}
	curPrime := 2
	for !done() {
		found := false
		for i := 0; i < len(counts); i++ {
			if counts[i]%curPrime == 0 {
				counts[i] = counts[i] / curPrime
				found = true
			}
		}
		if found {
			primes = append(primes, curPrime)
		} else {
			curPrime = nextPrime(curPrime)
		}
	}

	sum := uint64(1)
	for _, v := range primes {
		sum *= uint64(v)
	}

	return sum
}

func nextPrime(curPrime int) int {
	i := curPrime + 1
	for {
		stop := int(math.Sqrt(float64(i)))
		found := false
		for j := 2; j <= stop; j++ {
			if x := i / j; x*j == i {
				found = true
				break
			}
		}
		if !found {
			return i
		}
		i++
	}
}

type node struct {
	name  string
	left  *node
	right *node
}

func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	s.Scan()
	rules := s.Text()
	s.Scan()
	nodes := buildNodes(s)
	sum := walk(nodes, rules)
	return sum
}

func walk(nodes map[string]*node, rules string) int {
	count := 0
	pos := 0
	curNode := nodes["AAA"]
	for curNode.name != "ZZZ" {
		switch rules[pos] {
		case 'R':
			curNode = curNode.right
		case 'L':
			curNode = curNode.left
		}
		count++
		pos = (pos + 1) % len(rules)
	}
	return count
}

func buildNodes(s *bufio.Scanner) map[string]*node {
	nodes := map[string]*node{}
	for s.Scan() {
		name, rules, _ := strings.Cut(s.Text(), " = ")
		l, r, _ := strings.Cut(rules[1:len(rules)-1], ", ")
		var curNode *node
		var ok bool
		if curNode, ok = nodes[name]; !ok {
			curNode = &node{
				name: name,
			}
			nodes[name] = curNode
		}
		var lNode *node
		if lNode, ok = nodes[l]; !ok {
			lNode = &node{
				name: l,
			}
			nodes[l] = lNode
		}
		curNode.left = lNode
		var rNode *node
		if rNode, ok = nodes[r]; !ok {
			rNode = &node{
				name: r,
			}
			nodes[r] = rNode
		}
		curNode.right = rNode
	}
	return nodes
}
