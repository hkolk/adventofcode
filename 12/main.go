package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readInput(filename string) []string {
	var lines []string
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	_ = file.Close()
	return lines
}

type pattern struct {
	pattern string
	result string
}

type pots struct {
	pots string
	lowestIndex int
}

func (p *pots) zeroPad() {
	leftPadding := 0
	rightPadding := 0
	for i := 0; i < padding; i++ {
		if string(p.pots[i]) == "#" {
			leftPadding = 1
		}

		if string(p.pots[len(p.pots) - 1 - i]) == "#" {
			rightPadding = 1
		}
	}
	p.pots = strings.Repeat(".", leftPadding*padding) + p.pots + strings.Repeat(".", rightPadding*padding)
	p.lowestIndex -= leftPadding*padding
	//p.pots = "...."+p.pots+"...."
	//p.lowestIndex -= 4

	//highestIndex := len(p.pots) - 1
	//if string(p.pots[highestIndex]) == "#" || string(p.pots[highestIndex-1]) == "#" || string(p.pots[highestIndex-2]) == "#" {
	//	p.pots = p.pots + ".."
	//}
}

func printPots(pots *pots) {
	prefix := 10
	fmt.Print(strings.Repeat(".", pots.lowestIndex+prefix))
	fmt.Print(pots.pots)
	fmt.Print("\t", pots.lowestIndex)
	fmt.Print("\t", sum(pots))

	fmt.Println()
}

func sum(pots *pots) int {
	sum := 0
	for i, char := range pots.pots {
		if string(char) == "#" {
			sum += i + pots.lowestIndex
		}
	}
	return sum
}

var debug = false
var padding = 4
func main() {
	filename := "12/input.txt"
	initial := "###..#...####.#..###.....####.######.....##.#####.##.##..###....#....##...##...##.#..###..#.#...#..#"
	if debug {
		filename = "12/input.test.txt"
		initial = "#..#.#..##......###...###"
	}
	var patterns []pattern
	lines := readInput(filename)
	for _, line := range lines {
		var patternString, result string
		fmt.Sscanf(line, "%5s => %1s", &patternString, &result)
		patterns = append(patterns, pattern{pattern: patternString, result: result})
	}
	pots := &pots{pots:initial, lowestIndex:0}
	pots.zeroPad()

	fmt.Printf("%2d: ", 0)
	printPots(pots)

	diff := 0
	prevSum := 0
	for i := 0; i < 1000; i++ {
		match(pots, patterns)
		pots.lowestIndex += 2
		pots.zeroPad()
		if i == 19  || (i+1) % 100 == 0 {
			fmt.Printf("%2d:", i+1)
			printPots(pots)
		}
		sum := sum(pots)
		diff = sum - prevSum
		prevSum = sum;
		//}
	}

	fmt.Println("Skipping",50000000000-1000, "...", prevSum, diff,  prevSum + (diff*(50000000000-1000)))
}

func match(pots *pots, patterns []pattern) {
	sb := strings.Builder{}
	for i := 2; i < len(pots.pots)-2; i++ {
		result := "."
	patternLoop:
		for _, pattern := range patterns {

			for x := 0; x < 5; x++ {
				if !(pattern.pattern[x] == pots.pots[i+x-2]) {
					//fmt.Println("Miss: ", pattern.pattern[x], pots.pots[i+x-2])
					continue patternLoop
				} else {
					//fmt.Println("Hit: ", pattern.pattern[x], pots.pots[i+x-2])
				}
			}
			// match!
			//fmt.Println("Found match for", pattern.pattern, "at", i)
			result = pattern.result
			break patternLoop
		}
		sb.WriteString(result)

	}
	pots.pots = sb.String()
}