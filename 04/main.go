package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
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

func main() {
	var splitRe = regexp.MustCompile("\\[([0-9]{4})-([0-9]{2})-([0-9]{2}) ([0-9]{2}):([0-9]{2})] (.+)")
	var guardRe = regexp.MustCompile("Guard #([0-9]+)")

	guards := make(map[int]map[int]int)
	minutes := make(map[int]map[int]int)
	for i := 0; i < 60; i++ {
		minutes[i] = make(map[int]int)
	}
	currentGuard := 0
	startSleep := 0

	lines := readInput("04/input.txt")
	sort.Strings(lines)
	for _, line := range lines {
		parts := splitRe.FindStringSubmatch(line)
		timestamp, _ := strconv.Atoi(parts[1] + parts[2] + parts[3] + parts[4] + parts[5])
		text := parts[6]
		switch text {
		case "falls asleep":
			startSleep = timestamp
		case "wakes up":
			for i := startSleep; i < timestamp; i++ {
				guards[currentGuard][i % 100]++
				minutes[i % 100][currentGuard]++
			}
		default:
			guardParts := guardRe.FindStringSubmatch(text)
			currentGuard, _ = strconv.Atoi(guardParts[1])
			_, exists := guards[currentGuard]
			if !exists {
				guards[currentGuard] = make(map[int]int)
			}
		}
	}
	prevTotalTime := 0
	mostSleepyGuard := 0
	for guard, minutes := range guards {
		totalTime := 0
		for _, occurences := range minutes {
			totalTime += occurences
		}
		if totalTime > prevTotalTime {
			mostSleepyGuard = guard
			prevTotalTime = totalTime
		}
	}

	mostSleepyMinute := 0
	mostOccurrences := 0
	for minute, occurrences := range guards[mostSleepyGuard] {
		if occurrences > mostOccurrences {
			mostSleepyMinute = minute
			mostOccurrences = occurrences
		}
	}
	fmt.Println("Most sleepy guard:", mostSleepyGuard, "Most sleepy minute:", mostSleepyMinute, "with", mostOccurrences, "occurences, total:", mostSleepyGuard * mostSleepyMinute)

	mostOccurrences = 0
	mostSleepyGuard = 0
	mostSleepyMinute = 0
	for minute, guards := range minutes {
		for guard, occurrences := range guards {
			if occurrences > mostOccurrences {
				mostOccurrences = occurrences
				mostSleepyGuard = guard
				mostSleepyMinute = minute
			}
		}
	}
	fmt.Println("Most sleepy guard:", mostSleepyGuard, "Most sleepy minute:", mostSleepyMinute, "with", mostOccurrences, "occurences, total:", mostSleepyGuard * mostSleepyMinute)
}
