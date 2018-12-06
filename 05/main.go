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

func invert(input uint8) uint8 {
	if input >= 65 && input <= 90 {
		return input + 32
	} else {
		return input - 32
	}
}

func reactString(input string) (string, int) {
	result := strings.Builder{}
	reactions := 0
	for i := 0; i < len(input); i++ {
		if i != len(input) - 1 && input[i] == invert(input[i+1]) {
			reactions++
			i++
		} else {
			result.WriteByte(input[i])
		}
	}
	return result.String(), reactions
}

func fullyReact(input string) string {
	reactions := 1
	for reactions > 0 {
		input, reactions = reactString(input)
	}
	return input
}

func main() {
	//fmt.Println(reactString("aaBB"))
	//fmt.Println(reactString("aAbB"))

	lines := readInput("05/input.txt")
	input := lines[0]
	//input := "dabAcCaCBAcCcaDA"

	result := fullyReact(input)
	fmt.Println("Remaining units:", len(result))

	var i uint8
	shortest := len(input)
	for i = 65; i <= 90; i++ {
		intermediate := strings.Builder{}
		for x := 0; x < len(input); x++ {
			if input[x] != i && input[x] != invert(i) {
				intermediate.WriteByte(input[x])
			}
		}
		result := fullyReact(intermediate.String())
		//fmt.Println(i, len(result))

		if len(result) < shortest {
			shortest = len(result)
		}
	}
	fmt.Println("Shortest after removing", shortest)
}