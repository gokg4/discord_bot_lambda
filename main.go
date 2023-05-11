package main

import (
	"context"
	"google_news/discord"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, request events.APIGatewayProxyRequest) error {
	cloudTechnologies := []string{"Cloud%20Computing", "Amazon%20Web%20Services", "Google%20Cloud", "Azure"}
	for _, v := range cloudTechnologies {
		discord.NewsCheck(v)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
