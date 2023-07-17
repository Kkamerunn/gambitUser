package main

import (
	"context"
	"os"
	"fmt"
	"errors"
	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/Kkamerunn/gambitUser/awsgo"
	//"github.com/Kkamerunn/gambitUser/models"
)

func main() {
	lambda.Start(LambdaExec)
}

func LambdaExec(ctx context.Context, event events.CognitoEventUserPoolsPostConfirmation) (events.CognitoEventUserPoolsPostConfirmation, error) {
	awsgo.InitializeAWS()

	if !ValidParams() {
		fmt.Println("Error in the params. You must send 'SecretName'")
		err :=  errors.New("Error in the params. You must send 'SecretName'")
		return event, err
	}
}

func ValidParams() bool {
	var getParam bool
	_, getParam = os.LookupEnv("SecretName")
	return getParam
}