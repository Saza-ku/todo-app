package main

import (
	"context"
	"log"
	"todo-app/handler"
	"todo-app/infra"
	"todo-app/usecase"

	"github.com/labstack/echo/v4"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	echolamda "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echolamda.EchoLambda

func init() {
	todoRepo := infra.NewTodoDB()
	todoUc := usecase.NewTodoUseCase(todoRepo)
	controller := handler.NewController(todoUc)

	log.Printf("Echo cold start")

	e := echo.New()
	e.GET("/todo", controller.GetTodo)
	echoLambda = echolamda.New(e)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

func main() {
	lambda.Start(Handler)
}
