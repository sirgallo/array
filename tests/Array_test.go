package arraytests

import "testing"

import "github.com/sirgallo/array"


var testArray []int


func init() {
	testArray = []int{ 1, 2, 3, 4 }
}


func TestArray(t *testing.T) {
	t.Run("test chunk", func(t *testing.T) {
		expected := [][]int{ { 1, 2 }, { 3, 4 } }
		
		chunkArr, chunkErr := array.Chunk[int](testArray, 2)
		if chunkErr != nil { t.Errorf("error chunking array: %s", chunkErr.Error()) }

		t.Logf("actual: (%v), expected: (%v)", chunkArr, expected)
	})

	t.Run("test filter", func(t *testing.T) {
		expected := []int{ 1, 2 }
		
		filterFunc := func(elem int) bool { return elem < 3 }
		filtered := array.Filter[int](testArray, filterFunc)
		
		t.Logf("actual: (%v), expected: (%v)", filtered, expected)

		for idx, val := range filtered {
			if val != expected[idx] { t.Errorf("actual not equal to expected: actual(%v), expected(%v)", val, expected[idx]) }
		}
	})

	t.Run("test map", func(t *testing.T) {
		expected := []int{ 2, 4, 6, 8 }

		mapFunc := func(elem int) int { return elem * 2 }
		mapped := array.Map[int](testArray, mapFunc)

		t.Logf("actual: (%v), expected: (%v)", mapped, expected)

		for idx, val := range mapped {
			if val != expected[idx] { t.Errorf("actual not equal to expected: actual(%v), expected(%v)", val, expected[idx]) }
		}
	})
}