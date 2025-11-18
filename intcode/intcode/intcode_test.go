package intcode

import (
	"log"
	"os"
	"testing"
)

func findNounAndVerbTestProgram(program_file string, target_output int, expected_noun int, expected_verb int, t *testing.T) {
	var debugLog *log.Logger
	debugLog = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)
	Init(debugLog)

	program, err := ReadMemoryFromFile(program_file)
	if err != nil {
		t.Fatalf("Error reading program file: %v", err)
	}
	findNoun, findVerb := FindNounAndVerbForOutput(*program, target_output)

	if findNoun != expected_noun || findVerb != expected_verb {
		t.Errorf("FindNounAndVerbForOutput(%d, %d) = (%d, %d); want (%d, %d)",
			*program, target_output, findNoun, findVerb, expected_noun, expected_verb)
	}

}

func runTestProgram(program_file string, expected_file string, t *testing.T) {

	var debugLog *log.Logger
	debugLog = log.New(os.Stdout, "[DEBUG] ", log.LstdFlags)
	Init(debugLog)

	program, err := ReadMemoryFromFile(program_file)
	if err != nil {
		t.Fatalf("Error reading program file: %v", err)
	}
	expected, err := ReadMemoryFromFile(expected_file)
	if err != nil {
		t.Fatalf("Error reading expected file: %v", err)
	}
	results := Execute(*program)

	if MemoryEquals(results, *expected) == false {
		t.Errorf("Execute(%d) = %d; want %d", *program, results, *expected)
	}
}

func TestIntcode(t *testing.T) {
	runTestProgram("tests/program_1", "tests/expected_1", t)
	runTestProgram("tests/program_2", "tests/expected_2", t)
	runTestProgram("tests/program_3", "tests/expected_3", t)
	runTestProgram("tests/program_4", "tests/expected_4", t)
	runTestProgram("tests/program_5", "tests/expected_5", t)
	findNounAndVerbTestProgram("../input", 3706713, 12, 2, t)


}
