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
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

type Coord struct {
	x, y int
}

type SquareType string

const(
	LumberYard SquareType = "#"
	Forest SquareType = "|"
	Open SquareType = "."
)

func main() {
	dimension := 50
	lines := readInput("18/input.txt")
	coordMap := make(map[Coord]SquareType)
	for y, line := range lines {
		for x, char := range line {
			coord := Coord{x:x, y:y}
			switch string(char) {
			case "#": coordMap[coord] = LumberYard
			case "|": coordMap[coord] = Forest
			case ".": coordMap[coord] = Open
			}
		}
	}
	print(coordMap, dimension)
	for tick := 1; tick <= 1000000000; tick++ {
		newMap := make(map[Coord]SquareType)
		for coord, sq := range coordMap {
			forestCount, lumberyardCount, _ := countSquares(coord, coordMap)
			switch sq {
			case LumberYard:
				if forestCount >= 1 && lumberyardCount >= 1 {
					newMap[coord] = LumberYard
				} else {
					newMap[coord] = Open
				}
			case Forest:
				if lumberyardCount >= 3 {
					newMap[coord] = LumberYard
				} else {
					newMap[coord] = Forest
				}
			case Open:
				if forestCount >= 3 {
					newMap[coord] = Forest
				} else {
					newMap[coord] = Open
				}
			}
		}
		//fmt.Println("After tick:", tick)
		//print(newMap, dimension)
		forestCount := 0
		lumberyardCount := 0
		for _, sq := range coordMap {
			switch sq {
			case Forest:forestCount++
			case LumberYard:lumberyardCount++
			}
		}
		fmt.Printf("[%5d] resource value: %d\n", tick, forestCount*lumberyardCount)
		coordMap = newMap
	}
}
func countSquares(coord Coord, coordMap map[Coord]SquareType) (forestCount, lumberyardCount, openCount int) {
	for _, adjCoord := range adjacent(coord) {
		switch coordMap[adjCoord] {
		case LumberYard: lumberyardCount++
		case Forest: forestCount++
		case Open: openCount++
		}
	}
	//fmt.Println("Sum", lumberyardCount+forestCount+openCount)
	return
}

func adjacent(coord Coord) (ret []Coord) {
	for y := coord.y - 1; y <= coord.y + 1; y++ {
		for x := coord.x - 1; x <= coord.x + 1; x++ {
			if !(x == coord.x && y ==coord.y) {
				ret = append(ret, Coord{x:x, y:y})
			}
		}
	}
	return
}

func print(coordMap map[Coord]SquareType, dimension int) {
	for y := 0; y < dimension; y++ {
		for x := 0; x < dimension; x++ {
			fmt.Print(coordMap[Coord{x:x, y:y}])
		}
		fmt.Println()
	}
}

