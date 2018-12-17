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

type SquareType int

const (
	Sand SquareType = 1
	Clay SquareType = 2
	Flow SquareType = 3
	Settled SquareType = 4
	Droplet SquareType = 5
)

func main() {
	lines := readInput("17/input.test.txt")
	coordMap := make(map[Coord]SquareType)
	for _, line := range lines {
		var part1, part2, part3 int
		_, error := fmt.Sscanf(line, "y=%d, x=%d..%d", &part1, &part2, &part3)
		if error == nil {
			for x := part2; x <= part3; x++ {
				coordMap[Coord{y:part1, x:x}] = Clay
			}

		} else {
			_, error = fmt.Sscanf(line, "x=%d, y=%d..%d", &part1, &part2, &part3)
			if error == nil {
				for y := part2; y <= part3; y++ {
					coordMap[Coord{y: y, x: part1}] = Clay
				}
			}
		}
	}
	//fmt.Println(coordMap)
	minX := 9999999; maxX := 0; maxY := 0
	for coord, _ := range coordMap {
		if coord.y > maxY {
			maxY = coord.y
		}
		if coord.x > maxX {
			maxX = coord.x
		}
		if coord.x < minX {
			minX = coord.x
		}
	}
	//renderMap(coordMap, minX, maxX, maxY)

	tickLoop:
	for tick := 0; tick < 40; tick++ {
		droplet := Coord{x:500, y:0}
		nextDroplet:
		for {
			below := Coord{x: droplet.x, y: droplet.y + 1}
			belowSquare := coordMap[below]
			if belowSquare == Clay || belowSquare == Settled {
				// move left/right
				left := Coord{x: droplet.x - 1, y: droplet.y}
				leftSquare := coordMap[left]
				if leftSquare == Clay || leftSquare == Settled || leftSquare == Flow  {
					if !canFlow(droplet, coordMap, 1) {
						// settle
						if tick == 14 { fmt.Println("canflow was false") }
						coordMap[droplet] = Settled
						droplet = Coord{x: 500, y: 0}
						break nextDroplet
					} else {
						if tick == 14 { fmt.Println("canflow was true") }

						right := Coord{x: droplet.x + 1, y: droplet.y}
						rightSquare := coordMap[right]
						if rightSquare == Clay || rightSquare == Settled {
							// settle
							coordMap[droplet] = Settled
							droplet = Coord{x: 500, y: 0}
							break nextDroplet
						} else {
							coordMap[droplet] = Flow
							droplet = right
						}
					}
				} else {
					coordMap[droplet] = Flow
					droplet = left
				}

			} else {
				// move down?
				if below.y > maxY {
					// Dropped off the world! Done?
					break tickLoop
				}
				coordMap[droplet] = Flow
				droplet = below
			}

		}

		coordMap[droplet] = Droplet
		fmt.Println("===== Tick", tick)
		renderMap(coordMap, minX, maxX, maxY)
		for coord, squareType := range coordMap {
			if squareType == Flow {
				coordMap[coord] = Sand
			}
		}
	}
	renderMap(coordMap, minX, maxX, maxY)
}

func canFlow(coord Coord, coordMap map[Coord]SquareType, direction int) bool {
	leftCoord := Coord{x:coord.x + direction, y: coord.y}
	if coordMap[leftCoord] == Clay || coordMap[leftCoord] == Settled {
		return false
	}
	below := Coord{x:leftCoord.x, y:leftCoord.y+1}
	if coordMap[below] != Clay && coordMap[below] != Settled {
		return true
	}
	return canFlow(leftCoord, coordMap, direction)
}

func toChar(square SquareType) string {
	switch square {
	case Clay: return "#"
	case Flow: return "|"
	case Settled: return "~"
	case Droplet: return "*"
	default: return "."
	}
}

func renderMap(coordMap map[Coord]SquareType, minX, maxX, maxY int) {
	for y := 0; y <= maxY+2; y++ {
		for x := minX-2; x <= maxX+2; x++ {
			fmt.Print(toChar(coordMap[Coord{x:x, y:y}]))
		}
		fmt.Println()
	}

}
