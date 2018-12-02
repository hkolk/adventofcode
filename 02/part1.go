package main

import ( 
	"fmt"
	"bufio"
	"strconv"
	"os"
	"log"
)

func readInput(filename string) []string {
	lines := []string{}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	file.Close()
	return lines
}

func stringToInt(input []string) ([]int, error) {
	output := make([]int, 0)
	for _, i := range input {
		num, err := strconv.Atoi(i)
		if err != nil {
			return output, err
		}
		output = append(output, num)
	}
	return output, nil
}

func countChars(input string) map[rune]int {
	charmap := make(map[rune]int)
	for _, char := range input {
		charmap[char] += 1
	}
	return charmap
}

func countOccurences(input map[rune]int) map[int]bool {
	newmap := make(map[int]bool)
	for _, value := range input {
		newmap[value] = true
	}
	return newmap
}

func main() {
	twos := 0
	threes := 0
	for _, line := range readInput("input.txt") {
		fmt.Print(line, ": ")
		counts := countOccurences(countChars(line))
		fmt.Println(counts)
		if counts[2] {
			twos += 1
		}
		if counts[3] {
			threes += 1
		}

	}
	fmt.Println("twos: ", twos, ", threes: ", threes, ", multiplied: ", twos * threes)
}

