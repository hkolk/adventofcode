package main

import ( 
	"fmt"
	"bufio"
	"strconv"
	"os"
	"log"
	"regexp"
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

type area struct {
	name string
	startY int
	startX int
	width int
	height int
}

func unsafeInt(input string) int {
	num, _ := strconv.Atoi(input)
	return num
}

func parseInput(input []string) []area {
	areas := make([]area, 0)
	var re = regexp.MustCompile("^#([0-9]+) @ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)$")
	for _, line := range input {
		parts := re.FindStringSubmatch(line)
		areas = append(areas, area {
			name: parts[1],
			startX: unsafeInt(parts[2]),
			startY: unsafeInt(parts[3]),
			width: unsafeInt(parts[4]),
			height: unsafeInt(parts[5])})
	}
	return areas		
}

func main() {
	var grid [1000][1000]int
	
	lines := readInput("03/input.txt")
	areas := parseInput(lines)
	for _, area := range areas {
		for x := area.startX; x < area.startX + area.width; x++ {
			for y := area.startY; y < area.startY + area.height; y++ {
				grid[x][y]++
			}
		}
	}
	overlap := 0
	for _, line := range grid {
		//fmt.Println(line)
		for _, space := range line {
			if space > 1 {
				overlap++
			}
		}
	}
	fmt.Println("Overal: ", overlap)
	for _, area := range areas {
		overlapping := false
		for x := area.startX; x < area.startX + area.width; x++ {
			for y := area.startY; y < area.startY + area.height; y++ {
				if grid[x][y] != 1 {
					overlapping = true
				}
			}
		}
		if !overlapping {
			fmt.Println("Non-overlapping area: ", area.name)
		}
	}
}
