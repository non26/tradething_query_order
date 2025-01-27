run-echo:
	go run cmd/echo/main.go

go-zip:
	bash generate-function.sh

go-zip-cron-job:
	bash generate-function-cron-job.sh