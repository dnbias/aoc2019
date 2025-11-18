# Intcode Computer

This program implements an Intcode computer as described in the Advent of Code 2019, Day 2. The Intcode computer is a simple virtual machine that processes a list of integers (memory) according to a set of opcodes.

## Opcodes

The Intcode computer supports the following opcodes:

- `1` (Add): Adds together numbers read from two positions and stores the result in a third position. The three integers immediately after the opcode tell you these three positions - the first two indicate the positions from which to read the input values, and the third indicates the position at which to store the output.
- `2` (Multiply): Works like the `Add` opcode, but it multiplies the two inputs instead of adding them.
- `99` (Halt): Means that the program is finished and should immediately halt.

## Memory

The Intcode computer's memory is a list of integers. The program reads its initial memory from a file named `input`. The `input` file should contain a comma-separated list of integers.

## Execution

The `intcode.Execute` function takes a slice of integers representing the computer's memory and executes the Intcode program. It iterates through the memory, processing opcodes until it encounters a `Halt` opcode.

The `main` function reads the initial memory from the `input` file, restores the "1202 program alarm" state by setting `memory[1]` to `12` and `memory[2]` to `2`, and then executes the program.

## Finding Noun and Verb

The `intcode.FindNounAndVerbForOutput` function is used to find a pair of `noun` and `verb` values that, when placed in `memory[1]` and `memory[2]` respectively, produce a specific `targetOutput` at `memory[0]` after the program executes. This function iterates through all possible `noun` and `verb` values (from 0 to 99) and returns the pair that produces the desired output.

## Usage

To run the program, use the following command:

```bash
go run main.go
```

To enable verbose logging, use the `-v` flag:

```bash
go run main.go -v
```
