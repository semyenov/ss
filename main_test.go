package main

import (
	"math/rand"
	"reflect"
	"sort"
	"strconv"
	"testing"
	"time"
)

type testCase struct {
	name     string
	input    []string
	expected []string
}

var testCases = testCaseGenerator(500, 100000)

func TestMergeSortUniq(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := MergeSortUniq(tc.input)
			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("Expected %v, but got %v", tc.expected, result)
			}
		})
	}
}

func testCaseGenerator(numCases, inputSize int) []testCase {
	// Prepare test cases
	rand.New(rand.NewSource(time.Now().UnixNano())) // Create a new random number generator
	testCases := make([]testCase, numCases)         // Create a slice to store the test cases

	for i := 0; i < numCases; i++ {
		// Generate expected result
		expectedSize := rand.Intn(inputSize)              // Generate a random expected size
		expected := generatUniqeSortedSlice(expectedSize) // Generate a unique sorted slice of the expected size

		// Generate input
		input := make([]string, inputSize) // Create a slice to store the input
		for i := 0; i < inputSize; i++ {
			input[i] = expected[i%expectedSize] // Set each element of the input slice to a value from the expected slice
		}
		rand.Shuffle(inputSize, func(i, j int) {
			input[i], input[j] = input[j], input[i] // Shuffle the input slice randomly
		})

		// Generate test case
		testCases[i] = testCase{ // Create a new testCase struct with the generated input, expected result, and name
			name:     strconv.Itoa(i),
			input:    input,
			expected: expected,
		}
	}

	return testCases // Return the generated test cases
}

func generatUniqeSortedSlice(size int) []string {
	// generate a slice of size`
	slice := make([]string, size)
	for i := 0; i < size; i++ {
		slice[i] = strconv.Itoa(i)
	}

	sort.Strings(slice)
	return slice
}
