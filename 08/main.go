package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

type node struct {
	id int
	childNodeCount int
	metadataCount int
	childNodes []*node
	metadata []int
	parent *node
}

func (self node) print() {
	fmt.Println("Id:", self.id, "Children:", len(self.childNodes), "Metadata:", self.metadata)
}

func (self node) getValueFromNode() int {
	if self.childNodeCount == 0 {
		sum := 0
		for _, num := range self.metadata {
			sum += num
		}
		return sum
	} else {
		sum := 0
		for _, num := range self.metadata {
			if num > 0 && num <= len(self.childNodes) {
				childNode := self.childNodes[num-1]
				sum += childNode.getValueFromNode()
			}
		}
		return sum
	}
}

var debug = false
func main() {
	filename := "08/input.txt"
	if debug {
		filename = "08/input.test.txt"
	}
	lines := readInput(filename)
	var numbers []int
	for _, part := range strings.Split(lines[0], " ") {
		number, _ := strconv.Atoi(part)
		numbers = append(numbers, number)
	}
	fmt.Println(numbers)

	metaDataCounter := 0
	rootNode := &node{id:0, childNodeCount:numbers[0], metadataCount:numbers[1]}
	currentNode := rootNode
	for index := 2; index < len(numbers); index++ {
		if currentNode.childNodeCount != len(currentNode.childNodes) {
			// Create a node
			fmt.Println("Creating node at index", index)
			node := node{id: index, childNodeCount: numbers[index], metadataCount: numbers[index+1]}
			node.parent = currentNode
			currentNode = &node
			index++
		} else if currentNode.metadataCount != len(currentNode.metadata) {
			// read metadata
			currentNode.metadata = append(currentNode.metadata, numbers[index])
			metaDataCounter += numbers[index]
		} else {
			// done processing node, jump back to parent
			index--
			currentNode.parent.childNodes = append(currentNode.parent.childNodes, currentNode)
			currentNode = currentNode.parent
		}
		currentNode.print()
	}
	fmt.Println("Metadata sum:", metaDataCounter)
	fmt.Println("Value of rootnode:", rootNode.getValueFromNode())
}
