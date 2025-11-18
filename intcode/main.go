package main

import (
	"aoc2019/day2/intcode"
	"flag"
	"io"
	"log"
	"os"
)

var (
	verbose bool
)

var debugLog *log.Logger

func execute(memory []int) int {
	memory = intcode.Restore1202ProgramAlarm(memory)

	debugLog.Println("Initial memory:", memory)

	memory = intcode.Execute(memory)

	debugLog.Println("Final memory:", memory)

	return memory[0]
}

func findNounAndVerb(initialMemory []int, targetOutput int) (int, int) {
	return intcode.FindNounAndVerbForOutput(initialMemory, targetOutput)
}

func main() {
	flag.BoolVar(&verbose, "v", false, "enable verbose logging")
	flag.Parse()

	if verbose {
		debugLog = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)
	} else {
		debugLog = log.New(io.Discard, "", 0)
	}

	    intcode.Init(debugLog)
	
		memory, err := intcode.ReadMemoryFromFile("input")
		if err != nil {
			log.Fatal(err)
		}
		result := execute(*memory)
		log.Default().Printf("Value at position 0: %d\n", result)
	
		debugLog.Println("Memory:", memory)
		noun, verb := findNounAndVerb(*memory, 19690720)
		log.Default().Printf("Noun: %d, Verb: %d, 100 * noun + verb = %d\n", noun, verb, 100*noun+verb)
}
