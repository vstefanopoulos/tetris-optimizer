package main_test

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
	"time"
)

func TestMainOutput(t *testing.T) {
	testCasesGood := []struct {
		inputFile    string
		expectedDots int
	}{
		{"good00.txt", 0},
		{"good01.txt", 9},
		{"good02.txt", 4},
		{"good03.txt", 5},
		{"hard.txt", 1},
		{"hardvar.txt", 1},
	}
	testCasesBad := []struct {
		inputFile     string
		expectedError string
	}{
		{"00.txt", "ERROR"},
		{"01.txt", "ERROR"},
		{"02.txt", "ERROR"},
		{"03.txt", "ERROR"},
		{"04.txt", "ERROR"},
		{"05.txt", "ERROR"},
		{"06.txt", "ERROR"},
		{"07.txt", "ERROR"},
		{"08.txt", "ERROR"},
		{"09.txt", "ERROR"},
	}
	for _, tc := range testCasesBad {
		t.Run(tc.inputFile, func(t *testing.T) {
			cmd := exec.Command("go", "run", "main.go", tc.inputFile)
			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				t.Fatalf("Failed to run main.go with %s: %v", tc.inputFile, err)
			}

			output := strings.TrimSpace(out.String())
			lines := strings.Split(output, "\n")

			fmt.Printf("\nOutput for %s:\n%s\n", tc.inputFile, output)

			if len(lines) < 2 {
				t.Fatalf("Unexpected output format for %s, got: %v", tc.inputFile, output)
			}
			actualError := lines[len(lines)-1]
			if actualError != tc.expectedError {
				t.Errorf("For %s: expected %s error, but got %s error", tc.inputFile, tc.expectedError, actualError)
			}
		})
	}
	for _, tc := range testCasesGood {
		t.Run(tc.inputFile, func(t *testing.T) {
			start := time.Now()
			cmd := exec.Command("go", "run", "main.go", tc.inputFile)

			var out bytes.Buffer
			cmd.Stdout = &out
			err := cmd.Run()
			if err != nil {
				t.Fatalf("Failed to run main.go with %s: %v", tc.inputFile, err)
			}

			end := time.Since(start)
			output := strings.TrimSpace(out.String())
			board := strings.Split(output, "\n")

			fmt.Printf("\nOutput for %s:\n%s\n%v\n", tc.inputFile, output, end)

			if len(board) < 2 {
				t.Fatalf("Unexpected output format for %s, got: %v", tc.inputFile, output)
			}

			actualDots := 0
			for _, line := range board {
				actualDots += strings.Count(line, ".")
			}

			if actualDots != tc.expectedDots {
				t.Errorf("For %s: expected %d dots, but got %d dots", tc.inputFile, tc.expectedDots, actualDots)
			}
		})
	}
}
