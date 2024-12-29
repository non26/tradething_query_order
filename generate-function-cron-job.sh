#!/bin/bash
GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap cmd/cron_job/main.go
zip tradething_query_position_cron_job.zip bootstrap