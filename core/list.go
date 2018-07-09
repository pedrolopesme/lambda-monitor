package core

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/pedrolopesme/lambda-monitor/client"
)

// lambdaSummary contains the basic  interface
// for a Lambda in order to be printed
type lambdaSummary interface {
	GetName() string
}

// printLambdaSummary formats and print a lambda
func printLambdaSummary(lambda lambdaSummary) {
	fmt.Println(lambda.GetName())
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
	fmt.Println(lambdas)
	return
}

// Process list the lambdas associated to an
// AWS account and print them
// TODO refactor and add tests
func (lambdas Lambdas) Load() Lambdas {
	sess, err := session.NewSession()
	if err != nil {
		// Something went wrong... Log it.
		return lambdas
	}

	lambdasClient := client.NewLambdaClient(lambda.New(sess))
	awsLambdas, err := listFunctions(lambdasClient.Get())

	if err != nil {
		lambdas.Set(awsLambdas)
	}
	return lambdas
}

// Print func walks through all lambdas stored
// at Lambdas type and print them to the standard output
// TODO refactor and add tests1
func (lambdas Lambdas) Print() Lambdas {
	iterator := lambdas.GetIterator()
	for iterator.Next() {
		printLambdaSummary(iterator.Get())
	}
	return lambdas
}

func NewLambdaList() Lambdas {
	return Lambdas{}
}
