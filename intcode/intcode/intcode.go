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
	_memory := make([]int, len(memory))
	copy(_memory, memory)

	// I guess I mean #inputs + #parameters
	var nparams int
	for i := 0; i < len(_memory); i += nparams+1 {
		opcode := Opcode(_memory[i])
		switch opcode {
		case Add:
			nparams = 3
			params := getParameters(_memory, i, nparams)
			_memory[params[2]] = _memory[params[0]] + _memory[params[1]]

			debugLog.Printf("%d + %d = %d -> %d\n",
				_memory[params[0]],
				_memory[params[1]],
				_memory[params[2]],
				params[2])
		case Multiply:
			nparams = 3
			params := getParameters(_memory, i, nparams)
			_memory[params[2]] = _memory[params[0]] * _memory[params[1]]

			debugLog.Printf("%d * %d = %d -> %d\n",
				_memory[params[0]],
				_memory[params[1]],
				_memory[params[2]],
				params[2])
		case Store:
			nparams = 2
			params := getParameters(_memory, i, nparams)
			_memory[params[1]] = _memory[params[0]]
		case Outputs:
			nparams = 1
			params := getParameters(_memory, i, nparams)
			println(params[0])
			
		case Halt:
			debugLog.Println("HALT received, stopping execution")

			return _memory
		default:
			log.Fatalf("unknown opcode %d at position %d", opcode, i)
		}
	}

	return _memory
}

func getParameters(memory []int, index int, n_parameters int) []int {
	params := make([]int, n_parameters)
	for i := range n_parameters {
		params = append(params, memory[index+i+1])
	}
	return params
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
