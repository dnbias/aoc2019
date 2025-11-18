# Intcode Computer

This program implements an Intcode computer as described in the Advent of Code 2019. The Intcode computer is a simple virtual machine that processes a list of integers (memory) according to a set of opcodes. This implementation has been extended to support the full set of opcodes and parameter modes from later challenges.

## Opcodes

The Intcode computer supports the following opcodes:

- `1` (Add): Adds together numbers from two positions and stores the result in a third position.
- `2` (Multiply): Works like the `Add` opcode, but it multiplies the two inputs instead of adding them.
- `3` (Input): Takes a single integer as input and saves it to the position given by its only parameter.
- `4` (Output): Outputs the value of its only parameter.
- `5` (Jump-if-true): If the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter.
- `6` (Jump-if-false): If the first parameter is zero, it sets the instruction pointer to the value from the second parameter.
- `7` (Less than): If the first parameter is less than the second parameter, it stores `1` in the position given by the third parameter. Otherwise, it stores `0`.
- `8` (Equals): If the first parameter is equal to the second parameter, it stores `1` in the position given by the third parameter. Otherwise, it stores `0`.
- `99` (Halt): Means that the program is finished and should immediately halt.

## Parameter Modes

Instructions can have parameters in one of two modes:

- `0` (Position Mode): The parameter is interpreted as an address. The value stored at that address is used.
- `1` (Immediate Mode): The parameter is interpreted as a value directly.

Parameter modes are specified by the instruction's value. For an instruction like `1002`, `02` is the opcode (Multiply), `0` is the mode of the 1st parameter, `1` is the mode of the 2nd parameter, and the 3rd parameter's mode is `0` (omitted leading zero).

## Memory

The Intcode computer's memory is a list of integers. The program can read its initial memory from a file.

## Execution

The `intcode.Execute` function takes a slice of integers representing the computer's memory and executes the Intcode program. It iterates through the memory, processing opcodes and parameter modes until it encounters a `Halt` opcode.

## Helper Functions

- `FindNounAndVerbForOutput`: Used for the Day 2 challenge. It finds a pair of `noun` and `verb` values that, when placed in `memory[1]` and `memory[2]`, produce a specific `targetOutput` at `memory[0]`.
- `TEST`: Runs a diagnostic program for the ship's Air Conditioner Unit (ACU) and Thermal Radiator Controller (TRC).

## Intcode Programs

The `programs/` directory contains several sample Intcode programs:

- `cmp8.i`: Takes an integer input. Outputs `1000` if the input is equal to 8, `1001` if it is greater than 8, and `999` if it is less than 8.
- `diagnostic.i`: A diagnostic program for the thermal environment supervision unit. It verifies that the ship's air conditioner unit and thermal radiator controller are functioning correctly.
- `eq8.i`: Takes an integer input and outputs `1` if the input is equal to 8, and `0` otherwise. Uses position mode.
- `eq8immediate.i`: Takes an integer input and outputs `1` if the input is equal to 8, and `0` otherwise. Uses immediate mode.
- `lt8.i`: Takes an integer input and outputs `1` if the input is less than 8, and `0` otherwise. Uses position mode.
- `lt8immediate.i`: Takes an integer input and outputs `1` if the input is less than 8, and `0` otherwise. Uses immediate mode.
- `nonzero.i`: Takes an integer input and outputs `1` if the input is non-zero, and `0` otherwise. Uses position mode.
- `nonzero-immediate.i`: Takes an integer input and outputs `1` if the input is non-zero, and `0` otherwise. Uses immediate mode.

## Usage

To run the program, use the following command:

```bash
go run main.go
```

To enable verbose logging, use the `-v` flag:

```bash
go run main.go -v
```
