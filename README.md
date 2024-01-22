# Golang Environment Variables [![Go Reference](https://pkg.go.dev/badge/github.com/bored-engineer/env.svg)](https://pkg.go.dev/github.com/bored-engineer/env)
A small Golang package for parsing values from environment variables based on [github.com/spf13/cast](https://github.com/spf13/cast).

### Usage
This package was originally written for use in simple [AWS Lambda](https://aws.amazon.com/lambda/) functions to reduce boilerplate, ex:
```golang
package main

import (
    "fmt"
    "context"

    "github.com/aws/aws-lambda-go/lambda"
    "github.com/bored-engineer/env"
)

var FunctionName = env.String("AWS_LAMBDA_FUNCTION_NAME")

func invoke(ctx context.Context) {
    fmt.Printf("Hello from %s", FunctionName)
}

func main() {
    lambda.Start(invoke)
}
```
