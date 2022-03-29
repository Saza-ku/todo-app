package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"todo-app/handler"
	"todo-app/infra"
	"todo-app/usecase"

	"github.com/guregu/dynamo"
	"github.com/labstack/echo/v4"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	echolamda "github.com/awslabs/aws-lambda-go-api-proxy/echo"
)

var echoLambda *echolamda.EchoLambda
var dynamoEndpoint = os.Getenv("DYNAMODB_ENDPOINT")

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-northeast-1"),
		Endpoint: &dynamoEndpoint,
	})
	if err != nil {
		panic(err)
	}

	db := dynamo.New(sess)

	todoRepo := infra.NewTodoDB(db)
	todoUc := usecase.NewTodoUseCase(todoRepo)
	controller := handler.NewController(todoUc)

	log.Printf("Echo cold start")
	fmt.Println(dynamoEndpoint)

	e := echo.New()
	e.GET("/health", func(c echo.Context) error {
		fmt.Println("STATUS OK")
		return c.String(http.StatusOK, "STATUS OK")
	})
	e.GET("/todo", controller.GetTodo)
	e.POST("/todo", controller.AddTodo)
	e.PUT("/todo/:id", controller.EditTodo)
	e.DELETE("/todo/:id", controller.RemoveTodo)
	echoLambda = echolamda.New(e)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return echoLambda.ProxyWithContext(ctx, req)
}

// @title           Todo API
// @version         1.0
// @description     This is a Todo-app server.

// @host      localhost:8080
// @BasePath  /
func main() {
	lambda.Start(Handler)
}
