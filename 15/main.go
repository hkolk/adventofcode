package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func abs(n int) int {
	y := n >> 63
	return (n ^ y) - y
}

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

type Race int
const (
	GoblinRace Race = 1
	ElfRace Race = 2
)

type Unit struct {
	race Race
	postion Coord
	hitpoints int
	attackPower int
	dead bool
}

func (self *Unit) Adjacent(other *Unit) bool {
	if other.postion.x == self.postion.x {
		return abs(other.postion.y - self.postion.y) <= 1
	} else if other.postion.y == self.postion.y {
		return abs(other.postion.x - self.postion.x) <= 1
	}
	return false
}

func (a Coord) Less(b Coord) bool {
	if a.y == b.y {
		return a.x < b.x
	}
	return a.y < b.y
}

type ByCoords []*Unit

func (a ByCoords) Len() int { return len(a)}
func (a ByCoords) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoords) Less(i, j int) bool {
	return a[i].postion.Less(a[j].postion)
	/*
	if a[i].postion.y == a[j].postion.y {
		return a[i].postion.x < a[j].postion.x
	}
	return a[i].postion.y < a[j].postion.y
	*/
}

type Space int

const (
	Open Space = 1
	Wall Space = 2
	Goblin Space = 3
	Elf Space = 4
)

func toIdentifier(space Space) string {
	switch space {
	case Open: return "."
	case Wall: return "#"
	case Goblin: return "G"
	case Elf: return "E"
	}
	panic("unknown char")
}

func printMap(caveMap map[Coord]Space, maxX, maxY int ) {
	for y := 0; y < maxY; y++ {
		for x := 0; x < maxX; x++ {
			fmt.Print(toIdentifier(caveMap[Coord{x:x, y:y}]))
		}
		fmt.Println()
	}
}


