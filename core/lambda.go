package core

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"time"
)

// Lambda represents a AWS Lambda with all
// properties need by lambda-monitor to run
type Lambda struct {
	name          string
	lastModified  time.Time
	codeSha256    string
	codeSize      uint32
	description   string
	variables     map[string]string
	arn           string
	handler       string
	revisionId    string
	role          string
	runtime       string
	timeout       uint32
	tracingConfig map[string]string
	version       string
}

// GetName allows others to retrieve Lambda name
// TODO add tests
func (l Lambda) GetName() string {
	return l.name
}

// LambdaBuilder knows how to build a Lambda from a FunctionConfiguration
// provided by AWS package
// TODO add tests
func LambdaBuilder(awsFunction *lambda.FunctionConfiguration) (lambda *Lambda) {
	return &Lambda{
		name: *awsFunction.FunctionName,
	}
}

// Lambdas stores a collection of Lambdas
// and collection related funcs to operate over them
type Lambdas struct {
	lambdas []Lambda
}

// Get returns the internal array of lambdas from Lambdas type
// TODO add tests
func (lambdas Lambdas) Get() []Lambda {
	return lambdas.lambdas
}

// Set replaces the internal array of lambdas
// TODO add tests
func (lambdas Lambdas) Set(newLambdas []Lambda) {
	lambdas.lambdas = newLambdas
}

// Size returns the amount of lambdas found
// TODO add tests
func (lambdas Lambdas) Size() int {
	return len(lambdas.lambdas)
}

// Size returns the amount of lambdas found
// TODO add tests
func (lambdas *Lambdas) Put(lambda Lambda) {
	lambdas.lambdas = append(lambdas.lambdas, lambda)
}

// GetIterator retrieves an iterator to the
// current Lambdass
// TODO add tests
func (lambdas Lambdas) GetIterator() *LambdasIterator {
	return &LambdasIterator{
		lambdas: lambdas,
		head:    -1,
	}
}

// LambdasIterator contains a temporary data
// to iterate over a Lambdas type
type LambdasIterator struct {
	lambdas Lambdas
	head    int
}

// Next moves the internal head of the iterator to the
// next position if its possible
// TODO add tests
func (iterator *LambdasIterator) Next() bool {
	if iterator.head < len(iterator.lambdas.Get())-1 {
		iterator.head++
		return true
	}
	return false
}

// Retrieves the lambda stored in the current head position
// TODO add tests
func (iterator LambdasIterator) Get() (lambda Lambda) {
	if iterator.head >= 0 && iterator.head < len(iterator.lambdas.Get()) {
		lambda = iterator.lambdas.Get()[iterator.head]
	}
	return
}
