package main

import (
	"os"
	"testing"
)

func mockStdin(t *testing.T, input string, testFunc func()) {
	reader, writer, err := os.Pipe()
	if err != nil {
		t.Fatal(err)
	}

	origStdin := os.Stdin
	os.Stdin = reader

	defer func() { os.Stdin == origStdin }()
}
