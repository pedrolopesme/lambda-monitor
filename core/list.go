package core

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/pedrolopesme/lambda-monitor/client"
	"fmt"
)

// TODO add description and tests
type lambdaSummary interface {
	GetName() string
}

// TODO add description and tests
func printLambda(lambda lambdaSummary) {
	fmt.Println(lambda.GetName())
}

// TODO add description and tests
func printLambdas(lambdas []Lambda) {
	if len(lambdas) > 0 {
		for _, lambda := range lambdas {
			printLambda(lambda)
		}
	}
}

// TODO add description and tests
func listFunctions(client lambdaiface.LambdaAPI) (lambdas []Lambda, err error) {
	input := &lambda.ListFunctionsInput{}
	result, err := client.ListFunctions(input)

	if err != nil {
		return
	}

	for _, function := range result.Functions {
		lambdas = append(lambdas, *LambdaBuilder(function))
	}
	return
}

type LambdaList struct {}

// Process list the lambdas associated to an
// AWS account and print them
// TODO refactor and add tests
func (l LambdaList) Print() error {
	sess, err := session.NewSession()
	if err != nil {
		return ErrCannotCreateSession
	}

	client := client.NewLambdaClient(lambda.New(sess))
	lambdas, err := listFunctions(client.GetClient())
	if err != nil {
		fmt.Println("It was impossible to retrieve AWS Lambda Functions due to", err)
	}

	printLambdas(lambdas)
	return nil
}

func NewList() LambdaList {
	return LambdaList{}
}