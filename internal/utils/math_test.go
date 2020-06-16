package utils

import "testing"

func TestMin(t *testing.T) {
	AssertEqual(t, Min(3, 5), 3)
}

func TestMax(t *testing.T) {
	AssertEqual(t, Max(3, 5), 5)
}
