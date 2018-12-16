package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

type Track int
const (
	Horizontal 	Track 	= 1
	Vertical 	Track 	= 2
	Intersect 	Track 	= 3
	Collission  Track   = 4
	Empty		Track   = 5
	TurnLeftLean Track	= 6
	TurnRightLean Track	= 7



)

type Direction int
const (
	North	Direction 	= 0
	East 	Direction	= 1
	South 	Direction 	= 2
	West 	Direction 	= 3

)

func (current Direction) turnLeftLean() Direction {
	switch current {
		case East: return South
		case North: return West
		case South: return East
		case West: return North
	}
	panic("Unknown direction")
}

func (current Direction) turnRightLean() Direction {
	switch current {
	case East: return North
	case North: return East
	case South: return West
	case West: return South
	}
	panic("Unknown direction")
}

func (current Direction) intersect(intersectCount int) Direction {
	switch intersectCount % 3 {
	case 0:
		if current == East || current == West {
			return current.turnRightLean()
		} else {
			return current.turnLeftLean()
		}
	case 1:
		return current
	case 2:
		if current == East || current == West {
			return current.turnLeftLean()
		} else {
			return current.turnRightLean()
		}
	}
	panic("Math is broken")
}

type Coord struct {
	x, y int
}


type Cart struct {
	location Coord
	direction Direction
	intersectCount int
}

type ByCoords []*Cart

func (a ByCoords) Len() int { return len(a)}
func (a ByCoords) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoords) Less(i, j int) bool {
	if a[i].location.y == a[j].location.y {
		return a[i].location.x < a[j].location.x
	}
	return a[i].location.y < a[j].location.y
}


func createCart(char string) *Cart {
	switch char {
	case "<":
		return &Cart{direction: West}
	case ">":
		return &Cart{direction: East}
	case "^":
		return &Cart{direction: North}
	case "v":
		return &Cart{direction: South}
	}
	return nil
}

func createTrack(char string) Track {
	switch char {
	case "|":
		return Horizontal
	case "-":
		return Vertical
	case "/":
		return TurnRightLean
	case "\\":
		return TurnLeftLean
	case "+":
		return Intersect
	default:
		return Empty
	}
}

func (orig Coord) move(direction Direction) Coord {
	switch direction {
	case East:  return Coord{x:orig.x + 1, y:orig.y}
	case West:  return Coord{x:orig.x - 1, y:orig.y}
	case North: return Coord{x:orig.x, y:orig.y - 1}
	case South: return Coord{x:orig.x, y:orig.y + 1}
	}
	panic("Unknown direction passed")
}

func moveCart(cart *Cart, railroads map[Coord]Track) {
	newloc := cart.location.move(cart.direction)
	newTrack := railroads[newloc]
	newDirection := cart.direction
	switch newTrack {
	case TurnLeftLean:
		newDirection = cart.direction.turnLeftLean()
	case TurnRightLean:
		newDirection = cart.direction.turnRightLean()
	case Intersect:
		newDirection = cart.direction.intersect(cart.intersectCount)
		cart.intersectCount++
	}
	cart.direction = newDirection
	cart.location = newloc
}

var debug = false
func main() {
	railroads := make(map[Coord]Track)
	var carts []*Cart
	filename := "13/input.txt"
	if debug {
		filename = "13/input.test2.txt"
	}
	lines := readInput(filename)
	for y, line := range lines {
		for x , char := range line {
			coord := Coord{x: x, y: y}
			if cart := createCart(string(char)); cart != nil {
				cart.location = coord
				carts = append(carts, cart)
				track := Vertical
				if cart.direction == East || cart.direction == West {
					track = Horizontal
				}
				railroads[coord] = track

			} else if track := createTrack(string(char)); track != Empty {
				railroads[coord] = track
			}
		}
	}

	cartCount := len(carts)
	// let's start ticking
	outer:
	for tick := 0; tick < 10000000; tick++ {
		fmt.Println("== Tick:", tick)
		inner:
		for _, cart := range carts {
			if cart == nil {
				continue inner
			}
			fmt.Printf("%d %d %d %d\n", cart.location.x, cart.location.y, cart.direction % 4, cart.intersectCount % 3)
			moveCart(cart, railroads)
			//fmt.Printf("  Moved cart %d: %+v\n", cartId, cart )
			fmt.Printf("%d %d %d %d\n", cart.location.x, cart.location.y, cart.direction % 4, cart.intersectCount % 3)
			if cartCount == 1 {
				//fmt.Println("Done at tick:", tick)
				//fmt.Println("Remaining cart: ", cart)
				break outer
			}
			if collission, cart1, cart2 := detectCollission(carts); collission {
				carts[cart1] = nil
				carts[cart2] = nil
				cartCount -= 2
				fmt.Printf("!! Found colission at %+v, carts remaining: %d\n", cart.location, cartCount)

			}
		}
		carts = sortCartsAndPrune(carts);
	}
}


func detectCollission(carts []*Cart) (result bool, cart1, cart2 int) {
	coordMap := make(map[Coord]int)
	for index, cart := range carts {
		if cart != nil {
			if prevId, exists := coordMap[cart.location]; exists {
				return true, prevId, index
			} else {
				coordMap[cart.location] = index
			}
		}
	}
	return false, -1, -1
}

func sortCartsAndPrune(oldCarts []*Cart) []*Cart {
	var newCarts []*Cart
	for _, cart := range oldCarts {
		if cart != nil {
			newCarts = append(newCarts, cart)
		}
	}
	sort.Sort(ByCoords(newCarts))
	return newCarts
}
