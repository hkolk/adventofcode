package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func findClosest(y, x int) int {
	duplicate := false
	closestId := 0
	closestDistance := 99999999999
	for id, crd := range points {
		distance := crd.distance(x, y)
		if distance == closestDistance {
			duplicate = true
		} else if distance < closestDistance {
			duplicate = false
			closestId = id
			closestDistance = distance
		}
	}
	if duplicate {
		return 0
	} else {
		return closestId
	}
}

func isWithinDistance(x, y, distance int) bool {
	totalDistance := 0
	for _, crd := range points {
		totalDistance += crd.distance(x, y)
		if totalDistance >= distance {
			return false
		}
	}
	return true
}

type coord struct {
	x int
	y int
}

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

func (me coord) distance(x, y int) int {
	return abs(me.x - x) + abs(me.y - y)
}


func part2(maxY int, maxX int, distance int) {
	grid := make(map[int]map[int]bool)
	for y := 0; y <= maxY; y++ {
		grid[y] = make(map[int]bool)
		for x := 0; x <= maxX; x++ {
			grid[y][x] = isWithinDistance(y, x, distance)
		}
	}
	count := 0
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if grid[y][x] {
				count++
				if debug { fmt.Print("#") }
			} else {
				if debug { fmt.Print(".") }
			}
		}
		if debug { fmt.Println() }
	}
	fmt.Println("Part 2 count:", count)
}

func part1(maxY int, maxX int) {
	// Part 1 Grid
	grid := make(map[int]map[int]int)
	for y := 0; y <= maxY; y++ {
		grid[y] = make(map[int]int)
		for x := 0; x <= maxX; x++ {
			grid[y][x] = findClosest(y, x)
		}
	}
	if debug {
		for y := 0; y <= maxY; y++ {
			for x := 0; x <= maxX; x++ {
				fmt.Printf("%2d ", grid[y][x])
			}
			fmt.Println()
		}
	}
	exclusions := make(map[int]bool)
	counts := make(map[int]int)
	for x := 0; x <= maxX; x++ {
		exclusions[grid[0][x]] = true
		exclusions[grid[maxY][x]] = true
	}
	for y := 0; y <= maxY; y++ {
		exclusions[grid[y][0]] = true
		exclusions[grid[y][maxX]] = true
	}
	for y := 0; y <= maxY; y++ {
		for x := 0; x <= maxX; x++ {
			if !exclusions[grid[y][x]] {
				counts[grid[y][x]]++
			}
		}
	}
	if debug { fmt.Println(counts) }
	maxCount := 0
	for _, count := range counts {
		if count > maxCount {
			maxCount = count
		}
	}
	fmt.Println("Part 1:", maxCount)
}

var points = make(map[int]coord)

var debug = false
func main() {
	filename := "06/input.txt"
	distance := 10000
	if debug {
		filename = "06/input.test.txt"
		distance = 32
	}
	lines := readInput(filename)
	//points := make(map[int]map[int]int)
	var maxX, maxY int
	for id, line := range lines {
		var x, y int
		fmt.Sscanf(line, "%d, %d", &x, &y)
		if x > maxX {
			maxX = x
		}
		if y > maxY {
			maxY = y
		}
		points[id+1] = coord{x: x, y: y}
	}

	part1(maxY, maxX)
	part2(maxY, maxX, distance)
}
