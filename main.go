package main

import (
	"fmt"

	"context"

	"github.com/aws/aws-lambda-go/events"
	aws "github.com/aws/aws-lambda-go/lambda"
	/*
		"errors"
		"os"
		"strings"



		"github.com/livingstone23/gambit/awsgo"
		"github.com/livingstone23/gambit/bd"
		"github.com/livingstone23/gambit/models"*/)

func main() {
	fmt.Println("Iniciando aplicacion GAMBIT")

	aws.Start(EjecutoLambda)
}

func EjecutoLambda(ctx context.Context, event events.APIGatewayV2HTTPRequest) (*events.APIGatewayProxyResponse, error) {

}
