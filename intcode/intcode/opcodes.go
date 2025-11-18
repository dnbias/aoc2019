package intcode

type Opcode int

const (
	Add      Opcode = 1
	Multiply Opcode = 2
	Halt     Opcode = 99
)
