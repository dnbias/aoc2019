package main;

import (
	"aoc2019/day4/passwords"
	"strconv"
)

// damn elves threw out the password
// 6 digit number
// value in range of input
// two adjacent digits are the same
// from left to right, the digits never decrease
// -- only increase or stay the same 
//
// how many passwords are valid
func main() {
	input := "138241-674034"
	lower , _:= strconv.Atoi(input[:6])
	upper , _:=	strconv.Atoi(input[7:]) 

	result := passwords.CalculatePossiblePasswords(lower, upper)

	println(result)
}
