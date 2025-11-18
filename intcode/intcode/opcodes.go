package intcode

type Opcode int

// first 2 digits (rightmost) are the opcode
// each digit after that is the mode of the parameter in each position
//
// Opcode 3 takes a single integer as input and saves it to the position given by its only parameter.
// For example, the instruction 3,50 would take an input value and store it at address 50.

// Opcode 4 outputs the value of its only parameter.
// For example, the instruction 4,50 would output the value at address 50.
const (
	Add      Opcode = 1 // input i, input j, parameter: position
	Multiply Opcode = 2 // input i, input j, parameter: position
	Store 	 Opcode = 3 // input i, parameter: position
	Outputs	 Opcode = 4 // input -, parameter: position
	Halt     Opcode = 99
)

// ABCDE
//  1002
//
// DE - two-digit opcode,      02 == opcode 2
//  C - mode of 1st parameter,  0 == position mode
//  B - mode of 2nd parameter,  1 == immediate mode
//  A - mode of 3rd parameter,  0 == position mode,
//                                   omitted due to being a leading zero
type Mode int
const (
	Position	Mode = 0 // pointer to parameter value
	Immediate 	Mode = 1 // parameter value
)

// 1002,4,3,4,33
// opcode = 02
// parameter #1 = 4 in mode 0
// parameter #2 = 3 in mode 1
// parameter #3 = 4 in mode 0 (omitted)
// memory[4] = 33 * 3 = 99
// 1002,4,3,4,99
