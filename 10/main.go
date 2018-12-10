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

type elf struct {
	id int
	x int
	y int
	xVelo int
	yVelo int
}

func (s *elf) move() {
	s.x += s.xVelo
	s.y += s.yVelo
}

func dimensions(elfs []*elf) (minX, maxX, minY, maxY int) {
	for _, elf := range elfs {
		if elf.x < minX {
			minX = elf.x
		}
		if elf.x > maxX {
			maxX = elf.x
		}
		if elf.y < minY {
			minY = elf.y
		}
		if elf.y > maxY {
			maxY = elf.y
		}
	}
	return minX, maxX, minY, maxY
}

func printElfs(elfMap map[int]map[int]bool) {
	var minX, maxX, minY, maxY int
	hasX := false
	hasY := false
	for x, mapY := range elfMap {
		if !hasX || x < minX {
			minX = x
		}
		if !hasX || x > maxX {
			maxX = x
		}
		hasX = true
		for y, _ := range mapY {
			if !hasY || y < minY {
				minY = y
			}
			if !hasY || y > maxY {
				maxY = y
			}
			hasY = true
		}
	}
	//minX, maxX, minY, maxY := dimensions(elfs)
	//fmt.Println(minX, maxX, minY, maxY)
	//elfMap := generateElfMap(elfs)
	//fmt.Println(elfMap)
	for y := minY-2; y <= maxY+2; y++ {
		for x := minX-2; x <= maxX+2; x++ {
			if elfMap[x][y] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func generateElfMap(elfs []*elf) map[int]map[int]bool {
	elfMap := make(map[int]map[int]bool)
	for _, elf := range elfs {
		_, exists := elfMap[elf.x]
		if !exists {
			elfMap[elf.x] = make(map[int]bool)
		}
		elfMap[elf.x][elf.y] = true
	}
	return elfMap
}


var debug = false
func main() {
	filename := "10/input.txt"
	if debug {
		filename = "10/input.test.txt"
	}
	lines := readInput(filename)

	var elfs []*elf

	for elfId, line := range lines {
		var x, y, veloX, veloY int
		fmt.Sscanf(line, "position=<%8d, %8d> velocity=<%8d, %8d>", &x, &y, &veloX, &veloY)
		elf := elf{id: elfId, x: x, y: y, xVelo: veloX, yVelo: veloY}
		elfs = append(elfs, &elf)
		//fmt.Println(x, y, veloX, veloY)
	}

	//fmt.Println("==== Second", 0, "=====")
	//printElfs(elfs)
	minX, maxX, minY, maxY := dimensions(elfs)
	prevConvergence := (maxX - minX) + (maxY - minY)
	prevElfMap := generateElfMap(elfs)
	for i := 1; i < 1000000; i++ {
		for _, elf := range elfs {
			elf.move()
		}

		minX, maxX, minY, maxY := dimensions(elfs)
		convergence := (maxX - minX) + (maxY - minY)
		fmt.Println("Convergence at second", i, "is", convergence)
		if convergence < prevConvergence {
			prevConvergence = convergence
			prevElfMap = generateElfMap(elfs)
		} else {
			printElfs(prevElfMap)
			fmt.Println("Minimal convergence found at second", i-1, "with a convergence of", prevConvergence)
			break
		}


		//fmt.Println("==== Second", i, "=====")
		//printElfs(generateElfMap(elfs))


	}
}
