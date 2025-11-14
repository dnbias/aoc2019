package passwords

import "math"

// damn elves threw out the password
// 6 digit number
// value in range of input
// two adjacent digits are the same
// from left to right, the digits never decrease
// -- only increase or stay the same 
//
// how many passwords are valid
func CalculatePossiblePasswords(lower, upper int) int {
	result := 0

	for n := lower; n <= upper; n++ {
		if isValid(n) {
			result++
		}
	}

	return result
}

func isValid(n int) bool {
	contains_same_adjacent := false
	digits_adjacent := make(map[int]int)

	// check for non decreasing digits
	for i := 0; i < 5; i++ {
		m := n / int(math.Pow10(i))
		currentDigit := m % 10
		nextDigit := (m / 10) % 10
		if currentDigit < nextDigit {
			return false
		}
		
		// check for adjacent digits
		if currentDigit == nextDigit {
			digits_adjacent[currentDigit]++
		}
	}

	// for part 2
	for _ , v := range digits_adjacent {
		if v == 1 {
			contains_same_adjacent = true
		}
	}

	return contains_same_adjacent
}
