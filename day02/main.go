package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
*/
func main() {
	data :=
		`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green`
	dr := strings.NewReader(data)
	sum := calcPart1(dr)
	fmt.Println(sum)

	dr = strings.NewReader(data)
	sum = calcPart2(dr)
	fmt.Println(sum)

	f, _ := os.Open("day02/input.txt")
	sum = calcPart1(f)
	fmt.Println(sum)
	f.Close()

	f, _ = os.Open("day02/input.txt")
	sum = calcPart2(f)
	fmt.Println(sum)
	f.Close()
}

/*
In game 1, three sets of cubes are revealed from the bag (and then put back again). The first set is 3 blue cubes and 4 red cubes; the second set is 1 red cube, 2 green cubes, and 6 blue cubes; the third set is only 2 green cubes.

The Elf would first like to know which games would have been possible if the bag contained only 12 red cubes, 13 green cubes, and 14 blue cubes?

In the example above, games 1, 2, and 5 would have been possible if the bag had been loaded with that configuration. However, game 3 would have been impossible because at one point the Elf showed you 20 red cubes at once; similarly, game 4 would also have been impossible because the Elf showed you 15 blue cubes at once. If you add up the IDs of the games that would have been possible, you get 8.

Determine which games would have been possible if the bag had been loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes. What is the sum of the IDs of those games?
*/
func calcPart1(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	sum := 0
	for s.Scan() {
		curLine := s.Text()
		gameNum, numBlue, numRed, numGreen := parseLine(curLine)
		if numBlue <= 14 && numRed <= 12 && numGreen <= 13 {
			sum += gameNum
		}
	}
	return sum
}

/*
As you continue your walk, the Elf poses a second question: in each game you played, what is the fewest number of cubes of each color that could have been in the bag to make the game possible?

Again consider the example games from earlier:

Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
In game 1, the game could have been played with as few as 4 red, 2 green, and 6 blue cubes. If any color had even one fewer cube, the game would have been impossible.
Game 2 could have been played with a minimum of 1 red, 3 green, and 4 blue cubes.
Game 3 must have been played with at least 20 red, 13 green, and 6 blue cubes.
Game 4 required at least 14 red, 3 green, and 15 blue cubes.
Game 5 needed no fewer than 6 red, 3 green, and 2 blue cubes in the bag.
The power of a set of cubes is equal to the numbers of red, green, and blue cubes multiplied together. The power of the minimum set of cubes in game 1 is 48. In games 2-5 it was 12, 1560, 630, and 36, respectively. Adding up these five powers produces the sum 2286.

For each game, find the minimum set of cubes that must have been present. What is the sum of the power of these sets?
*/
func calcPart2(dr io.Reader) int {
	s := bufio.NewScanner(dr)
	sum := 0
	for s.Scan() {
		curLine := s.Text()
		_, numBlue, numRed, numGreen := parseLine(curLine)
		power := numBlue * numRed * numGreen
		sum += power
	}
	return sum
}

func parseLine(line string) (int, int, int, int) {
	// split off the game num
	gameString, rest, _ := strings.Cut(line, ":")
	gameNum, _ := strconv.Atoi(gameString[5:])
	rounds := strings.Split(rest, ";")
	var maxBlue, maxRed, maxGreen int
	for _, v := range rounds {
		parts := strings.Split(v, ",")
		for _, v2 := range parts {
			numS, color, _ := strings.Cut(strings.TrimSpace(v2), " ")
			num, _ := strconv.Atoi(numS)
			switch color {
			case "blue":
				if num > maxBlue {
					maxBlue = num
				}
			case "red":
				if num > maxRed {
					maxRed = num
				}
			case "green":
				if num > maxGreen {
					maxGreen = num
				}
			}
		}
	}
	return gameNum, maxBlue, maxRed, maxGreen
}
