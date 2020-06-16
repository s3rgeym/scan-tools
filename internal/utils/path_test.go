package utils

import (
	"os"
	"path"
	"testing"
)

func TestExpandPath(t *testing.T) {
	home, _ := os.UserHomeDir()
	result, err := ExpandPath("~/foo/bar")
	AssertEqual(t, err, nil)
	AssertEqual(t, result, path.Join(home, "foo", "bar"))
}
