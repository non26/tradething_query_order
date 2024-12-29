run-echo:
	go run cmd/echo/main.go

go-zip:
	bash generate-function.sh
	bash generate-function-cron-job.sh