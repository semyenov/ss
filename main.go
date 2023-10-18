package main

import (
	"fmt"
)

func MergeSortUniq(items []string) []string {
	if len(items) < 2 {
		return items
	}

	middle := len(items) / 2
	left := MergeSortUniq(items[:middle])
	right := MergeSortUniq(items[middle:])

	return merge(left, right)
}

// merge merges two sorted slices of strings into a single sorted slice.
//
// The function takes two parameters:
// - left: a slice of strings representing the first sorted slice
// - right: a slice of strings representing the second sorted slice
//
// The function returnsgo test -coverprofile cover.out a sorted slice of strings.
func merge(left, right []string) []string {
	res := make([]string, 0, len(left)+len(right))

	i, j := 0, 0
	for i < len(left) && j < len(right) {
		switch {
		case left[i] < right[j]:
			res = append(res, left[i])
			i++
		case left[i] > right[j]:
			res = append(res, right[j])
			j++
		default:
			// equal rewind left and right
			// and append left[i] and right[j]
			res = append(res, left[i])
			i++
			j++
		}
	}

	res = append(res, left[i:]...)
	res = append(res, right[j:]...)

	return res
}

// Difference calculates the difference between two sorted sets of strings.
//
// Parameters:
//   - sortedSet1: a sorted set of strings.
//   - sortedSet2: a sorted set of strings.
//
// Returns:
// - []string: the difference between sortedSet1 and sortedSet2.
func Difference(sortedSet1 []string, sortedSet2 []string) []string {
	difference := []string{}

	i, j := 0, 0
	for i < len(sortedSet1) && j < len(sortedSet2) {
		if sortedSet1[i] < sortedSet2[j] {
			difference = append(difference, sortedSet1[i])
			i++
		} else if sortedSet1[i] > sortedSet2[j] {
			j++
		} else {
			i++
			j++
		}
	}

	return difference
}

// Intersection finds the intersection of two sorted sets.
//
// Parameters:
//   - sortedSet1: a sorted set of strings
//   - sortedSet2: a sorted set of strings
//
// Returns:
//   - a sorted set of strings representing the intersection of sortedSet1 and sortedSet2
func Intersection(sortedSet1 []string, sortedSet2 []string) []string {
	union := []string{}

	i, j := 0, 0
	for i < len(sortedSet1) && j < len(sortedSet2) {
		if sortedSet1[i] < sortedSet2[j] {
			i++
		} else if sortedSet1[i] > sortedSet2[j] {
			j++
		} else {
			union = append(union, sortedSet1[i])
			i++
			j++
		}
	}

	return union
}

// Union takes two sorted sets of strings and returns their union.
//
// Parameters:
//   - sortedSet1: the first sorted set of strings.
//   - sortedSet2: the second sorted set of strings.
//
// Returns:
//   - []string: the sorted union of the two sets.
func Union(sortedSet1 []string, sortedSet2 []string) []string {
	union := make([]string, 0)
	i, j := 0, 0

	for i < len(sortedSet1) && j < len(sortedSet2) {
		if sortedSet1[i] < sortedSet2[j] {
			if len(union) == 0 || sortedSet1[i] != union[len(union)-1] {
				union = append(union, sortedSet1[i])
			}
			i++
		} else if sortedSet1[i] > sortedSet2[j] {
			if len(union) == 0 || sortedSet2[j] != union[len(union)-1] {
				union = append(union, sortedSet2[j])
			}
			j++
		} else {
			if len(union) == 0 || sortedSet1[i] != union[len(union)-1] {
				union = append(union, sortedSet1[i])
			}
			i++
			j++
		}
	}

	for i < len(sortedSet1) {
		if len(union) == 0 || sortedSet1[i] != union[len(union)-1] {
			union = append(union, sortedSet1[i])
		}
		i++
	}

	for j < len(sortedSet2) {
		if len(union) == 0 || sortedSet2[j] != union[len(union)-1] {
			union = append(union, sortedSet2[j])
		}
		j++
	}

	return union
}

func main() {
	set1 := []string{"kiwi", "apple", "banana", "apple", "cherry", "banana", "cherry", "cherry", "kiwi", "apple", "banana", "apple", "cherry", "banana", "cherry", "cherry"}
	set2 := []string{"grape", "apple", "apple", "orange", "kiwi", "lemon", "mango"}

	res := MergeSortUniq(append(set1, set2...))
	fmt.Println(res)
}
