package intcode

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"math"
)

var debugLog *log.Logger //nolint:gochecknoglobals

func Init(logger *log.Logger) {
	debugLog = logger
}

type CPU struct {
	IP int
	Code int
	Modes []Mode
	Parameters []int
	JUMP bool
}

var cpu CPU
var _memory []int

func Execute(memory []int) []int {
	_memory = make([]int, len(memory))
	copy(_memory, memory)

	// I guess I mean #inputs + #parameters
	var nparams int
	for cpu.IP = 0; cpu.IP < len(_memory); {
		cpu.Code = _memory[cpu.IP]
		opcode := parseCode(cpu.Code)
		switch opcode {
		case ADD:
			nparams = 3
			padModes(nparams)
			params := setParameters(nparams)
			debugState()
			debugLog.Printf("ADD")
			debugLog.Printf("memory[%d] = %d + %d\n",
				params[2],
				params[0],
				params[1])

			_memory[params[2]] = params[0] + params[1]

		case MULT:
			nparams = 3
			padModes(nparams)
			params := setParameters(nparams)
			debugState()
			debugLog.Printf("MULT")
			debugLog.Printf("memory[%d] = %d * %d\n",
				params[2],
				params[0],
				params[1])

			_memory[params[2]] = params[0] * params[1]

		case INPUT:
			nparams = 1
			debugState()
			debugLog.Printf("INPUT")
			var input int
			print("> ")
			fmt.Scan(&input)
			out_ptr := getValue(1)
			debugLog.Printf("memory[%d] = %d\n", out_ptr, input)
			_memory[out_ptr] = input

		case OUT:
			nparams = 1
			params := setParameters(nparams)
			debugState()
			debugLog.Printf("OUT")
			println(">>", params[0])
			
		case JIT:
			nparams = 2
			params := setParameters(nparams)
			debugState()
			debugLog.Printf("JIT")
			if params[0] != 0 {
				debugLog.Printf("JUMP %d", params[1])
				cpu.IP = params[1]
				cpu.JUMP = true
			} else {
				debugLog.Printf("NOP")
			}
			
		case JIF:
			nparams = 2
			params := setParameters(nparams)
			debugState()
			debugLog.Printf("JIF")
			if params[0] == 0 {
				debugLog.Printf("JUMP %d", params[1])
				cpu.IP = params[1]
				cpu.JUMP = true
			} else {
				debugLog.Printf("NOP")
			}
			
		case LT:
			nparams = 3
			padModes(nparams)
			debugState()
			debugLog.Printf("LT")
			params := setParameters(nparams)
			if params[0] < params[1] {
				_memory[params[2]] = 1
			} else {
				_memory[params[2]] = 0
			}

		case EQ:
			nparams = 3
			padModes(nparams)
			params := setParameters(nparams)
			debugState()
			debugLog.Printf("EQ")
			if params[0] == params[1] {
				_memory[params[2]] = 1
			} else {
				_memory[params[2]] = 0
			}

		case HALT:
			debugLog.Println("HALT received, stopping execution")
			return _memory

		default:
			debugState()
			log.Fatalf("unknown opcode %d at position %d", opcode, cpu.IP)
		}

		if cpu.JUMP {
			cpu.JUMP = false
		} else {
			cpu.IP += nparams+1
		}
	}

	return _memory
}

func padModes(nparams int) {
	if len(cpu.Modes) < nparams {
		for i := len(cpu.Modes); i < nparams -1; i++ {
			cpu.Modes = append(cpu.Modes, 0)
		}
		// we need to not dereference the output
		// so we get the value of the pointer
		cpu.Modes = append(cpu.Modes, 1)
	}
}

func getStar(parameter int) int {
	return getParameter(_memory, cpu.IP, parameter, Position)
}

func getValue(parameter int) int {
	return getParameter(_memory, cpu.IP, parameter, Immediate)
}

func getParameter(memory []int, index int, parameter int, mode Mode) int {
	var res int

	switch mode {
	case Position:
		res = memory[memory[index+parameter]]
	case Immediate:
		res = memory[index+parameter]
	default:
		log.Fatalf("unknown mode %d", mode)
	}

	return res
}


func setParameters(n_parameters int) []int {
	if len(cpu.Modes) > n_parameters {
		debugLog.Printf("number of modes: %d, number of parameters: %d", 
			len(cpu.Modes), n_parameters)
		log.Fatal("number of modes must less or equal to parameters")
	}
	params := make([]int, n_parameters)
	for i := range n_parameters {
		if i < len(cpu.Modes) {
			switch cpu.Modes[i] {
			case Position:
				params[i] = _memory[_memory[cpu.IP+i+1]]
			case Immediate:
				params[i] = _memory[cpu.IP+i+1]
			default:
				log.Fatalf("unknown mode %d", cpu.Modes[i])
			}
		} else { // omitted modes are Position
			params[i] = _memory[_memory[cpu.IP+i+1]]
		}
	}
	cpu.Parameters = params
	return params
}

func parseCode(code int) Opcode {
	length := intLength(code)
	cpu.Modes = make([]Mode,0)

	// first 2 digits are opcode
	opcode := Opcode(code%100)
	// rest of the digits are modes for each parameter
	// leftmost 0s are omitted
	var leftmost int
	for i := range length-2 {
		leftmost = code/int(math.Pow10(i+2))
		cpu.Modes = append(cpu.Modes, Mode((leftmost)%10))
	}
	
	return opcode
}

func intLength(n int) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n != 0 {
		n /= 10
		count++
	}

	return count
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

func debugState() {
	memory_slice := make([]int, 0)
	for i := range len(cpu.Parameters)+1 {
		memory_slice = append(memory_slice, _memory[cpu.IP+i])
	}
	debugLog.Print("IP: ",cpu.IP,"    ",memory_slice)
}
