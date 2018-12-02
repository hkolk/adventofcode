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

func roughMatch(left string, right string) bool {
	mismatched := 0
	if len(left) != len(right) {
		return false
	}
	for i, _ := range left {
		if right[i] != left[i] {
			mismatched++
		}
	}
	return mismatched == 1
}

func main() {
	lines := readInput("input.txt")
	for _, left := range lines {
		for _, right := range lines {
			if roughMatch(left, right) {
				fmt.Println(left, right)
			}
		}
	}
	//fmt.Println(roughMatch("aabbccdd", "aabbccdd"))
	//fmt.Println(roughMatch("aabbccdd", "aabbccde"))
	//fmt.Println(roughMatch("aabbccdd", "aabbccee"))
}
