package main

import (
	"context"
	"log"
	"todo-app/handler"
	"todo-app/infra"
	"todo-app/usecase"

	"github.com/guregu/dynamo"
	"github.com/labstack/echo/v4"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	echolamda "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echolamda.EchoLambda

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: credentials.NewStaticCredentials("dummy", "dummy", "dummy"),
	})
	if err != nil {
		panic(err)
	}

	db := dynamo.New(sess)

	todoRepo := infra.NewTodoDB(db)
	todoUc := usecase.NewTodoUseCase(todoRepo)
	controller := handler.NewController(todoUc)

	log.Printf("Echo cold start")

	e := echo.New()
	e.GET("/todo", controller.GetTodo)
	e.POST("/todo", controller.AddTodo)
	echoLambda = echolamda.New(e)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
