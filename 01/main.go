package main

import ( 
	"fmt"
	"bufio"
	"strconv"
	"os"
	"log"
)

func readInput(filename string) []string {
	lines := make([]string, 0)
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

func main() {
	found := make(map[int]bool)
	result := 0
	numbers, err := stringToInt(readInput("input.txt"))
	if err != nil {
		log.Fatal(err)
	}
	for {
		for _, value := range numbers {
			result += value
			if found[result] {
				fmt.Println("Found duplicate: ", result)
				os.Exit(-1)
			}
			found[result] = true
		}
		fmt.Println("Result: ", result)
	}	
}

