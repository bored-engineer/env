# env
A small Golang package for reading values from environment variables, originally written for use in [AWS Lambda](https://aws.amazon.com/lambda/) functions, ex:

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
