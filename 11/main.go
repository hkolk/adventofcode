package main

import "fmt"

func makeGrid(serial int) map[int]map[int]int {
	grid := make(map[int]map[int]int)
	for x := 1; x <= 300; x++ {
		grid[x] = make(map[int]int)
		for y := 1; y <= 300; y++ {
			rackId := x + 10
			power := ((rackId * y) + serial) * rackId
			power = (power / 100) % 10
			grid[x][y] = power - 5
		}
	}
	return grid
}

func highestPower(grid map[int]map[int]int, squareSize int) (maxX, maxY, power int) {
	highestSum := 0
	for x := 1; x <= 301-squareSize; x++ {
		for y := 1; y <= 301-squareSize; y++ {
			sum := 0
			for i := 0; i < squareSize; i++ {
				for j := 0; j < squareSize; j++ {
					sum += grid[x+i][y+j]
				}
			}
			if sum > highestSum {
				highestSum = sum
				maxX = x;
				maxY = y;
				power = highestSum
			}
		}
	}
	return
}

func main() {
	grid := makeGrid(57)
	fmt.Println(grid[122][79])

	grid = makeGrid(39)
	fmt.Println(grid[217][196])

	grid = makeGrid(71)
	fmt.Println(grid[101][153])

	fmt.Println(highestPower(makeGrid(18), 3))
	fmt.Println(highestPower(makeGrid(42), 3))
	fmt.Println(highestPower(makeGrid(4151), 3))

	grid = makeGrid(4151)
	var maxX, maxY, maxSize, maxPower int
	// Let's hope we find it in the first 30
	for squareSize := 1; squareSize <= 30; squareSize++ {
		fmt.Println("Loop:", squareSize)
		x, y, power := highestPower(grid, squareSize)
		if power > maxPower {
			maxPower = power
			maxX = x
			maxY = y
			maxSize = squareSize
		}
		fmt.Println(maxX, maxY, maxSize, maxPower)
	}
}
