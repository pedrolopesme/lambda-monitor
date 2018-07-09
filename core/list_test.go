package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintLambda(test *testing.T) {
	output := captureOutput(func() {
		printLambda(buildMockedLambda("Dummy"))
	})

	assert.Equal(test, "Dummy\n", output)
}