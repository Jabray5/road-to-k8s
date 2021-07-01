build:
	go build -o bin/logger main.go
	docker build -t loggerimage ./

run:
	docker run loggerimage:latest