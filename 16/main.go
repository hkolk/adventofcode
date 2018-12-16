package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
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

func intersection(a []string, b []string) (inter []string) {
	// interacting on the smallest list first can potentailly be faster...but not by much, worse case is the same
	low, high := a, b
	if len(a) > len(b) {
		low = b
		high = a
	}

	done := false
	for i, l := range low {
		for j, h := range high {
			// get future index values
			f1 := i + 1
			f2 := j + 1
			if l == h {
				inter = append(inter, h)
				if f1 < len(low) && f2 < len(high) {
					// if the future values aren't the same then that's the end of the intersection
					if low[f1] != high[f2] {
						done = true
					}
				}
				// we don't want to interate on the entire list everytime, so remove the parts we already looped on will make it faster each pass
				high = high[:j+copy(high[j:], high[j+1:])]
				break
			}
		}
		// nothing in the future so we are done
		if done {
			break
		}
	}
	return
}

// Intersect returns a slice of values that are present in all of the input slices
//
// [1, 1, 3, 4, 5, 6] & [2, 3, 6] >> [3, 6]
//
// [1, 1, 3, 4, 5, 6] >> [1, 3, 4, 5, 6]
func Intersect(arrs ...interface{}) (reflect.Value, bool) {
	// create a map to count all the instances of the slice elems
	arrLength := len(arrs)
	var kind reflect.Kind
	var kindHasBeenSet bool

	tempMap := make(map[interface{}]int)
	for _, arg := range arrs {
		tempArr, ok := Distinct(arg)
		if !ok {
			return reflect.Value{}, ok
		}

		// check to be sure the type hasn't changed
		if kindHasBeenSet && tempArr.Len() > 0 && tempArr.Index(0).Kind() != kind {
			return reflect.Value{}, false
		}
		if tempArr.Len() > 0 {
			kindHasBeenSet = true
			kind = tempArr.Index(0).Kind()
		}

		c := tempArr.Len()
		for idx := 0; idx < c; idx++ {
			// how many times have we encountered this elem?
			if _, ok := tempMap[tempArr.Index(idx).Interface()]; ok {
				tempMap[tempArr.Index(idx).Interface()]++
			} else {
				tempMap[tempArr.Index(idx).Interface()] = 1
			}
		}
	}

	// find the keys equal to the length of the input args
	numElems := 0
	for _, v := range tempMap {
		if v == arrLength {
			numElems++
		}
	}
	out := reflect.MakeSlice(reflect.TypeOf(arrs[0]), numElems, numElems)
	i := 0
	for key, val := range tempMap {
		if val == arrLength {
			v := reflect.ValueOf(key)
			o := out.Index(i)
			o.Set(v)
			i++
		}
	}

	return out, true
}

// Distinct returns the unique vals of a slice
//
// [1, 1, 2, 3] >> [1, 2, 3]
func Distinct(arr interface{}) (reflect.Value, bool) {
	// create a slice from our input interface
	slice, ok := takeArg(arr, reflect.Slice)
	if !ok {
		return reflect.Value{}, ok
	}

	// put the values of our slice into a map
	// the key's of the map will be the slice's unique values
	c := slice.Len()
	m := make(map[interface{}]bool)
	for i := 0; i < c; i++ {
		m[slice.Index(i).Interface()] = true
	}
	mapLen := len(m)

	// create the output slice and populate it with the map's keys
	out := reflect.MakeSlice(reflect.TypeOf(arr), mapLen, mapLen)
	i := 0
	for k := range m {
		v := reflect.ValueOf(k)
		o := out.Index(i)
		o.Set(v)
		i++
	}

	return out, ok
}

func takeArg(arg interface{}, kind reflect.Kind) (val reflect.Value, ok bool) {
	val = reflect.ValueOf(arg)
	if val.Kind() == kind {
		ok = true
	}
	return
}


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

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

type TrainData struct {
	before []int
	after []int
	opcode int
	a, b, c int
}

