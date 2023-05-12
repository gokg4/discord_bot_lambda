![GitHub](https://img.shields.io/github/license/gokg4/discord_bot_lambda) ![GitHub code size in bytes](https://img.shields.io/github/languages/code-size/gokg4/discord_bot_lambda) ![GitHub repo size](https://img.shields.io/github/repo-size/gokg4/discord_bot_lambda)

# discord_bot_lambda

A Discord Bot(webhook) written in Go and hosted as an AWS Lambda Function. This application is designed to call the Google News RSS feed and post latest news articles to the discord bot based on given search query. We can automate the daily articles posting by setting up AWS EventBridge to schedule event triggers targetting the lambda function.

Install go from the official website - [link](https://go.dev/dl/).

Google News RSS Feed usage blog - [link](https://newscatcherapi.com/blog/google-news-rss-search-parameters-the-missing-documentaiton)

Create your Discord Bot - [link](https://support.discord.com/hc/en-us/articles/360045093012)

Discord Webhook Documentation - [link](https://discord.com/developers/docs/resources/webhook)

## Installation

Give a Try, clone this repository and run the following command in the terminal.

```go run main.go```

Build, Compile and Zip with the following command.

```GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o main main.go```

```zip -r function.zip main```

## Usage

Once built and compiled you can run from command line using the following command.

```./main```

Upload the Zip file to AWS Lambda Function.

## Contributors

- [gokg4](https://github.com/gokg4) - creator and maintainer
