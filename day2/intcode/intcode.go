package intcode

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var debugLog *log.Logger //nolint:gochecknoglobals

func Init(logger *log.Logger) {
	debugLog = logger
}

func Execute(memory []int) []int {
	result := make([]int, len(memory))
	copy(result, memory)

	for i := 0; i < len(result); i += 4 {
		opcode := Opcode(result[i])
		switch opcode {
		case Add:
			p1 := result[i+1]
			p2 := result[i+2]
			p3 := result[i+3]

			debugLog.Printf("Adding %d and %d, storing at position %d\n", result[p1], result[p2], p3)
			result[p3] = result[p1] + result[p2]
		case Multiply:
			p1 := result[i+1]
			p2 := result[i+2]
			p3 := result[i+3]

			debugLog.Printf("Multiplying %d and %d, storing at position %d\n", result[p1], result[p2], p3)
			result[p3] = result[p1] * result[p2]
		case Halt:
			debugLog.Println("Halting execution")

			return result
		default:
			log.Fatalf("unknown opcode %d at position %d", opcode, i)
		}
	}

	return result
}

func Restore1202ProgramAlarm(memory []int) []int {
	memory[1] = 12
	memory[2] = 2

	return memory
}

func FindNounAndVerbForOutput(initialMemory []int, targetOutput int) (int, int) {
	memory := make([]int, len(initialMemory))

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(memory, initialMemory)
			initialMemory[1] = noun
			initialMemory[2] = verb
			result := Execute(initialMemory)[0]

			if result == targetOutput {
				return noun, verb
			}
		}
	}

	return -1, -1
}

func ReadMemoryFromFile(filename string) []int {
	memory := []int{}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()

		for _, s := range strings.Split(str, ",") {
			instruction, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			memory = append(memory, instruction)
		}
	}

	return memory
}

func MemoryEquals(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
