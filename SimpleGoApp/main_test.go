package main

import (
	"io"
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

	go func() {
		defer writer.Close()
		io.WriteString(writer, input)
	}()

	testFunc()
}

func TestAskForNumber(t *testing.T) {
	input := "42\n"
	expected := 42
	result := 0

	mockStdin(t, input, func() {
		result = askForNumber()
	})

	if result != expected {
		t.Errorf("")
	}
}
