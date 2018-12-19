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

var coordMap map[Coord]SquareType
var springs map[Coord]bool

var minX, maxX, minY, maxY int

func main() {
	lines := readInput("17/input.txt")
	coordMap = make(map[Coord]SquareType)
	springs = make(map[Coord]bool)
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
	minX = 9999999; maxX = 0; maxY = 0; minY = 99999999
	for coord, _ := range coordMap {
		if coord.y < minY {
			minY = coord.y
		}
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

	runDroplet(Coord{x: 500, y:0}, maxY)
	//renderMap(coordMap, minX, maxX, maxY)
	// All the basins are full... now let's do the flow
	/*for coord, squareType := range coordMap {
		if squareType == Flow {
			coordMap[coord] = Sand
		}
	}*/
	fmt.Println("Starting final flow")
	droplet := Coord{x: 500, y: 0}
	belowType := coordMap[Coord{x: droplet.x, y: droplet.y + 1}]
	for belowType != Settled && belowType != Clay {
		if droplet.y > maxY {
			break
		}
		coordMap[droplet] = Flow
		droplet = Coord{x: droplet.x, y: droplet.y + 1}
		belowType = coordMap[Coord{x: droplet.x, y: droplet.y + 1}]
	}
	flow(droplet, -1, maxY)
	flow(droplet, 1, maxY)

	renderMap(coordMap, minX, maxX, maxY)
	flowSum := 0
	settledSum := 0
	for coord, squareType := range coordMap {
		if coord.y >= minY && coord.y <= maxY {
			if squareType == Flow {
				flowSum++
			}
			if squareType == Settled {
				settledSum++
			}
		}
	}
	fmt.Println("Sum of water:", flowSum + settledSum)
	fmt.Println("Sum of settled water:", settledSum)
}


func runDroplet(spring Coord, maxY int) {
	spawnCounter := 0
tickLoop:
	for tick := 0; tick < 50000; tick++ {
		droplet := spring
	nextDroplet:
		for {
			below := Coord{x: droplet.x, y: droplet.y + 1}
			belowSquare := coordMap[below]
			if belowSquare == Clay || belowSquare == Settled {
				canDropLeft, leftMost := canDrop(droplet, -1)
				canDropRight, rightMost := canDrop(droplet, 1)
				if !canDropLeft && !canDropRight {
					// find a place to settle
					left := Coord{x: droplet.x - 1, y: droplet.y}
					leftSquare := coordMap[left]
					if leftSquare == Clay || leftSquare == Settled {
						right := Coord{x: droplet.x + 1, y: droplet.y}
						rightSquare := coordMap[right]
						for rightSquare != Clay && rightSquare != Settled {
							droplet = right;
							right = Coord{x: droplet.x + 1, y: droplet.y}
							rightSquare = coordMap[right]
						}
						coordMap[droplet] = Settled
						droplet = Coord{x: 500, y: 0}
						break nextDroplet
					} else {
						for leftSquare != Clay && leftSquare != Settled {
							droplet = left;
							left = Coord{x: droplet.x - 1, y: droplet.y}
							leftSquare = coordMap[left]
						}
						coordMap[droplet] = Settled
						droplet = Coord{x: 500, y: 0}
						break nextDroplet
					}
				} else if canDropRight && canDropLeft {
					if springs[droplet] {
						fmt.Println("In a loop at", droplet)
						return
					}
					springs[droplet] = true
					fmt.Println("spawn at", leftMost)
					runDroplet(leftMost, maxY)
					fmt.Println("spawn at", rightMost)
					runDroplet(rightMost, maxY)
					spawnCounter++
					if spawnCounter > 100 {
						//renderMap(coordMap, minX, maxX, maxY)
						return
					}
					fmt.Println("stopped at", droplet)
					break nextDroplet
				} else if canDropLeft {
					left := Coord{x: droplet.x - 1, y: droplet.y}
					coordMap[droplet] = Flow
					droplet = left
				} else {
					right := Coord{x: droplet.x + 1, y: droplet.y}
					coordMap[droplet] = Flow
					droplet = right
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
		//fmt.Println("===== Tick", tick)
		//renderMap(coordMap, minX, maxX, maxY)
		/*for coord, squareType := range coordMap {
			if squareType == Flow {
				coordMap[coord] = Sand
			}
		}*/
	}
}


func flow(droplet Coord, direction int, maxY int) {
	belowType := Clay
	for belowType == Clay || belowType == Settled {
		newCoord := Coord{x: droplet.x + direction, y: droplet.y}
		if coordMap[newCoord] != Clay && coordMap[newCoord] != Settled {
			// move into that direction
			coordMap[droplet] = Flow
			droplet = newCoord
			belowType = coordMap[Coord{x: droplet.x, y: droplet.y + 1}]
		} else {
			// stuck
			coordMap[droplet] = Flow
			return
		}
	}
	for belowType != Clay && belowType != Settled {
		if droplet.y > maxY {
			return
		}
		coordMap[droplet] = Flow
		droplet = Coord{x: droplet.x, y: droplet.y + 1}
		belowType = coordMap[Coord{x: droplet.x, y: droplet.y + 1}]
	}
	flow(droplet, -1, maxY)
	flow(droplet, 1, maxY)
}

func canDrop(coord Coord, direction int) (bool, Coord) {
	leftCoord := Coord{x:coord.x + direction, y: coord.y}
	if coordMap[leftCoord] == Clay || coordMap[leftCoord] == Settled {
		return false, Coord{}
	}
	below := Coord{x:leftCoord.x, y:leftCoord.y+1}
	if coordMap[below] != Clay && coordMap[below] != Settled {
		return true, leftCoord
	}
	return canDrop(leftCoord, direction)
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
