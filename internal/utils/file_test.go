package utils

import (
	"strings"
	"testing"
)

var str string = `this is
multiline
string`

func TestReadLines(t *testing.T) {
	r := strings.NewReader(str)
	lines, err := ReadLines(r)
	AssertEqual(t, err, nil)
	AssertEqual(t, len(lines), 3)
	AssertEqual(t, lines[0], "this is")
	AssertEqual(t, lines[1], "multiline")
	AssertEqual(t, lines[2], "string")
}
