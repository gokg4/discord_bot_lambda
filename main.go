package main

import (
	"google_news/discord"

	"github.com/aws/aws-lambda-go/lambda"
)

func handler() error {
	cloudTechnologies := []string{"Cloud%20Computing", "Amazon%20Web%20Services", "Google%20Cloud", "Azure"} // You can change to some other topics to get relevant articles.
	for _, v := range cloudTechnologies {
		discord.NewsCheck(v)
	}

	return nil
}

func main() {
	lambda.Start(handler)
}
