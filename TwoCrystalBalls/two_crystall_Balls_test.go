package twocrystalballs

import (
	"math/rand"
	"testing"
)

func TestTwoCrystallBalls(t *testing.T) {

	t.Run("test crystall balls with  breaking point", func(t *testing.T) {
		idx := rand.Intn(10000)

		data := make([]bool, 10000)
		for i := range data {
			data[i] = false
		}

		for i := idx; i < 10000; i++ {
			data[i] = true
		}

		res := two_crystall_balls(data)

		if res != idx {
			t.Errorf("Error want %d got %d", idx, res)
		}
	})

	t.Run("test crystall balls with no breaking point", func(t *testing.T) {
		data := make([]bool, 821)
		for i := range data {
			data[i] = false
		}

		res := two_crystall_balls(data)

		if res != -1 {
			t.Errorf("Error want %d got %d", -1, res)
		}

	})

}
