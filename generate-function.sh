#!/bin/bash
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap cmd/aws_lambda/main.go
zip tradething_query_position.zip bootstrap