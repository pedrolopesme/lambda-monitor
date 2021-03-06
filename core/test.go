package core

import (
	"bytes"
	"io"
	"os"
)

// Captures the output by replacing the
// standard stdout with a buffer and then restore it again.
// Returns the buffer contents as String
func captureOutput(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f()

	w.Close()
	os.Stdout = old

	var buf bytes.Buffer
	io.Copy(&buf, r)
	return buf.String()
}

func buildMockedLambda(name string) *Lambda {
	return &Lambda{
		name: name,
	}
}
