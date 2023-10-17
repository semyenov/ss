package main

import (
	"fmt"
)

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

func MergeSortUniq(items []string) []string {
	if len(items) < 2 {
		return items
	}

	middle := len(items) / 2
	left := MergeSortUniq(items[:middle])
	right := MergeSortUniq(items[middle:])

	return merge(left, right)
}

func main() {
	set1 := []string{"kiwi", "apple", "banana", "apple", "cherry", "banana", "cherry", "cherry", "kiwi", "apple", "banana", "apple", "cherry", "banana", "cherry", "cherry"}
	set2 := []string{"grape", "apple", "apple", "orange", "kiwi", "lemon", "mango"}

	res := MergeSortUniq(append(set1, set2...))
	fmt.Println(res)
}