func run(inputFile string, elfAttackPower int, debug bool) (deadElves int, outcome int) {
	caveMap := make(map[Coord]Space)
	lines := readInput(inputFile)
	var goblins, elfs, units []*Unit
	var maxX, maxY int
	maxY = len(lines)
	for y, line := range lines {
		maxX = len(line)
		for x, char := range line {
			coord := Coord{x: x, y: y}
			switch string(char) {
			case "G":
				unit := &Unit{race: GoblinRace, postion: coord, hitpoints:200, attackPower:3}
				goblins = append(goblins, unit)
				units = append(units, unit)
				caveMap[coord] = Goblin
			case "E":
				unit := &Unit{race: ElfRace, postion: coord, hitpoints:200, attackPower:elfAttackPower}
				elfs = append(elfs, unit)
				units = append(units, unit)
				caveMap[coord] = Elf
			case "#":
				caveMap[coord] = Wall
			case ".":
				caveMap[coord] = Open
			}
		}
	}

	// let's start ticking
	var round int
	tickLoop:
	for round = 1; round < 1000; round++ {
		if debug { fmt.Println("Starting round:", round) }
		if debug { printMap(caveMap, maxX, maxY) }
		sort.Sort(ByCoords(units))
		for _, unit := range units {
			if unit.dead {
				continue
			}
			// find target(s)
			var targets []*Unit
			if unit.race == GoblinRace {
				targets = elfs
			} else {
				targets = goblins
			}
			if len(targets) == 0 {
				break tickLoop
			}
			// find movable space
			explorableSpace := findExplorableSpace(unit.postion, caveMap)
			attacking := false
			var options []Coord
			for _, target := range targets {
				if unit.Adjacent(target) {
					attacking = true
					break
				} else {
					options = append(options, surroundingCoords(target.postion)...)
				}

			}
			if !attacking {
				foundSpot := false
				lowestDistance := 99999999999
				var closestCoord Coord
				for _, coord := range options {
					distance, exists := explorableSpace[Coord{x: coord.x, y: coord.y}]
					if exists {
						//fmt.Println("Reachable:", coord)
						foundSpot = true
						if distance < lowestDistance {
							lowestDistance = distance
							closestCoord = coord
						} else if distance == lowestDistance && coord.Less(closestCoord) {
							closestCoord = coord
						}
					}
				}
				if foundSpot {
					//fmt.Println("Found a nice spot:", closestCoord, "with distance", lowestDistance)
					distanceToTarget := findExplorableSpace(closestCoord, caveMap)
					lowestDistanceToTarget := 99999999999
					var closestCoordToTarget Coord
					for _, coord := range surroundingCoords(unit.postion) {
						distance, exists := distanceToTarget[coord]
						if exists {
							if distance < lowestDistanceToTarget {
								lowestDistanceToTarget = distance
								closestCoordToTarget = coord
							} else if distance == lowestDistanceToTarget && coord.Less(closestCoordToTarget) {
								closestCoordToTarget = coord
							}
						}
					}
					//fmt.Println("Will move to position", closestCoordToTarget)
					// actually move
					caveMap[unit.postion] = Open
					if unit.race == GoblinRace {
						caveMap[closestCoordToTarget] = Goblin
					} else {
						caveMap[closestCoordToTarget] = Elf
					}
					unit.postion = closestCoordToTarget
				}
			}
			//attacking = false
			pickedTarget := false
			attackingTarget := &Unit{hitpoints: 9999999999}
			for _, target := range targets {
				if unit.Adjacent(target) {
					pickedTarget = true
					if target.hitpoints < attackingTarget.hitpoints {
						attackingTarget = target
					} else if target.hitpoints == attackingTarget.hitpoints && target.postion.Less(attackingTarget.postion) {
						attackingTarget = target
					}
				}
			}
			if pickedTarget {
				attackingTarget.hitpoints -= unit.attackPower
				if attackingTarget.hitpoints <= 0 {
					attackingTarget.dead = true
					caveMap[attackingTarget.postion] = Open
					if attackingTarget.race == ElfRace {
						deadElves++
					}
				}
				// mortal kombat!
			}
			elfs = pruneUnitList(elfs)
			goblins = pruneUnitList(goblins)
		}
		units = pruneUnitList(units)
		if debug {
			fmt.Printf("After round %d there are %d units, of wich %d goblins and %d elfs\n", round, len(units), len(goblins), len(elfs))
		}
	}
	fmt.Printf("[%3d] == Combat ended in round: %d\n", elfAttackPower, round-1)
	units = pruneUnitList(units)
	healthSum := 0
	for _, unit := range units {
		healthSum += unit.hitpoints
	}
	fmt.Printf("[%3d] == Remaining health: %d\n", elfAttackPower, healthSum)
	outcome = healthSum * (round - 1)
	fmt.Printf("[%3d] == Outcome: %d\n", elfAttackPower, outcome)
	fmt.Printf("[%3d] == Dead elves: %d\n", elfAttackPower, deadElves)
	return
}

func pruneUnitList(list []*Unit) []*Unit {
	var newList []*Unit
	for _, unit := range list {
		if !unit.dead {
			newList = append(newList, unit)
		}
	}
	return newList
}

func findExplorableSpace(originalPosition Coord, caveMap map[Coord]Space) map[Coord]int {
	explorableSpace := make(map[Coord]int)
	var toVisit []Coord
	toVisit = append(toVisit, originalPosition)
	//explorableSpace[unit.postion] = 0
	distanceCounter := 0
	for len(toVisit) > 0 {
		var newVisit []Coord
		for _, coord := range toVisit {
			_, alreadyExplored := explorableSpace[coord]
			if !alreadyExplored && (caveMap[coord] == Open || originalPosition == coord) {
				explorableSpace[coord] = distanceCounter
				newVisit = append(newVisit, surroundingCoords(coord)...)
			}
		}
		toVisit = newVisit
		distanceCounter++
		//fmt.Println(explorableSpace)
	}
	return explorableSpace
}

func surroundingCoords(coord Coord) []Coord {
	ret := []Coord{
		{x:coord.x - 1, y: coord.y},
		{x:coord.x + 1, y: coord.y},
		{x:coord.x, y: coord.y - 1},
		{x:coord.x, y: coord.y + 1}}
	return ret
}

func main() {
	_, part1 := run("15/input.txt", 3, false)

	deadElves := 1
	attackPower := 3
	part2 := 0
	for deadElves > 0 {
		attackPower++
		deadElves, part2 = run("15/input.txt", attackPower, false)
	}
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}