package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func addr(a int, b int, c int, registers []int) []int {
	registers[c] = registers[a] + registers[b]
	return registers
}

func addi(a int, b int, c int, registers []int) []int {
	registers[c] = registers[a] + b
	return registers
}

func mulr(a int, b int, c int, registers []int) []int {
	registers[c] = registers[a] * registers[b]
	return registers
}

func muli(a int, b int, c int, registers []int) []int {
	registers[c] = registers[a] * b
	return registers
}

func banr(a, b, c int, registers []int) []int {
	registers[c] = registers[a] & registers[b]
	return registers
}

func bani(a, b, c int, registers []int) []int {
	registers[c] = registers[a] & b
	return registers
}

func borr(a, b, c int, registers []int) []int {
	registers[c] = registers[a] | registers[b]
	return registers
}
func bori(a, b, c int, registers []int) []int {
	registers[c] = registers[a] | b
	return registers
}
func setr(a, b, c int, registers []int) []int {
	registers[c] = registers[a]
	return registers
}
func seti(a, b, c int, registers []int) []int {
	registers[c] = a
	return registers
}

func gtir(a, b, c int, registers []int) []int {
	if a > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}
func gtri(a, b, c int, registers []int) []int {
	if registers[a] > b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}
func gtrr(a, b, c int, registers []int) []int {
	if registers[a] > registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}
func eqir(a, b, c int, registers []int) []int {
	if a == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}
func eqri(a, b, c int, registers []int) []int {
	if registers[a] == b {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}
func eqrr(a, b, c int, registers []int) []int {
	if registers[a] == registers[b] {
		registers[c] = 1
	} else {
		registers[c] = 0
	}
	return registers
}

var opcodeMap = map[string]interface{}{
"addr": addr,
"addi": addi,
"mulr": mulr,
"muli": muli,
"banr": banr,
"bani": bani,
"borr": borr,
"bori": bori,
"setr": setr,
"seti": seti,
"gtir": gtir,
"gtri": gtri,
"gtrr": gtrr,
"eqir": eqir,
"eqri": eqri,
"eqrr": eqrr }

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

type ProgramLine struct {
	opcode string
	a, b, c int
}

func (instruction ProgramLine) execute(registers []int) []int {
	after := opcodeMap[instruction.opcode].(func(int, int, int, []int)([]int))(instruction.a, instruction.b, instruction.c, append(registers[:0:0], registers...))
	return after
}

func main()  {
	lines := readInput("19/input.txt")
	var instructions []ProgramLine
	var ip int
	for _, line := range lines {
		var p1 string
		var p2, p3, p4 int
		_, error := fmt.Sscanf(line, "%s %d %d %d", &p1, &p2, &p3, &p4)
		if error == nil {
			instructions = append(instructions, ProgramLine{opcode:p1, a:p2, b:p3, c:p4})
		} else {
			fmt.Sscanf(line, "#ip %d", &ip)
		}
	}
	fmt.Println(ip, instructions)

	registers := []int{1, 0, 0, 0, 0, 0}
	debug := false
	maxTick := 100000000
	if debug {
		maxTick = 100000
	}
	for tick := 0; tick < maxTick; tick++ {
		if registers[ip] == 3 {
			factor := registers[3]
			hitcount := 0
			for i := 1; i <= factor; i++ {
				if factor % i == 0 {
					hitcount += i
					fmt.Println(hitcount)
				}

			}
			fmt.Println("Result: ", hitcount)
			break
		}

		if registers[ip] < 0 || registers[ip] >= len(instructions) {
			fmt.Println("Result: ", registers[0])
			break
		}
		if debug { fmt.Printf("ip:%d %+v %+v ", registers[ip], registers, instructions[registers[ip]]) }
		if registers[ip] == 9 {
			//registers[5] = registers[3] + 1
		}
		registers = instructions[registers[ip]].execute(registers)
		if debug { fmt.Printf("=> %+v\n", registers) }
		registers[ip]++
	}



}
