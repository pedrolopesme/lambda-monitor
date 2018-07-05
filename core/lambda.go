package core

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"time"
)

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

func (l Lambda) GetName() string {
	return l.name
}

// TODO add tests
func LambdaBuilder(awsFunction *lambda.FunctionConfiguration) (lambda *Lambda) {
	return &Lambda{
		name: *awsFunction.FunctionName,
	}
}