func matching(data TrainData, opcodeMap map[string]interface{}) (result []string) {
	for name, function := range opcodeMap {
		//fmt.Println(before)
		after := function.(func(int, int, int, []int)([]int))(data.a, data.b, data.c, append(data.before[:0:0], data.before...))
		//fmt.Println("After", name, "registers became", after)
		if Equal(after, data.after) {
			//fmt.Println(name, "is a match")
			result = append(result, name)
		}
	}
	return
}

func main() {
	opcodeMap := map[string]interface{}{
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

	quickTrain := TrainData{before:[]int{3, 2, 1, 1}, after: []int{3, 2, 2, 1}, opcode: 9, a: 2, b: 1, c: 2}
	result := matching(quickTrain, opcodeMap)
	fmt.Println(result)

	potentials := make(map[int][]string)
	for i := 0; i < 16; i++ {
		potentials[i] = []string{}
		for key, _ := range opcodeMap {
			potentials[i] = append(potentials[i], key)
		}
	}

	lines := readInput("16/input.train.txt")
	var train TrainData
	var allTraining []TrainData
	for _, line := range lines {
		var p1, p2, p3, p4 int
		_, error := fmt.Sscanf(line, "Before: [%d, %d, %d, %d]", &p1, &p2, &p3, &p4)
		if error == nil {
			register := []int{p1, p2, p3, p4}
			train = TrainData{before: register}
		}
		_, error = fmt.Sscanf(line, "%d %d %d %d", &p1, &p2, &p3, &p4)
		if error == nil {
			train.opcode = p1
			train.a = p2
			train.b = p3
			train.c = p4
		}
		_, error = fmt.Sscanf(line, "After: [%d, %d, %d, %d]", &p1, &p2, &p3, &p4)
		if error == nil {
			register := []int{p1, p2, p3, p4}
			train.after = register
			allTraining = append(allTraining, train)
		}
	}
	//fmt.Println(allTraining)

	hits := 0
	for _, trainData := range allTraining {
		result := matching(trainData, opcodeMap)
		if len(result) >= 3 {
			hits++
		}
		var newResult []string
		for _, potential := range potentials[trainData.opcode] {
			for _, actual := range result {
				if potential == actual {
					newResult = append(newResult, actual)
				}
			}
		}
		potentials[trainData.opcode] = newResult
	}
	actualOpcodes := make(map[int]string)
	found := make(map[string]bool)
	for len(actualOpcodes) != 16 {
		for opcode, options := range potentials {
			var postFilter []string
			for _, option := range options {
				if !found[option] {
					postFilter = append(postFilter, option)
				}
			}
			if len(postFilter) == 1 {
				actualOpcodes[opcode] = postFilter[0]
				potentials[opcode] = []string{}
				found[postFilter[0]] = true
			}
		}
	}
	fmt.Printf("Found %d hits for 3 results\n", hits)
	for i := 0; i < 16; i++ {
		fmt.Printf("%d: %+v\n", i, actualOpcodes[i])
	}

	program := readInput("16/input.program.txt")
	var programLines []ProgramLine
	for _, line := range program {
		var p1, p2, p3, p4 int
		_, _ = fmt.Sscanf(line, "%d %d %d %d", &p1, &p2, &p3, &p4)
		programLine := ProgramLine{opcode: p1, a: p2, b: p3, c: p4}
		programLines = append(programLines, programLine)
	}
	registers := []int{0, 0, 0, 0}
	for _, line := range programLines {
		opcodeString := actualOpcodes[line.opcode]
		registers = execute(opcodeString, line, opcodeMap, registers)
	}
	fmt.Println("Register 0 now contains:", registers[0])
}

func execute(opcodeString string, line ProgramLine, opcodeMap map[string]interface{}, registers []int) []int {
	after := opcodeMap[opcodeString].(func(int, int, int, []int)([]int))(line.a, line.b, line.c, append(registers[:0:0], registers...))
	return after
}

type ProgramLine struct {
	opcode int
	a, b, c int
}