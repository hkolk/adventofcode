package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sort"
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
	name string
	upstream map[string]node
	downstream map[string]node
	processing bool
}

type worker struct {
	id int
	working bool
	taskName string
	workingTime int
}

var debug = false
func main() {
	filename := "07/input.txt"
	workerCount := 5
	timePerTask := 60
	if debug {
		filename = "07/input.test.txt"
		workerCount = 2
		timePerTask = 0
	}
	//part1(createNodes(readInput(filename)))
	part2(createNodes(readInput(filename)), workerCount, timePerTask)

}

func part2(nodes map[string]*node, workerCount int, timePerTask int) {
	printNodes(nodes)

	sb := strings.Builder{}

	var workers []*worker
	for i := 0; i < workerCount; i++ {
		workers = append(workers, &worker{id: i})
	}

	loopCounter := 0

	for {
		// Tick loop
		for _ , worker := range workers {
			if worker.working {
				if worker.workingTime == 0 {
					name := worker.taskName

					for _, node := range nodes[name].downstream {
						delete(node.upstream, name)
					}
					delete(nodes, name)
					sb.WriteString(name)
					worker.working = false
					worker.taskName = ""
					worker.workingTime = 0
				} else {
					worker.workingTime--
				}
			}
		}
		for _ , worker := range workers {
			if !worker.working {
				// take work?
				var processing []string
				for _, node := range nodes {
					if len(node.upstream) == 0 && !node.processing {
						processing = append(processing, node.name)
					}
				}
				if len(processing) != 0 {
					//fmt.Println(processing)
					sort.Strings(processing)
					name := processing[0]
					nodes[name].processing = true
					worker.working = true
					worker.taskName = name
					worker.workingTime = (timePerTask - 1) + int(name[0] - 64)
				}
			}
		}
		fmt.Print(loopCounter, "\t")
		for _, worker := range workers {
			fmt.Print(worker.taskName, "\t")
		}
		fmt.Println(sb.String())
		loopCounter++
		//fmt.Println("Loopcounter: ", loopCounter)
		//printNodes(nodes)
		//printWorkers(workers)
		if debug && loopCounter > 100 {
			break
		}
		if len(nodes) == 0 {
			break
		}

	}
	fmt.Println("Result:", sb.String(), "in", loopCounter - 1, "seconds")
}

func printWorkers(workers []worker) {
	for _, worker := range workers {
		fmt.Printf("%+v \n", worker)
	}
}
func part1(nodes map[string]*node) {
	printNodes(nodes)

	sb := strings.Builder{}

	for {
		// find fulfilled
		processing := []string{}
		for _, node := range nodes {
			if len(node.upstream) == 0 {
				processing = append(processing, node.name)
			}
		}
		if len(processing) == 0 {
			break
		}
		sort.Strings(processing)
		name := processing[0]
		sb.WriteString(name)
		fmt.Println("Hit: ", name)
		for _, node := range nodes[name].downstream {
			delete(node.upstream, name)
		}
		delete(nodes, name)
		for name, node := range nodes {
			fmt.Println("Name:", name, ", upstream", reflect.ValueOf(node.upstream).MapKeys(), ", downstream:", reflect.ValueOf(node.downstream).MapKeys())
		}
	}
	fmt.Println("Result:", sb.String())
}

func printNodes(nodes map[string]*node) {
	for name, node := range nodes {
		fmt.Println("Name:", name, ", upstream", reflect.ValueOf(node.upstream).MapKeys(), ", downstream:", reflect.ValueOf(node.downstream).MapKeys())
	}
}

func createNodes(lines []string) map[string]*node {
	nodes := make(map[string]*node)
	for _, line := range lines {
		var start, end string
		fmt.Sscanf(line, "Step %1s must be finished before step %1s can begin.", &start, &end)
		fmt.Println(start, end)
		_, exists := nodes[start]
		if !exists {
			nodes[start] = &node{name: start, upstream: make(map[string]node), downstream: make(map[string]node)}
		}
		_, exists = nodes[end]
		if !exists {
			nodes[end] = &node{name: end, upstream: make(map[string]node), downstream: make(map[string]node)}
		}

		nodes[start].downstream[end] = *nodes[end]
		nodes[end].upstream[start] = *nodes[start]
	}
	return nodes
}
