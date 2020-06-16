package utils

import "testing"

func TestFilterInt(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	results := FilterInt(nums, func(x int) bool {
		return x > 3
	})
	AssertEqual(t, len(results), 2)
	AssertEqual(t, results[0], 4)
	AssertEqual(t, results[1], 5)
}
