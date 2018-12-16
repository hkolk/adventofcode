package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	next *Node
	prev *Node
	recipeScore int
}

func (current *Node) append(recipeScore int) {
	newNode := &Node{recipeScore: recipeScore, next: current.next, prev: current}
	current.next.prev = newNode
	current.next = newNode
}

func (current *Node) print() {
	fmt.Print(current.recipeScore, " ")
	pointer := current.next
	for pointer != current {
		fmt.Print(pointer.recipeScore, " ")
		pointer = pointer.next
	}
	fmt.Println()
}

func (current *Node) count() int {
	count := 1
	pointer := current.next
	for pointer != current {
		count += 1
		pointer = pointer.next
	}
	return count
}

func (current *Node) getScore(skip int) string {
	pointer := current
	for skip > 0 {
		pointer = pointer.next
		skip--
	}
	sb := strings.Builder{}
	for i := 0; i < 10; i++ {
		sb.WriteString(strconv.Itoa(pointer.recipeScore))
		pointer = pointer.next
	}
	return sb.String()
}

func (current *Node) getTailAsString() string {
	sb := strings.Builder{}
	pointer := current;
	for i := 0; i < 7; i++ {
		pointer = pointer.prev
	}
	for i := 0; i < 7; i++ {
		sb.WriteString(strconv.Itoa(pointer.recipeScore))
		pointer = pointer.next
	}
	return sb.String()
}

type Elf struct {
	id int
	current *Node
}

func (elf *Elf) advance() {
	steps := 1 + elf.current.recipeScore
	for i := 0; i < steps; i++ {
		elf.current = elf.current.next
	}
}

func intToArray(input int) (output []int) {
	if input == 0 {
		output = append(output, 0)
	}
	for input > 0 {
		output = append(output, input % 10)
		input /= 10
	}
	for left, right := 0, len(output)-1; left < right; left, right = left+1, right-1 {
		output[left], output[right] = output[right], output[left]
	}
	return
}

func run(stopAfter int, debug bool) string {
	// setup
	head := &Node{recipeScore:3}
	node2 := &Node{recipeScore:7, next: head, prev:head}
	head.prev = node2
	head.next = node2

	elfs := []*Elf{
		&Elf{id: 1, current: head},
		&Elf{id: 2, current: node2}}


	//for round := 1; round < 10; round++ {
	round := 0

	for head.count() < stopAfter + 10 {
		round++
		sum := 0
		for _, elf := range elfs {
			sum += elf.current.recipeScore
		}
		if debug { fmt.Println("Sum for round", round, "is", sum) }
		digits := intToArray(sum)
		for _, digit := range digits {
			head.prev.append(digit)
		}
		if debug { head.print() }

		for _, elf := range elfs {
			elf.advance()
			if debug { fmt.Println("Moved elf", elf.id, "to node with score", elf.current.recipeScore) }
		}
		if debug { fmt.Println("Number of nodes:", head.count()) }

	}
	return head.getScore(stopAfter)

}

func runB(stopWhenFound string, debug bool) int {
	// setup
	head := &Node{recipeScore:3}
	node2 := &Node{recipeScore:7, next: head, prev:head}
	head.prev = node2
	head.next = node2

	elfs := []*Elf{
		&Elf{id: 1, current: head},
		&Elf{id: 2, current: node2}}


	//for round := 1; round < 10; round++ {
	round := 0

	for {
		round++
		sum := 0
		for _, elf := range elfs {
			sum += elf.current.recipeScore
		}
		if debug { fmt.Println("Sum for round", round, "is", sum) }
		digits := intToArray(sum)
		for _, digit := range digits {
			head.prev.append(digit)
		}
		if debug { head.print() }

		for _, elf := range elfs {
			elf.advance()
			if debug { fmt.Println("Moved elf", elf.id, "to node with score", elf.current.recipeScore) }
		}
		if debug { fmt.Println("Number of nodes:", head.count()) }
		//if head.count() > 170 {
		//	os.Exit(-1)
		//}
		if debug { fmt.Println("Tail: ", head.getTailAsString()) }
		tail := head.getTailAsString()
		if strings.Contains(tail, stopWhenFound) {
			return head.count() - 7 + strings.Index(tail, stopWhenFound)
		}
	}
}

func main() {
	fmt.Println(run(5, false))
	fmt.Println(run(18, false))
	fmt.Println(run(2018, false))
	fmt.Println(run(47801, false))

	fmt.Println(runB("51589", false))
	fmt.Println(runB("01245", false))
	fmt.Println(runB("92510", false))
	fmt.Println(runB("59414", false))
	fmt.Println(runB("047801", false))


}
