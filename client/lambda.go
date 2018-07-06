package client

import "github.com/aws/aws-sdk-go/service/lambda/lambdaiface"

// TODO add comments
type lambdaClient struct {
	client lambdaiface.LambdaAPI
}

func (lc lambdaClient) Get() lambdaiface.LambdaAPI {
	return lc.client
}

func NewLambdaClient(client lambdaiface.LambdaAPI) *lambdaClient {
	return &lambdaClient{
		client: client,
	}
}
