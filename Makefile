build:
	go build -o bin/logger main.go
	docker build -t logger ./

release: build
# Implement docker push when new version
	docker tag logger:latest jabray5/logger:$(ver)
	docker push jabray5/logger:$(ver)
	docker tag logger:latest jabray5/logger:latest
	docker push jabray5/logger:latest


run:
	docker run --name logcontainer jabray/loggerimage:latest

deploy:
	kubectl apply -f deployment.yaml