package core

import (
	"github.com/aws/aws-sdk-go/service/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"fmt"
	"github.com/aws/aws-sdk-go/service/lambda/lambdaiface"
	"github.com/pedrolopesme/lambda-monitor/client"
)

func listFunctions(client lambdaiface.LambdaAPI) (lambdas []Lambda, err error){
	input := &lambda.ListFunctionsInput{}
	result, err := client.ListFunctions(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case lambda.ErrCodeServiceException:
				fmt.Println(lambda.ErrCodeServiceException, aerr.Error())
			case lambda.ErrCodeTooManyRequestsException:
				fmt.Println(lambda.ErrCodeTooManyRequestsException, aerr.Error())
			case lambda.ErrCodeInvalidParameterValueException:
				fmt.Println(lambda.ErrCodeInvalidParameterValueException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
	return
}

// ListLambdas list the lambdas associated to an
// AWS account
// TODO refactor
func PrintLambdas() (err error) {
	sess, err := session.NewSession()
	client := client.NewLambdaClient(lambda.New(sess))
	listFunctions(client.GetClient())
	return
}
