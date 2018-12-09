package main

import "fmt"


type node struct {
	value int
	prev *node
	next *node
}

func printNodes(start *node) {
	cur := start
	fmt.Print(cur.value, " ")
	cur = cur.next
	safety := 0
	for cur != start && safety < 100 {
		fmt.Print(cur.value, " ")
		cur = cur.next
		safety++
	}
	fmt.Println()
}

func playFast(players int, marbleCount int) int {
	var currentNode *node
	currentNode = &node{value: 0}
	currentNode.next = currentNode
	currentNode.prev = currentNode

	scores := make([]int, players)
	for i := 1; i <= marbleCount; i++ {
		if i % 23 == 0 {
			score := i
			for x := 0; x < 7; x++ {
				currentNode = currentNode.prev
			}
			//fmt.Println("Removing item", currentNode.value)
			score += currentNode.value
			currentNode.prev.next = currentNode.next
			currentNode.next.prev = currentNode.prev
			currentNode = currentNode.next
			player := (i-1) % players
			scores[player] += score
		} else {
			//fmt.Println("Doing round", i)
			currentNode = currentNode.next
			newNode := &node{value: i, next: currentNode.next, prev: currentNode}
			newNode.prev.next = newNode
			newNode.next.prev = newNode
			currentNode = newNode
		}
		//printNodes(currentNode)
	}
	maxScore := 0
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func play(players int, marbleCount int) int {
	marbles := []int{0}
	currentIndex := 0
	scores := make([]int, players)
	for i := 1; i <= marbleCount; i++ {
		if i % 23 == 0 {
			score := i
			remove := (currentIndex + len(marbles) - 7) % len(marbles)
			score += marbles[remove]
			marbles = append(marbles[:remove], marbles[remove+1:]...)
			currentIndex = remove
			//fmt.Println(marbles)
			player := (i-1) % players
			scores[player] += score
			//fmt.Println("Player", player+1, "now has a score of", scores[player])
		} else {
			currentIndex = ((currentIndex + 1) % len(marbles)) + 1
			marbles = append(marbles[:currentIndex], append([]int{i}, marbles[currentIndex:]...)...)
			//fmt.Println(marbles)
		}
	}
	maxScore := 0
	for _, score := range scores {
		if score > maxScore {
			maxScore = score
		}
	}
	return maxScore
}

func main() {
	fmt.Println("Result: ", play(9, 25))
	fmt.Println("Result: ", playFast(9, 25))
	fmt.Println("Result: ", play(10, 1618))
	fmt.Println("Result: ", playFast(10, 1618))
	fmt.Println("Result: ", play(13, 7999))
	fmt.Println("Result: ", playFast(13, 7999))
	fmt.Println("Result: ", play(17, 1104))
	fmt.Println("Result: ", playFast(17, 1104))
	fmt.Println("Result: ", playFast(21, 6111))
	fmt.Println("Result: ", playFast(30, 5807))
	fmt.Println("Result: ", playFast(438, 71626))
	fmt.Println("Result: ", playFast(438, 71626*100))




}