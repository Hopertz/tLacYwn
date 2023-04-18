package twocrystalballs

import (
	"math"
)

// Question
// Given two crystal balls that will break if dropped from high enough
// distance, determine the exact spot in which it will break in the most
// optimized way.

func two_crystall_balls(array []bool) int {

	arrayLength := float64(len(array))
	jumpAmount := math.Floor(math.Sqrt(arrayLength))

	i := jumpAmount

	for ; i < arrayLength; i += jumpAmount {
		if array[int(i)] {
			break
		}
	}

	i -= jumpAmount

	for j := 0; j < int(jumpAmount) && i < arrayLength; j, i = j+1, i+1 {
		if array[int(i)] {
			return int(i)
		}
	}

	return -1

}
