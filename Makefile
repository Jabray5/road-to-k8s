
build:
# Build Go and Docker images
	go build -o bin/logger main.go
	docker build -t logger ./

release: build
# Push new image to docker hub as user-specified version AND as latest
# "make ver=0.1 release"
	docker tag logger:latest jabray5/logger:$(ver)
	docker push jabray5/logger:$(ver)
	docker tag logger:latest jabray5/logger:latest
	docker push jabray5/logger:latest


run:
# Run latest image in a docker container
	docker run --name logcontainer jabray5/logger:latest

deploy:
# Kubernetes
	kubectl apply -f deployment.yaml