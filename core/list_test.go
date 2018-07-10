package core

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintLambda(test *testing.T) {
	output := captureOutput(func() {
		printLambdaSummary(buildMockedLambda("Dummy"))
	})

	assert.Equal(test, "Dummy\n", output)
}

func TestPrintLambdasWhenTheresNoLambda(test *testing.T) {
	lambdas := Lambdas{}
	output := captureOutput(func() {
		lambdas.Print()
	})
	assert.Equal(test, "No Lambdas Found\n", output)
}

func TestPrintLambdasWhenThereIsOneLambda(test *testing.T) {
	lambdas := Lambdas{}
	lambdas.Put(*buildMockedLambda("one"))

	output := captureOutput(func() {
		lambdas.Print()
	})
	assert.Equal(test, "one\n", output)
}

func TestPrintLambdasWhenThereAreMultipleLambda(test *testing.T) {
	lambdas := Lambdas{}
	lambdas.Put(*buildMockedLambda("one"))
	lambdas.Put(*buildMockedLambda("two"))
	lambdas.Put(*buildMockedLambda("three"))

	output := captureOutput(func() {
		lambdas.Print()
	})
	assert.Equal(test, "one\ntwo\nthree\n", output)
}