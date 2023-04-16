package binarysearch

import (
	"fmt"
	"testing"
)

func TestBinarySearch(t *testing.T) {

	var foo = []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	var tests = []struct {
		num     int
		isFound bool
	}{
		{69, true},
		{1336, false},
		{69420, true},
		{69421, false},
		{1, true},
		{0, false},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Test for %d ", tt.num)

		t.Run(testname, func(t *testing.T) {
			isFound := binary_search(foo, tt.num)
			if isFound != tt.isFound {
				t.Errorf("got %v, want %v", isFound, tt.isFound)
			}
		})
	}

}

func TestBinarySearchForDiffList(t *testing.T) {

	var tests = []struct {
		lists   []int
		num     int
		isFound bool
	}{
		{[]int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}, 69, true},
		{[]int{1}, 1, true},
		{[]int{98, 99}, 99, true},
		{[]int{5, 6, 7}, 7, true},
		{[]int{5, 6, 7, 8}, 8, true},
		{[]int{5, 6, 7}, 5, true},
		{[]int{5, 6, 7, 8, 9, 11}, 11, true},
	}
	for _, tt := range tests {
		testname := fmt.Sprintf("Test for %v  find %d ", tt.lists, tt.num)

		t.Run(testname, func(t *testing.T) {
			isFound := binary_search(tt.lists, tt.num)
			if isFound != tt.isFound {
				t.Errorf("got %v, want %v", isFound, tt.isFound)
			}
		})
	}

}
