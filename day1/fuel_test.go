package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestGetFuel(t *testing.T) {
	testFile, err := os.Open("tests/test")
	if err != nil {
		t.Fatalf("failed to open test file: %s", err)
	}
	defer testFile.Close()

	resultsFile, err := os.Open("tests/results")
	if err != nil {
		t.Fatalf("failed to open results file: %s", err)
	}
	defer resultsFile.Close()

	testScanner := bufio.NewScanner(testFile)
	resultsScanner := bufio.NewScanner(resultsFile)

	for testScanner.Scan() {
		if !resultsScanner.Scan() {
			t.Fatal("results file has fewer lines than test file")
		}

		mass, err := strconv.Atoi(testScanner.Text())
		if err != nil {
			t.Fatalf("failed to parse mass: %s", err)
		}

		expectedFuel, err := strconv.Atoi(resultsScanner.Text())
		if err != nil {
			t.Fatalf("failed to parse expected fuel: %s", err)
		}

		actualFuel := getFuel(mass)

		if actualFuel != expectedFuel {
			t.Errorf("getFuel(%d) = %d; want %d", mass, actualFuel, expectedFuel)
		}
	}

	if resultsScanner.Scan() {
		t.Fatal("test file has fewer lines than results file")
	}

	if err := testScanner.Err(); err != nil {
		t.Fatalf("error scanning test file: %s", err)
	}

	if err := resultsScanner.Err(); err != nil {
		t.Fatalf("error scanning results file: %s", err)
	}

	if recursiveFuel(14) != 2 {
		t.Fatalf("recursiveFuel(14) = %d; want 2", recursiveFuel(14))
	}
	if recursiveFuel(1969) != 966 {
		t.Fatalf("recursiveFuel(1969) = %d; want 966", recursiveFuel(1969))
	}
	if recursiveFuel(100756) != 50346 {
		t.Fatalf("recursiveFuel(100756) = %d; want 50346", recursiveFuel(100756))
	}
}
