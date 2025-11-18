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

			result[p3] = result[p1] + result[p2]

			debugLog.Printf("%d + %d = %d -> %d\n", result[p1], result[p2], result[p3], p3)
		case Multiply:
			p1 := result[i+1]
			p2 := result[i+2]
			p3 := result[i+3]

			result[p3] = result[p1] * result[p2]

			debugLog.Printf("%d * %d = %d -> %d\n", result[p1], result[p2], result[p3], p3)
		case Halt:
			debugLog.Println("HALT received, stopping execution")

			return result
		default:
			log.Fatalf("unknown opcode %d at position %d", opcode, i)
		}
	}

	return result
}

func Restore1202ProgramAlarm(memory []int) []int {
	debugLog.Println("Restoring 1202 program alarm state")
	memory[1] = 12
	memory[2] = 2

	return memory
}

const (
	minNoun = 0
	maxNoun = 99
	minVerb = 0
	maxVerb = 99
)

func FindNounAndVerbForOutput(initialMemory []int, targetOutput int) (int, int) {
	debugLog.Printf("Finding noun and verb for target output %d\n", targetOutput)

	for noun := minNoun; noun <= maxNoun; noun++ {
		for verb := minVerb; verb <= maxVerb; verb++ {
			memory := make([]int, len(initialMemory))
			copy(memory, initialMemory)
			memory[1] = noun
			memory[2] = verb
			result := Execute(memory)[0]

			if result == targetOutput {
				debugLog.Printf("Found noun %d and verb %d for target output %d\n", noun, verb, targetOutput)
				return noun, verb
			}
		}
	}

	return -1, -1
}

func ReadMemoryFromFile(filename string) (*[]int, error) {
	memory := []int{}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()

		for _, s := range strings.Split(str, ",") {
			instruction, err := strconv.Atoi(s)
			if err != nil {
				return nil, err
			}
			memory = append(memory, instruction)
		}
	}

	return &memory, nil
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
