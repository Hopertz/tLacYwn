package linearsearch

import (
	"fmt"
	"testing"
)

func TestLinearSearch(t *testing.T) {

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
			isFound := linear_search(foo, tt.num)
			if isFound != tt.isFound {
				t.Errorf("got %v, want %v", isFound, tt.isFound)
			}
		})
	}

}
